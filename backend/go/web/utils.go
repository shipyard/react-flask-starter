package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func writeErr(w http.ResponseWriter, errcode int, err error) {
	w.WriteHeader(errcode)
	_, _ = w.Write([]byte(err.Error()))
}

func getRequest(w http.ResponseWriter, r *http.Request, target any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return err
	}

	if len(body) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("invalid request"))
		return err
	}

	if err := json.Unmarshal(body, target); err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return err
	}
	return nil
}

func writeResponse(w http.ResponseWriter, data any) {
	j, err := json.Marshal(data)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(j); err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}
}
