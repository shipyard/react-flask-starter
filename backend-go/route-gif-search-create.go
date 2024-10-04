package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
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
    req.URL.RawQuery = q.Encode()

    resp, err := client.Do(req)
    if err != nil {
        serverError(w, "Could not execute request.", err)
        return
    }

    defer resp.Body.Close()
    bodyBytes, err := io.ReadAll(resp.Body)
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
        "INSERT INTO gif_search (search_terms, gif_urls) VALUES ($1, $2) RETURNING id",
        searchTerms,
        gifUrls,
    ).Scan(&id)
    if err != nil {
        serverError(w, "Creating record in database failed.", err)
        return
    }

    responseData := SearchResponse{ID: id, GiphyResponse: searchResults}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(responseData)
}