package usecase

import (
	"context"
	"github/architecture/internal/entity"
)

// ===============================
// ========= REPOSITORY ==========
// ===============================
type CustomerRepo interface {
	Create(ctx context.Context, m *entity.Customer) error
	Update(ctx context.Context, m *entity.Customer) error
	Get(ctx context.Context, m map[string]string) (*entity.Customer, error)
	List(ctx context.Context, m map[string]string) ([]*entity.Customer, error)
}

type FileRepo interface {
	Get(ctx context.Context, guid string) (*entity.File, error)
	List(ctx context.Context, limit, offset uint64, filter entity.Parameter) ([]*entity.File, error)
	Create(ctx context.Context, m *entity.File) error
	Update(ctx context.Context, m *entity.File) error
}

// ===============================
// =========== WEB API ===========
// ===============================
