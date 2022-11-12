package storage

import (
	"context"

	"github.com/google/uuid"
)

type GetResponse struct {
	Data []*Data `json:"data,omitempty"`
}

type Data struct {
	URL string `json:"url,omitempty"`
}

func (s *Storage) Get(ctx context.Context, u uuid.UUID) (*GetResponse, error) {
	resp := GetResponse{}
	rw, err := s.db.QueryContext(
		ctx,
		`
		SELECT giphy_url FROM giphy_data WHERE uuid = $1
		ORDER BY id;
		`,
		u.String(),
	)
	if err != nil {
		// FIXME:  not found != something broke
		return nil, err
	}
	defer rw.Close()

	for rw.Next() {
		val := ""
		if err := rw.Scan(&val); err != nil {
			return nil, err
		}
		resp.Data = append(
			resp.Data,
			&Data{
				URL: val,
			},
		)
	}
	return &resp, nil
}
