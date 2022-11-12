package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// GetResults fetches the results of a search event
func (s *Webservice) GetResults(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("missing search key"))
		return
	}

	uid, err := uuid.Parse(key)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	dat, err := s.store.Get(r.Context(), uid)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}

	writeResponse(w, dat)
}
