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

// ===============================
// =========== WEB API ===========
// ===============================
