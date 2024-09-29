package models

import "time"

type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required`
	Description string    `json:"description" binding:"required`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
