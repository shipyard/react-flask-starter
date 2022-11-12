package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
)

func (s *Storage) QueueSearch(ctx context.Context, args []string) (*uuid.UUID, error) {
	u, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	t, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = t.Rollback()
	}()

	for _, val := range args {
		url, err := s.search.Search(ctx, val)
		if err != nil {
			return nil, err
		}

		if _, err := t.Exec(
			`
			INSERT INTO giphy_data 
			(uuid, giphy_url)
			VALUES
			($1, $2)
			`,
			u,
			url,
		); err != nil {
			return nil, err
		}
		go s.uploadFile(val, url)
	}

	if err := t.Commit(); err != nil {
		return nil, err
	}

	return &u, nil
}

type uploadResponse struct {
	Message string `json:"message"`
}

func (s *Storage) uploadFile(queryparam, url string) {
	r, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	var (
		buf    = new(bytes.Buffer)
		writer = multipart.NewWriter(buf)
	)

	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s.gif", queryparam))
	if err != nil {
		log.Println(err)
		return
	}

	if _, err := io.Copy(part, r.Body); err != nil {
		log.Println(err)
		return
	}
	_ = writer.Close()

	req, err := http.NewRequest(http.MethodPost, s.uploadendpoint, buf)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := s.httpclient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	uplr := uploadResponse{}
	_ = json.Unmarshal(b, &uplr)

	log.Printf("response body from upload is %s\n", uplr.Message)
}
