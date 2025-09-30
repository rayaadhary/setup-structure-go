package service

import (
	"context"
	"errors"

	"github.com/rayaadhary/social-go/internal/posts"
)

type PostService struct {
	repo posts.Repository
}

func NewPostService(repo posts.Repository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) ListPosts(ctx context.Context, limit, offset int) ([]*posts.Post, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	params := posts.ListParams{
		Limit:  limit,
		Offset: offset,
	}

	return s.repo.List(ctx, params)
}

func (s *PostService) CreatePost(ctx context.Context, p *posts.Post) error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	if len(p.Content) < 5 {
		return errors.New("content too short")
	}

	return s.repo.Create(ctx, p)
}

func (s *PostService) GetPost(ctx context.Context, id int64) (*posts.Post, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *PostService) UpdatePost(ctx context.Context, p *posts.Post) error {
	if p.ID <= 0 {
		return errors.New("invalid id")
	}

	if p.Title == "" {
		return errors.New("title is required")
	}

	if len(p.Content) < 5 {
		return errors.New("content too short")
	}

	return s.repo.Update(ctx, p)
}

func (s *PostService) DeletePost(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}
