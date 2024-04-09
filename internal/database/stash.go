package database

import (
	"context"
	"time"
)

type Stash struct {
	ID         int       `json:"id"`
	Author     *string   `json:"author"`
	Title      *string   `json:"title"`
	Body       *string   `json:"body"`
	Stars      int       `json:"stars"`
	Created_at time.Time `json:"created_at"`
}

type StashDetail struct {
	ID         int        `json:"id"`
	Author     *string    `json:"author"`
	Title      *string    `json:"title"`
	Body       *string    `json:"body"`
	Stars      int        `json:"stars"`
	IsPublic   bool       `json:"is_public"`
	Created_at time.Time  `json:"created_at"`
	Links      []*Link    `json:"links"`
	Comments   []*Comment `json:"comments"`
}

func (database *DB) GetPublicStashes() ([]*Stash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	query := `
  SELECT username, title, body, stashes.id, stashes.created_at, (
    SELECT count(1) FROM stars WHERE stashes.id = stars.stash_id
  )
  FROM stashes INNER JOIN users ON stashes.owner_id = users.id
  WHERE stashes.is_public = true`
	rows, err := database.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stashArr := []*Stash{}
	for rows.Next() {
		stash := new(Stash)
		err := rows.Scan(&stash.Author, &stash.Title, &stash.Body, &stash.ID, &stash.Created_at, &stash.Stars)
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
