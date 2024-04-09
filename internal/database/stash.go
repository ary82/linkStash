package database

import (
	"context"
	"time"
)

type Stash struct {
	Title      *string   `json:"title"`
	Body       *string   `json:"body"`
	Points     int       `json:"points"`
	Created_at time.Time `json:"created_at"`
}

func (database *DB) GetPublicStashes() ([]*Stash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	query := `select title, body, points, created_at from stashes where is_public = true`
	rows, err := database.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stashArr := []*Stash{}
	for rows.Next() {
		stash := new(Stash)
		err := rows.Scan(&stash.Title, &stash.Body, &stash.Points, &stash.Created_at)
		if err != nil {
			return nil, err
		}
		stashArr = append(stashArr, stash)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return stashArr, nil
}
