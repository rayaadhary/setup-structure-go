package store

import (
	"context"
	"database/sql"
	"time"

	sqlcgen "github.com/rayaadhary/social-go/internal/sqlc"
)

type postsRepo struct {
	q *sqlcgen.Queries
}

func NewPostsRepo(db *sql.DB) PostRepository {
	return &postsRepo{q: sqlcgen.New(db)}
}

func (r *postsRepo) List(ctx context.Context, limit, offset int) ([]*Post, error) {
	items, err := r.q.ListPosts(ctx, sqlcgen.ListPostsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	out := make([]*Post, 0, len(items))
	for _, it := range items {
		out = append(out, toStorePost(it))
	}
	return out, nil
}

func (r *postsRepo) Create(ctx context.Context, p *Post) error {
	row, err := r.q.CreatePost(ctx, sqlcgen.CreatePostParams{
		Title:   p.Title,
		Content: p.Content,
	})
	if err != nil {
		return err
	}
	fillStorePost(p, row)
	return nil
}

func (r *postsRepo) GetByID(ctx context.Context, id int64) (*Post, error) {
	it, err := r.q.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}
	return toStorePost(it), nil
}

func (r *postsRepo) Update(ctx context.Context, p *Post) error {
	row, err := r.q.UpdatePost(ctx, sqlcgen.UpdatePostParams{
		Title:   p.Title,
		Content: p.Content,
		ID:      p.ID,
	})
	if err != nil {
		return err
	}
	fillStorePost(p, row)
	return nil
}

func (r *postsRepo) Delete(ctx context.Context, id int64) error {
	return r.q.DeletePost(ctx, id)
}

// helpers

func toStorePost(x sqlcgen.Post) *Post {
	return &Post{
		ID:        x.ID,
		Title:     x.Title,
		Content:   x.Content,
		CreatedAt: x.CreatedAt,
		UpdatedAt: nullTimePtr(x.UpdatedAt),
	}
}

func fillStorePost(dst *Post, x sqlcgen.Post) {
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
