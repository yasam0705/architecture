package usecase

import (
	"context"
	"github/architecture/internal/entity"
	"time"

	"github.com/google/uuid"
)

type file struct {
	ctxTimeout time.Duration
	repo       FileRepo
}

type FileService interface {
	Create(ctx context.Context, m *entity.File) error
	Update(ctx context.Context, m *entity.File) error
	Get(ctx context.Context, guid string) (*entity.File, error)
	List(ctx context.Context, m map[string]string) ([]*entity.File, error)
}

func NewFile(ctxTimeout time.Duration, repo FileRepo) *file {
	return &file{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (f *file) beforeCreate(m *entity.File) error {
	m.Guid = uuid.New().String()
	m.CreatedAt = time.Now().UTC()
	m.UpdatedAt = time.Now().UTC()
	return nil
}

func (f *file) beforeUpdate(m *entity.File) error {
	m.UpdatedAt = time.Now().UTC()
	return nil
}

func (f *file) Create(ctx context.Context, m *entity.File) error {
	f.beforeCreate(m)
	return f.repo.Create(ctx, m)
}

func (f *file) Update(ctx context.Context, m *entity.File) error {
	f.beforeUpdate(m)
	return f.repo.Update(ctx, m)
}

func (f *file) Get(ctx context.Context, guid string) (*entity.File, error) {
	return f.repo.Get(ctx, guid)
}

func (f *file) List(ctx context.Context, m map[string]string) ([]*entity.File, error) {
	return f.repo.List(ctx, m)
}
