package model

type File struct {
	ID          int    `json:"id"`
	Filename    string `json:"filename"`
	UserID      int    `json:"user_id"`
	IsProcessed int    `json:"is_processed"`
}