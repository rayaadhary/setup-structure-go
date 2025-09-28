package service

import (
	"context"
	"errors"

	"github.com/rayaadhary/social-go/internal/store"
)

type PostService struct {
	repo store.PostRepository
}

func NewPostService(repo store.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) ListPosts(ctx context.Context, limit, offset int) ([]*store.Post, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	return s.repo.List(ctx, limit, offset)
}

func (s *PostService) CreatePost(ctx context.Context, p *store.Post) error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	if len(p.Content) < 5 {
		return errors.New("content too short")
	}

	return s.repo.Create(ctx, p)
}

func (s *PostService) GetPost(ctx context.Context, id int64) (*store.Post, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *PostService) UpdatePost(ctx context.Context, p *store.Post) error {
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
