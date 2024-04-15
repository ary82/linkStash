package database

import (
	"context"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	AuthorId  int       `json:"author_id"`
	Author    string    `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (pg *Postgres) GetComments(stashId int) ([]*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	commentsQuery := `
  SELECT comments.id, users.id, username, body, comments.created_at
  FROM comments INNER JOIN users
  ON comments.author = users.id
  WHERE stash_id = $1
  `
	rows, err := pg.Pool.Query(ctx, commentsQuery, stashId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArr := []*Comment{}
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
		commentArr = append(commentArr, comment)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return commentArr, nil
}
