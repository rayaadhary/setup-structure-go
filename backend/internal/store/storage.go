package store

import (
	"database/sql"

	"github.com/rayaadhary/social-go/internal/posts"
	"github.com/rayaadhary/social-go/internal/users"
)

type Repos struct {
	Posts posts.Repository
	Users users.Repository
}

func NewRepos(db *sql.DB) Repos {
	return Repos{
		Posts: posts.NewSQLCRepo(db),
		Users: users.NewSQLCRepo(db),
	}
}
