package store

import (
	"database/sql"

	"github.com/rayaadhary/social-go/internal/posts"
)

type Repos struct {
	Posts posts.Repository
}

func NewRepos(db *sql.DB) Repos {
	return Repos{
		Posts: posts.NewSQLCRepo(db),
	}
}
