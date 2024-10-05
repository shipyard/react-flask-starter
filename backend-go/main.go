package main
import (
    "fmt"
    "log"
    "net/http"
)

func serverError(w http.ResponseWriter, message string, err error) {
    errMsg := fmt.Sprintf("%s (%v)", message, err)
    log.Println(errMsg)
    http.Error(w, errMsg, http.StatusInternalServerError)
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /gif-search/{id}", gifSearchGet)
    mux.HandleFunc("POST /gif-search", gifSearchCreate)

    http.ListenAndServe(":8081", mux)
}