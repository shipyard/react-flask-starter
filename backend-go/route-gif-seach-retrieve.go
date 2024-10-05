package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strconv"

    "github.com/jackc/pgx/v5"
)

func gifSearchGet(w http.ResponseWriter, r *http.Request) {
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        serverError(w, "Could not connect to database.", err)
        return
    }
    defer conn.Close(context.Background())

    idString := r.PathValue("id")
    _, err = strconv.Atoi(idString)
    if err != nil {
        http.Error(w, fmt.Sprintf("Invalid search ID '%s'; must be an integer.", idString), http.StatusBadRequest)
        return
    }

    var gifUrls []string
    err = conn.QueryRow(context.Background(), "SELECT gif_urls FROM gif_search WHERE id=$1", idString).Scan(&gifUrls)
    if err != nil {
        if err == pgx.ErrNoRows {
            http.Error(w, fmt.Sprintf("No such search ID '%s'.", idString), http.StatusBadRequest)
            return
        }
        serverError(w, "There was an error while retrieving your search.", err)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(gifUrls)
}