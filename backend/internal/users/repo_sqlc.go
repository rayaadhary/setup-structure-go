package users

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sqlcgen "github.com/rayaadhary/social-go/internal/sqlc"
)

type SQLCRepo struct {
	q *sqlcgen.Queries
}

func mapRow(x sqlcgen.User) *User {
	return &User{
		ID:        x.ID,
		Username:  x.Username,
		Password:  x.Password,
		CreatedAt: x.CreatedAt,
		UpdatedAt: nullTimePtr(x.UpdatedAt),
	}
}

func nullTimePtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func NewSQLCRepo(db *sql.DB) Repository {
	return &SQLCRepo{q: sqlcgen.New(db)}
}

func (r *SQLCRepo) GetByUsername(ctx context.Context, username string) (*User, error) {
	row, err := r.q.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("users not found")
		}
		return nil, err
	}
	return mapRow(row), nil
}
