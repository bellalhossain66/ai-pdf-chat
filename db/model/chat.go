package model

type Chat struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	FileId   int    `json:"file_id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}