package store

import "database/sql"

type Storage struct {
	Posts PostRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: NewPostsRepo(db),
	}
}
