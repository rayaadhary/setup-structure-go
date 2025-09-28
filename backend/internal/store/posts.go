package store

import (
	"context"
	"database/sql"
	"time"
)

type Post struct {
	ID        int64      `json:"id"`
	Content   string     `json:"content"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type PostStore struct {
	db *sql.DB
}

type PostRepository interface {
	Create(context.Context, *Post) error
	GetByID(context.Context, int64) (*Post, error)
	List(context.Context, int, int) ([]*Post, error)
	Update(context.Context, *Post) error
	Delete(context.Context, int64) error
}

func (s *PostStore) List(ctx context.Context, limit, offset int) ([]*Post, error) {
	query := `SELECT id, content, title, created_at, updated_at 
		FROM posts ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(
			&p.ID,
			&p.Content,
			&p.Title,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, &p)
	}

	return posts, nil
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title)
	VALUES ($1, $2)
	RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(ctx, query,
		post.Content,
		post.Title,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	return err
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
	SELECT id, content, title, created_at, updated_at
	FROM posts
	WHERE id = $1
	`

	var p Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Content,
		&p.Title,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = now()
		WHERE id = $3
		RETURNING updated_at
	`

	return s.db.QueryRowContext(ctx, query,
		post.Title,
		post.Content,
		post.ID,
	).Scan(&post.UpdatedAt)
}

func (s *PostStore) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	return err
}
