package database

import (
	"context"
	"time"
)

type Stash struct {
	ID         int       `json:"id"`
	Author     string    `json:"author"`
	AuthorId   int       `json:"author_id"`
	Title      string    `json:"title"`
	Body       *string   `json:"body"` // Optional(Can be nil)
	Stars      int       `json:"stars"`
	Created_at time.Time `json:"created_at"`
}

type StashDetail struct {
	// Embed Stash
	Stash
	IsPublic bool       `json:"is_public"`
	Links    []*Link    `json:"links"`
	Comments []*Comment `json:"comments"`
}

func (pg *Postgres) GetPublicStashes() ([]*Stash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	query := `
  SELECT username, users.id, title, body,
  stashes.id, stashes.created_at,
  (SELECT count(1) FROM stars WHERE stashes.id = stars.stash_id)
  FROM stashes INNER JOIN users
  ON stashes.owner_id = users.id
  WHERE stashes.is_public = true
  `
	rows, err := pg.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stashArr := []*Stash{}
	for rows.Next() {
		stash := new(Stash)
		err := rows.Scan(
			&stash.Author,
			&stash.AuthorId,
			&stash.Title,
			&stash.Body,
			&stash.ID,
			&stash.Created_at,
			&stash.Stars,
		)
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

func (pg *Postgres) GetPublicStashesUser(userId int) ([]*Stash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	getStashQuery := `
  SELECT username, users.id, title, body,
  stashes.id, stashes.created_at,
  (SELECT count(1) FROM stars WHERE stashes.id = stars.stash_id)
  FROM stashes INNER JOIN users
  ON stashes.owner_id = users.id
  WHERE stashes.is_public = true AND stashes.owner_id = $1
  `
	rows, err := pg.Pool.Query(ctx, getStashQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stashArr := []*Stash{}

	for rows.Next() {
		stash := new(Stash)
		err := rows.Scan(
			&stash.Author,
			&stash.AuthorId,
			&stash.Title,
			&stash.Body,
			&stash.ID,
			&stash.Created_at,
			&stash.Stars,
		)
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

func (pg *Postgres) GetUserStashes(userId int) ([]*Stash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	getStashQuery := `
  SELECT username, users.id, title, body,
  stashes.id, stashes.created_at,
  (SELECT count(1) FROM stars WHERE stashes.id = stars.stash_id)
  FROM stashes INNER JOIN users
  ON stashes.owner_id = users.id
  WHERE stashes.owner_id = $1
  `
	rows, err := pg.Pool.Query(ctx, getStashQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stashArr := []*Stash{}

	for rows.Next() {
		stash := new(Stash)
		err := rows.Scan(
			&stash.Author,
			&stash.AuthorId,
			&stash.Title,
			&stash.Body,
			&stash.ID,
			&stash.Created_at,
			&stash.Stars,
		)
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

func (pg *Postgres) GetStashDetailed(stashId int) (*StashDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Populate Stash details
	stashQuery := `
  SELECT username, users.id, title, body,
  stashes.id, stashes.created_at, stashes.is_public,
  (SELECT count(1) FROM stars WHERE stashes.id = stars.stash_id)
  FROM stashes INNER JOIN users
  ON stashes.owner_id = users.id
  WHERE stashes.id = $1
  `

	stash := new(StashDetail)
	err := pg.Pool.QueryRow(ctx, stashQuery, stashId).Scan(
		&stash.Author,
		&stash.AuthorId,
		&stash.Title,
		&stash.Body,
		&stash.ID,
		&stash.Created_at,
		&stash.IsPublic,
		&stash.Stars,
	)
	if err != nil {
		return nil, err
	}

	// Populate Links
	stash.Links, err = pg.GetLinks(stashId)
	if err != nil {
		return nil, err
	}

	// Populate comments
	stash.Comments, err = pg.GetComments(stashId)
	if err != nil {
		return nil, err
	}

	return stash, nil
}

func (pg *Postgres) CheckStashPublic(stashId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	query := `
  SELECT is_public FROM stashes
  WHERE id = $1
  `
	var isPublic bool
	err := pg.Pool.QueryRow(ctx, query, stashId).Scan(&isPublic)
	return isPublic, err
}

func (pg *Postgres) CheckOwner(userId int, stashId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	query := `
  SELECT owner_id FROM stashes
  WHERE id = $1
  `
	// Initialize as -1 to avoid zero case
	ownerId := -1
	err := pg.Pool.QueryRow(ctx, query, stashId).Scan(&ownerId)
	if err != nil {
		return false, err
	}
	if ownerId == userId {
		return true, nil
	}
	return false, nil
}
