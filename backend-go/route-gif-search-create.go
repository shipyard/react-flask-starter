package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/jackc/pgx/v5"
    "github.com/stoewer/go-strcase"
    "golang.org/x/sync/errgroup"
)

const SEARCH_LIMIT = "5"

func gifSearchCreate(w http.ResponseWriter, r *http.Request) {
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        serverError(w, "Could not connect to database.", err)
        return
    }
    defer conn.Close(context.Background())

    var requestUrl = "https://api.giphy.com/v1/gifs/search"

    client := &http.Client{}

    var apiKey = os.Getenv("GLIPHY_API_KEY")
    var searchTerms = r.URL.Query()["q"]
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
    const retryAttempts = 4
    retries := 0
    for giphyResp == nil && retries < retryAttempts { // will actually exit before failing loop, but for clarity
        giphyResp, err = client.Do(req)
        retries += 1

        if err != nil { // network error or something? let's retry
            log.Println("Recevied error '%v' attempting Giphy request.", err)
            if retries >= retryAttempts {
                serverError(w, "Could not execute request.", err)
                return
            }
        } else if giphyResp.StatusCode >= 400 { // if err is nil, then giphyResp is not
            // none of these errors are recoverable
            if giphyResp.StatusCode == 414 {
                http.Error(w, "Query too long.", http.StatusBadRequest)
                return
            } // else
            serverError(
                w, 
                fmt.Sprintf("Giphy request failed with unrecoverable error code %d.", 
                giphyResp.StatusCode,
            ), err)
            return
        }
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
    filesReq, err := http.NewRequest("GET", "http://backend:8080/api/v1/files/", nil)
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
    fmt.Println(filesMsg) // DEBUG

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