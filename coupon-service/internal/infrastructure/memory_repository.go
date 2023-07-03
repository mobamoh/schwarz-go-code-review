package infrastructure

import (
	"fmt"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
)

type MemoryRepository struct {
	entries map[string]domain.Coupon
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (r *MemoryRepository) FindByCode(code string) (*domain.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("Coupon not found")
	}
	return &coupon, nil
}

func (r *MemoryRepository) Save(coupon *domain.Coupon) error {
	//r.entries[coupon.Code] = coupon
	return nil
}
