package users

import (
	"context"
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Repository interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
}
