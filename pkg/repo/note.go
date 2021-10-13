package repo

import (
	"context"
	"notes/pkg/db"
	"notes/pkg/models"
)

type Note struct {
	adapter *db.Storage
}

func New(adapter *db.Storage) *Note {
	return &Note{adapter: adapter}
}

func (u *Note) GetNotes(ctx context.Context) ([]models.Note, error) {
	notes, err := u.adapter.Get(ctx)
	if err != nil {
		return nil, err
	}
	for i := range notes {
		data := []rune(notes[i].Text)
		if len(data) > 30 {
			notes[i].Text = string(data[0:27]) + "..."
		}
	}
	return notes, err
}
