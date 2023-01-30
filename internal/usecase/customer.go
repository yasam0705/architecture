package usecase

import (
	"context"
	"github/architecture/internal/entity"
	"time"

	"github.com/google/uuid"
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

func (c *customer) beforeCreate(m *entity.Customer) error {
	m.GUID = uuid.New().String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	return nil
}
func (c *customer) beforeUpdate(m *entity.Customer) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (c *customer) Create(ctx context.Context, m *entity.Customer) error {
	c.beforeCreate(m)
	return c.repo.Create(ctx, m)
}

func (c *customer) Update(ctx context.Context, m *entity.Customer) error {
	c.beforeUpdate(m)
	return c.repo.Update(ctx, m)
}

func (c *customer) Get(ctx context.Context, m map[string]string) (*entity.Customer, error) {
	return c.repo.Get(ctx, m)
}

func (c *customer) List(ctx context.Context, m map[string]string) ([]*entity.Customer, error) {
	return c.repo.List(ctx, m)
}
