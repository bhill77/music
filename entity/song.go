package entity

import "time"

type Song struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Artist     string    `json:"artist"`
	Cover      string    `json:"cover"`
	File       string    `json:"file"`
	PlayCounts int       `json:"play_counts"`
	CreatedAt  time.Time `json:"created_at"`
}
