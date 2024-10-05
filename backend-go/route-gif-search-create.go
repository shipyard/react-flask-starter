package main

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/cenkalti/backoff/v4"
    "github.com/jackc/pgx/v5"
    "github.com/stoewer/go-strcase"
    "golang.org/x/sync/errgroup"
)

const SEARCH_LIMIT = "5"
const GIPHY_QUERY_KEY = "giphy_api_key"
const GIPHY_ENV_KEY = "GIPHY_API_KEY"

func gifSearchCreate(w http.ResponseWriter, r *http.Request) {
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        serverError(w, "Could not connect to database.", err)
        return
    }
    defer conn.Close(context.Background())

    var requestUrl = "https://api.giphy.com/v1/gifs/search"

    client := &http.Client{}

    reqQuery := r.URL.Query()

    var apiKey string
    apiKeyParam := reqQuery[GIPHY_QUERY_KEY]
    if apiKeyParam == nil {
        apiKey = os.Getenv(GIPHY_ENV_KEY)
    } else {
        apiKey = apiKeyParam[0]
    }
    if apiKey == "" {
        http.Error(w, 
            fmt.Sprintf("Did not find Giphy API key in request parameter '%s' or server environment variable '%s'.",
                GIPHY_QUERY_KEY, 
                GIPHY_ENV_KEY,
            ),
            http.StatusBadRequest,
        )
        return
    }
    searchTerms := reqQuery["q"]
    if searchTerms == nil {
        http.Error(w, "Must provide at least one query parameter 'q'.", http.StatusBadRequest)
        return
    }

    req, err := http.NewRequest("GET", requestUrl, nil)
    if err != nil {
        serverError(w, fmt.Sprintf("Could not create request for URL '%s'.", requestUrl), err)
        return
    }

    q := req.URL.Query()
    q.Add("api_key", apiKey)
    q.Add("q", strings.Join(searchTerms, ","))
    q.Add("limit", SEARCH_LIMIT)
    req.URL.RawQuery = q.Encode() // go is weird how you set the URL parameters

    var giphyResp *http.Response
    GiphySearch := func () error {
        giphyResp, err = client.Do(req)

        if err != nil { // network error or something? let's retry
            log.Println("Received error '%v' attempting Giphy request.", err)
            return err
        } else if giphyResp.StatusCode >= 400 { // if err is nil, then giphyResp is not
            // none of these errors are recoverable
            permError := backoff.Permanent(errors.New("Donso"))
            if giphyResp.StatusCode == 414 {
                http.Error(w, "Query too long.", http.StatusBadRequest)
                return permError
            } // else
            serverError(
                w, 
                fmt.Sprintf("Giphy request failed with unrecoverable error code %d.", 
                giphyResp.StatusCode,
            ), err)
            return permError
        }
        return nil
    }

    expBackoff := backoff.NewExponentialBackOff()
    expBackoff.InitialInterval = 250 * time.Millisecond
    expBackoff.MaxInterval = 3 * time.Second
    expBackoff.MaxElapsedTime = 15 * time.Second

    err = backoff.Retry(GiphySearch, expBackoff)
    if err != nil {
        // if it's not a permanent error, then we exceeded the retry timeout and haven't set the HTTP response yet
        if _, isPerm := err.(*backoff.PermanentError); isPerm == false {
            // then we timed out
            serverError(w, "Exceeded maximum retries while attempting to contact Giphy service.", err)
        }
        // else it's a PermanentError and we've already set the HTTP response, only need to return
        return // we're done in any case
    }

    defer giphyResp.Body.Close()
    bodyBytes, err := io.ReadAll(giphyResp.Body)
    if err != nil {
        serverError(w, "Could not read response body.", err)
        return
    }

    var searchResults GiphySearchResults
    if err := json.Unmarshal(bodyBytes, &searchResults); err != nil {  // Parse []byte to the go struct pointer
        serverError(w, "Could not unmarshal results.", err)
        return
    }

    // Must call this to create the S3 bucket.
    filesReq, err := http.NewRequest("GET", "http://backend:8082/api/v1/files/", nil)
    if err != nil {
        serverError(w, "Error creating S3 bucket request.", err)
        return
    }
    filesRes, err := client.Do(filesReq)
    if err != nil {
        serverError(w, "Error creating S3 bucket.", err)
        return
    }
    buf := new(strings.Builder)
    io.Copy(buf, filesRes.Body)
    filesMsg := fmt.Sprintln(buf.String())

    g := new(errgroup.Group)
    var gifUrls []string
    var s3Paths []string
    for _, searchResult := range searchResults.Data {
        gifUrl := searchResult.Images.Original.URL
        gifUrls = append(gifUrls, gifUrl)

        var termBits []string
        for _, term := range searchTerms {
            termBits = append(termBits, strcase.KebabCase(term))
        }
        fileName := searchResult.ID +
            "_" + strcase.KebabCase(searchResult.Title) +
            "_" + strings.Join(termBits, "_") +
            ".gif"
        s3Path := "test-bucket/" + fileName
        s3Paths = append(s3Paths, s3Path)

        // Download GIFs and upload to S3 asynchronously for efficiency
        g.Go(func () error {
            return UploadFileFromUrl(client, gifUrl, fileName)
        })
    }

    // wait on all pending downloads to succeed or fail
    if err = g.Wait(); err != nil {
        serverError(w, "Error uploading GIFs to S3 storage.", err)
        return
    }

    var id int32
    err = conn.QueryRow(
        context.Background(),
        "INSERT INTO gif_search (search_terms, gif_urls, s3_paths) VALUES ($1, $2, $3) RETURNING id",
        searchTerms,
        gifUrls,
        s3Paths,
    ).Scan(&id)
    if err != nil {
        serverError(w, "Creating record in database failed.", err)
        return
    }

    responseData := SearchResponse{ID: id, GifUrls: gifUrls, S3Paths: s3Paths}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(responseData)
}