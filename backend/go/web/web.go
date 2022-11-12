package web

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/dietb/react-flask-starter/backend/go/storage"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type searcher interface {
	QueueSearch(context.Context, []string) (*uuid.UUID, error)
}

type datastore interface {
	Get(context.Context, uuid.UUID) (*storage.GetResponse, error)
}

// Webservice ...is not just a clever name
type Webservice struct {
	search searcher
	store  datastore
}

func New(s searcher, store datastore) *Webservice {
	return &Webservice{
		search: s,
		store:  store,
	}
}

func (s *Webservice) Start(ctx context.Context, port string) error {
	r := chi.NewRouter()
	r.Post("/api/gif-search", s.CreateSearch)
	r.Get("/api/gif-search/{key}", s.GetResults)

	srv := http.Server{
		ReadHeaderTimeout: 2 * time.Second,
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           r,
		BaseContext:       func(_ net.Listener) context.Context { return ctx },
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()
	return srv.Shutdown(ctx)
}
