package interactor

import (
	"context"
	"fmt"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/interactor/dto"
)

type Coupon struct {
	repository domain.CouponRepository
}

func NewCoupon(repo domain.CouponRepository) *Coupon {
	return &Coupon{
		repository: repo,
	}
}

func (c *Coupon) Apply(ctx context.Context, data dto.ApplyRequest) (*dto.ApplyResponse, error) {
	coupon, err := c.repository.FindByCode(ctx, data.CouponCode)
	if err != nil {
		return nil, fmt.Errorf("no valid coupon with code %v found", data.CouponCode)
	}

	cart, err := c.repository.FindCartByCode(ctx, data.CartId)
	if err != nil {
		return nil, fmt.Errorf("no basket with id %v found", data.CartId)
	}

	// if we assume there is a cart-service,
	// this should trigger and event which will be handled by Cart Aggregate
	// This should be handled using SAGA
	if cart.Value > 0 {
		cart.AppliedDiscount = coupon.Discount
		cart.ApplicationSuccessful = true
		// Call save cart
	} else {
		return nil, fmt.Errorf("failed to apply coupon %v to cart %v", data.CouponCode, data.CartId)
	}

	res := dto.ApplyResponse{
		Basket: cart,
		Err:    nil,
	}
	return &res, nil
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
