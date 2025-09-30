package posts

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

func NewSQLCRepo(db *sql.DB) Repository {
	return &SQLCRepo{q: sqlcgen.New(db)}
}

func (r *SQLCRepo) List(ctx context.Context, p ListParams) ([]*Post, error) {
	rows, err := r.q.ListPosts(ctx, sqlcgen.ListPostsParams{
		Limit:  int32(p.Limit),
		Offset: int32(p.Offset),
	})
	if err != nil {
		return nil, err
	}
	out := make([]*Post, 0, len(rows))
	for _, row := range rows {
		out = append(out, mapRow(row))
	}
	return out, nil
}

func (r *SQLCRepo) Create(ctx context.Context, p *Post) error {
	row, err := r.q.CreatePost(ctx, sqlcgen.CreatePostParams{
		Title:   p.Title,
		Content: p.Content,
	})
	if err != nil {
		return err
	}
	fill(p, row)
	return nil
}

func (r *SQLCRepo) GetByID(ctx context.Context, id int64) (*Post, error) {
	row, err := r.q.GetPost(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return mapRow(row), nil
}

func (r *SQLCRepo) Update(ctx context.Context, p *Post) error {
	row, err := r.q.UpdatePost(ctx, sqlcgen.UpdatePostParams{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}
	fill(p, row)
	return nil
}

func (r *SQLCRepo) Delete(ctx context.Context, id int64) error {
	if err := r.q.DeletePost(ctx, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

// mapping helpers
func mapRow(x sqlcgen.Post) *Post {
	return &Post{
		ID:        x.ID,
		Title:     x.Title,
		Content:   x.Content,
		CreatedAt: x.CreatedAt,
		UpdatedAt: nullTimePtr(x.UpdatedAt),
	}
}

func fill(dst *Post, x sqlcgen.Post) {
	dst.ID = x.ID
	dst.Title = x.Title
	dst.Content = x.Content
	dst.CreatedAt = x.CreatedAt
	dst.UpdatedAt = nullTimePtr(x.UpdatedAt)
}

func nullTimePtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
