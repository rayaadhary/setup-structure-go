package service

import (
	"context"
	"errors"

	"github.com/rayaadhary/social-go/internal/users"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo users.Repository
}

func NewUserService(repo users.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Authenticate(ctx context.Context, username, password string) (*users.User, error) {
	u, err := s.repo.GetByUsername(ctx, username)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return u, nil
}
