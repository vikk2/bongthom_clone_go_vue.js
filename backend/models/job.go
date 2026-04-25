package models

type Job struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Company string `json:"company"`
	New     bool   `json:"new"`
}
