package store

import (
	"context"
	"time"
)

type Post struct {
	ID        int64      `json:"id"`
	Content   string     `json:"content"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type PostRepository interface {
	Create(context.Context, *Post) error
	GetByID(context.Context, int64) (*Post, error)
	List(context.Context, int, int) ([]*Post, error)
	Update(context.Context, *Post) error
	Delete(context.Context, int64) error
}
