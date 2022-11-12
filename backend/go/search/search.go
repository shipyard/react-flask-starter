package search

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	searchurl          = "https://api.giphy.com/v1/gifs/search?api_key=%s&q=%s&limit=1&offset=0&rating=g&lang=en"
	httpTimeoutSeconds = 5
)

type Searcher struct {
	apikey string
	client http.Client
}

func New(key string) *Searcher {
	return &Searcher{
		apikey: key,
		client: http.Client{
			Timeout: time.Second * httpTimeoutSeconds,
		},
	}
}

// Search accepts a list of strings and returns a list of associated urls
func (s *Searcher) Search(ctx context.Context, key string) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(
			searchurl,
			s.apikey,
			key,
		),
		nil,
	)
	if err != nil {
		return "", err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	dat := GiphyResponse{}
	if err := json.Unmarshal(body, &dat); err != nil {
		return "", err
	}

	if len(dat.Data) == 0 {
		return "", fmt.Errorf("no search match found for %s", key)
	}

	return dat.Data[0].Images.Downsized.URL, nil
}
