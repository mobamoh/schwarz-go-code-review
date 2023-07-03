package interactor

import (
	"context"
	"fmt"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
)

type Coupon struct {
	repository domain.CouponRepository
}

func NewCoupon(repo domain.CouponRepository) *Coupon {
	return &Coupon{
		repository: repo,
	}
}

func (c *Coupon) Apply(ctx context.Context, basket domain.Basket, code string) (b *domain.Basket, e error) {
	b = &basket
	coupon, err := c.repository.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if b.Value > 0 {
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true
	}
	if b.Value == 0 {
		return
	}

	return nil, fmt.Errorf("Tried to apply discount to negative value")
}

func (c *Coupon) Create(ctx context.Context, couponData domain.CouponRequestData) (*domain.Coupon, error) {
	coupon := domain.NewCoupon(couponData)
	err := c.repository.Save(ctx, coupon)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (c *Coupon) GetAll(ctx context.Context, codes []string) ([]domain.Coupon, error) {
	coupons := make([]domain.Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := c.repository.FindByCode(ctx, code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, idx)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, idx)
			}
		}
		coupons = append(coupons, *coupon)
	}

	return coupons, e
}
