package infrastructure

import (
	"context"
	"fmt"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
)

type MemoryRepository struct {
	entries map[string]*domain.Coupon
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		entries: make(map[string]*domain.Coupon),
	}
}

func (r *MemoryRepository) Save(ctx context.Context, coupon *domain.Coupon) (*domain.Coupon, error) {
	r.entries[coupon.ID] = coupon
	return coupon, nil
}

func (r *MemoryRepository) FindByCode(ctx context.Context, code string) (*domain.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("coupon not found")
	}
	return coupon, nil
}

func (r *MemoryRepository) List(ctx context.Context) ([]*domain.Coupon, error) {
	return values(r.entries), nil
}

func (r *MemoryRepository) FindCartByCode(ctx context.Context, code string) (*domain.Basket, error) {
	basket := domain.Basket{ID: "", Value: 11, AppliedDiscount: 11}
	return &basket, nil
}

func values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
