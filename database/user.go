package database

import (
	"context"
	"log"
	"strings"
	"time"

	"google.golang.org/api/idtoken"
)

type User struct {
	ID         int       `json:"id"`
	Email      *string   `json:"email"`
	Username   *string   `json:"username"`
	Name       *string   `json:"name"`
	Points     int       `json:"points"`
	Picture    *string   `json:"picture"`
	Created_at time.Time `json:"created_at"`
}

// This function takes in the google idtoken payload as the input
// and inserts user into the database if they don't exist.
func (database *DB) InsertUserByPayload(payload *idtoken.Payload) error {

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Get username from email
	email := payload.Claims["email"].(string)
	username := strings.Split(email, "@")[0]

	insertUserQuery := `
  INSERT INTO users (username, name, email, picture)
  VALUES ($1, $2, $3, $4)
  ON CONFLICT DO UPDATE
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
func (database *DB) GetUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	getUserQuery := `SELECT * from users WHERE email = $1`
	user := new(User)
	row := database.Pool.QueryRow(ctx, getUserQuery, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Name,
		&user.Points,
		&user.Picture,
		&user.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
