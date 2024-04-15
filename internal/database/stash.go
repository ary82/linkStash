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

func (pg *Postgres) CheckStashPublic(id int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	query := `
  SELECT is_public FROM stashes
  WHERE id = $1
  `
	var isPublic bool
	err := pg.Pool.QueryRow(ctx, query, id).Scan(&isPublic)
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

func (pg *Postgres) GetStashDetailed(id int) (*StashDetail, error) {
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
	err := pg.Pool.QueryRow(ctx, stashQuery, id).Scan(
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
	linksQuery := `
  SELECT id, url, comment FROM links
  WHERE stash_id = $1
  `
	rows, err := pg.Pool.Query(ctx, linksQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
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
		stash.Links = append(stash.Links, link)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	// Populate comments
	commentsQuery := `
  SELECT comments.id, users.id, username, body, comments.created_at
  FROM comments INNER JOIN users
  ON comments.author = users.id
  WHERE stash_id = $1
  `
	rows, err = pg.Pool.Query(ctx, commentsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := new(Comment)
		err := rows.Scan(
			&comment.ID,
			&comment.AuthorId,
			&comment.Author,
			&comment.Body,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		stash.Comments = append(stash.Comments, comment)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return stash, nil
}
