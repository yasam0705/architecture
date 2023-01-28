package usecase

import (
	"context"
	"github/architecture/internal/entity"
	"time"
)

type customer struct {
	ctxTimeout time.Duration
	repo       CustomerRepo
}

type CustomerService interface {
	Create(ctx context.Context, m *entity.Customer) error
	Update(ctx context.Context, m *entity.Customer) error
	Get(ctx context.Context, m map[string]string) (*entity.Customer, error)
	List(ctx context.Context, m map[string]string) ([]*entity.Customer, error)
}

func NewCustomer(ctxTimeout time.Duration, repo CustomerRepo) *customer {
	return &customer{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (c *customer) Create(ctx context.Context, m *entity.Customer) error {
	return c.repo.Create(ctx, m)
}

func (c *customer) Update(ctx context.Context, m *entity.Customer) error {
	return c.repo.Update(ctx, m)
}

func (c *customer) Get(ctx context.Context, m map[string]string) (*entity.Customer, error) {
	return c.repo.Get(ctx, m)
}

func (c *customer) List(ctx context.Context, m map[string]string) ([]*entity.Customer, error) {
	return c.repo.List(ctx, m)
}
