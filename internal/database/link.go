package database

type Link struct {
	ID      int     `json:"id"`
	Url     string  `json:"url"`
	Comment *string `json:"comment"` // Optional(Can be nil)
}
