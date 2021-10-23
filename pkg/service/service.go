package service

import (
	"context"
	"fmt"
	"notes/pkg/models"
)

type Repository interface {
	GetNotes(ctx context.Context) ([]models.Note, error)
	GetNote(ctx context.Context, note models.Note) (models.Note, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetNotes(ctx context.Context) ([]models.Note, error) {
	return s.repo.GetNotes(ctx)
}

func (s Service) GetNote(ctx context.Context, note models.Note) (models.Note, error) {
	fmt.Println(note.Id)
	return s.repo.GetNote(ctx, note)
}
