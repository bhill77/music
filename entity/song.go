package entity

import (
	"time"

	"github.com/thedevsaddam/govalidator"
)

type Song struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Artist     string    `json:"artist"`
	Cover      string    `json:"cover"`
	File       string    `json:"file"`
	PlayCounts int       `json:"play_counts"`
	CreatedAt  time.Time `json:"created_at"`
}

func (s *Song) Rules() govalidator.MapData {
	return govalidator.MapData{
		"title":  []string{"required"},
		"artist": []string{"required"},
	}
}
