package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	// Stash operations
	CheckOwner(userId int, stashId int) (bool, error)
	CheckStashPublic(id int) (bool, error)
	GetPublicStashes() ([]*Stash, error)
	GetStashDetailed(id int) (*StashDetail, error)

	//User operations
	GetUserByEmail(email string) (*User, error)
	GetUserProfile(id int) (*UserDetail, error)
	UpsertUser(username string, name string, email string, pictue string) error
}

type Postgres struct {
	Pool *pgxpool.Pool
}

// Takes in a connection string and returns a DB interface,
// which is a pointer to the Postgres Database
func NewPostgresDB(connStr string) (DB, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &Postgres{Pool: pool}, nil
}
