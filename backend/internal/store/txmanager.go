package store

import (
	"context"
	"database/sql"

	sqlcgen "github.com/rayaadhary/social-go/internal/sqlc"
)

type TxManager struct {
	db *sql.DB
}

func NewTxManager(db *sql.DB) *TxManager {
	return &TxManager{db: db}
}

func (m *TxManager) Run(ctx context.Context, fn func(*sqlcgen.Queries) error) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	qtx := sqlcgen.New(tx)

	if err := fn(qtx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}
