package web

import (
	"net/http"

	"github.com/google/uuid"
)

type searchrequest struct {
	SearchKeys []string `json:"search_keys,omitempty"`
}

type searchresponse struct {
	Key *uuid.UUID `json:"key,omitempty"`
}

// CreateSearch creates a search event
func (s *Webservice) CreateSearch(w http.ResponseWriter, r *http.Request) {
	req := searchrequest{}
	if err := getRequest(w, r, &req); err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}

	id, err := s.search.QueueSearch(r.Context(), req.SearchKeys)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err)
		return
	}

	resp := searchresponse{
		Key: id,
	}

	writeResponse(w, resp)
}
