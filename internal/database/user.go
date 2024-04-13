package database

import (
	"context"
	"log"
	"strings"
	"time"

	"google.golang.org/api/idtoken"
)

type User struct {
	ID       int     `json:"id"`
	Username string `json:"username"`
	Picture  *string `json:"picture"`
}

type UserDetail struct {
	ID            int       `json:"id"`
	Username      string   `json:"username"`
	Stars         int       `json:"stars"`
	Picture       *string   `json:"picture"`
	Created_at    time.Time `json:"created_at"`
	PublicStashes []*Stash  `json:"public_stashes"`
}

func (database *DB) GetUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	query := `
  SELECT id, username, picture 
  FROM users 
  WHERE email = $1
  `
	user := new(User)
	err := database.Pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Picture,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// This function takes in the google idtoken payload as the input
// and inserts user into the database if they don't exist.
func (database *DB) UpsertUserByPayload(payload *idtoken.Payload) error {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Get username from email
	email := payload.Claims["email"].(string)
	username := strings.Split(email, "@")[0]

	insertUserQuery := `
  INSERT INTO users (username, name, email, picture)
  VALUES ($1, $2, $3, $4)
  ON CONFLICT (email)
  DO UPDATE SET
  name=EXCLUDED.name, picture=EXCLUDED.picture
  `

	resp, err := database.Pool.Exec(
		ctx,
		insertUserQuery,
		username,
		payload.Claims["name"].(string),
		email, payload.Claims["picture"].(string),
	)
	if err != nil {
		return err
	}
	log.Printf("%+v\n", resp)
	return nil
}

// Returns User by email
func (database *DB) GetUserProfile(id int) (*UserDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Populate user's info
	getUserQuery := `
  SELECT id, username, (
    SELECT COUNT(1)
    FROM stashes INNER JOIN stars
    ON stashes.id = stars.stash_id
    WHERE stashes.owner_id = $1
  ), picture, created_at
  FROM users WHERE id = $1
  `

	user := new(UserDetail)

	row := database.Pool.QueryRow(ctx, getUserQuery, id)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Stars,
		&user.Picture,
		&user.Created_at,
	)
	if err != nil {
		return nil, err
	}

	// Populate user's public stashes
	getStashQuery := `
  SELECT username, users.id, title, body,
  stashes.id, stashes.created_at,
  (SELECT count(1) FROM stars WHERE stashes.id = stars.stash_id)
  FROM stashes INNER JOIN users
  ON stashes.owner_id = users.id
  WHERE stashes.is_public = true AND stashes.owner_id = $1
  `
	rows, err := database.Pool.Query(ctx, getStashQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
		user.PublicStashes = append(user.PublicStashes, stash)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return user, nil
}
