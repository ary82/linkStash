package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	// Stash operations
	GetPublicStashes() ([]*Stash, error)
	GetPublicStashesUser(userId int) ([]*Stash, error)
	GetStashDetailed(stashId int) (*StashDetail, error)
	CheckOwner(userId int, stashId int) (bool, error)
	CheckStashPublic(stashId int) (bool, error)

	//User operations
	GetUserByEmail(email string) (*User, error)
	GetUserProfile(userId int) (*UserDetail, error)
	UpsertUser(username string, name string, email string, pictue string) error

	// Comment operations
	GetComments(stashId int) ([]*Comment, error)

	//Link operations
	GetLinks(stashId int) ([]*Link, error)
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
