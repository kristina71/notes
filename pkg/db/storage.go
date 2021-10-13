package db

import (
	"context"
	"notes/pkg/models"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

const (
	tableName = "notes"
)

func (s *Storage) Get(ctx context.Context) ([]models.Note, error) {
	query, _, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("id", "text", "created_time", "updated_time").From(tableName).ToSql()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	notes := []models.Note{}
	err = s.db.Select(&notes, query)

	return notes, errors.WithStack(err)
}
