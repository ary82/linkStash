package database

import (
	"context"
	"time"
)

type Link struct {
	ID      int     `json:"id"`
	Url     string  `json:"url"`
	Comment *string `json:"comment"` // Optional(Can be nil)
}

func (pg *Postgres) GetLinks(stashId int) ([]*Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	linksQuery := `
  SELECT id, url, comment FROM links
  WHERE stash_id = $1
  `
	rows, err := pg.Pool.Query(ctx, linksQuery, stashId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := []*Link{}

	for rows.Next() {
		link := new(Link)
		err := rows.Scan(
			&link.ID,
			&link.Url,
			&link.Comment,
		)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return links, nil
}
