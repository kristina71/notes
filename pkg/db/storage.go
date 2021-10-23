package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"notes/pkg/models"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
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
		log.Println(err)
		return nil, err
	}

	notes := []models.Note{}
	err = s.db.Select(&notes, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return notes, nil
}

func (s *Storage) GetById(ctx context.Context, note models.Note) (models.Note, error) {
	fmt.Println(note.Id)
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("id", "text", "created_time", "updated_time").From(tableName).Where(squirrel.Eq{"id": note.Id}).ToSql()
	if err != nil {
		log.Println(err)
		return models.Note{}, err
	}

	note_ := models.Note{}
	err = s.db.Get(&note_, query, args...)

	if err == sql.ErrNoRows {
		return models.Note{}, err
	}

	if err != nil {
		return models.Note{}, err
	}

	return note_, nil
}
