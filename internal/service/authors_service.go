package service

import (
	"context"

	"github.com/go-starter/internal/db"
)

type AuthorService interface {
	GetAllAuthors(ctx context.Context) ([]db.Author, error)
	CreateAuthor(ctx context.Context, name string) (db.Author, error)
	GetAuthor(ctx context.Context, id int64) (db.Author, error)
	DeleteAuthor(ctx context.Context, id int64) error
}

type authorService struct {
	q *db.Queries
}

func NewAuthorService(q *db.Queries) AuthorService {
	return &authorService{q: q}
}

func (s *authorService) GetAllAuthors(ctx context.Context) ([]db.Author, error) {
	return s.q.ListAuthors(ctx)
}

func (s *authorService) CreateAuthor(ctx context.Context, name string) (db.Author, error) {
	return s.q.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: name,
	})
}

func (s *authorService) GetAuthor(ctx context.Context, id int64) (db.Author, error) {
	return s.q.GetAuthor(ctx, id)
}

func (s *authorService) DeleteAuthor(ctx context.Context, id int64) error {
	return s.q.DeleteAuthor(ctx, id)
}
