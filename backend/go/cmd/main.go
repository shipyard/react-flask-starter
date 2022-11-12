package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dietb/react-flask-starter/backend/go/search"
	"github.com/dietb/react-flask-starter/backend/go/storage"
	"github.com/dietb/react-flask-starter/backend/go/web"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	giphykey := os.Getenv("GIPHY_API_KEY")
	if giphykey == "" {
		log.Fatal("giphy key does not exist")
	}

	pguri := os.Getenv("DATABASE_URL")
	if pguri == "" {
		log.Fatal("postgres uri does not exist")
	}

	port := os.Getenv("WEB_SERVICE_PORT")
	if port == "" {
		log.Fatal("service port configuration does not exist")
	}

	ue := os.Getenv("FILE_UPLOAD_ENDPOINT")
	if port == "" {
		log.Fatal("file upload endpoint configuration does not exist")
	}

	db, err := sql.Open(
		"pgx",
		pguri,
	)
	if err != nil {
		log.Fatalf("pg conn error: %v", err)
	}

	srch := search.New(giphykey)

	str := storage.New(
		srch,
		db,
		ue,
	)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	w := web.New(str, str)

	if err := w.Start(ctx, port); err != nil {
		log.Println(err)
	}
}
