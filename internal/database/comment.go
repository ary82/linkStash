package database

import "time"

type Comment struct {
	ID        int       `json:"id"`
	AuthorId  int       `json:"author_id"`
	Author    string    `json:"author"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
