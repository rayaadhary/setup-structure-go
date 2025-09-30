package posts

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("post not found")
)

type Post struct {
	ID        int64      `json:"id"`
	Content   string     `json:"content"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ListParams struct {
	Limit  int
	Offset int
}

type Repository interface {
	Create(ctx context.Context, p *Post) error
	GetByID(ctx context.Context, id int64) (*Post, error)
	List(ctx context.Context, p ListParams) ([]*Post, error)
	Update(ctx context.Context, p *Post) error
	Delete(ctx context.Context, id int64) error
}
