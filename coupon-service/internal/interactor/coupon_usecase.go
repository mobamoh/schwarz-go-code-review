package interactor

import (
	"context"
	"fmt"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/interactor/dto"
)

// Coupon Domain Service
type Coupon struct {
	repository domain.CouponRepository
}

func NewCoupon(repo domain.CouponRepository) *Coupon {
	return &Coupon{
		repository: repo,
	}
}

func (c *Coupon) Apply(ctx context.Context, data dto.ApplyRequest) *dto.ApplyResponse {
	res := dto.ApplyResponse{}
	coupon, err := c.repository.FindByCode(ctx, data.CouponCode)
	if err != nil {
		res.Basket = nil
		res.Error = fmt.Errorf("no valid coupon with code %v found", data.CouponCode)
		return &res
	}

	cart, err := c.repository.FindCartByCode(ctx, data.CartId)
	if err != nil {
		res.Basket = nil
		res.Error = fmt.Errorf("no basket with id %v found", data.CartId)
		return &res
	}

	// if we assume there is a cart-service,
	// this should trigger and event which will be handled by Cart Aggregate
	// This should be handled using SAGA
	if cart.Value > 0 {
		cart.AppliedDiscount = coupon.Discount
		cart.ApplicationSuccessful = true
		// Call save cart
	} else {
		res.Basket = nil
		res.Error = fmt.Errorf("failed to apply coupon %v to cart %v", data.CouponCode, data.CartId)
		return &res
	}

	res.Basket = cart
	res.Error = nil
	return &res
}

func (c *Coupon) Create(ctx context.Context, data dto.CreateRequest) *dto.CreateResponse {
	res := dto.CreateResponse{}
	coupon := domain.NewCoupon(data.Code, data.Discount, data.MinBasketValue)
	savedCoupon, err := c.repository.Save(ctx, coupon)
	if err != nil {
		res.Error = fmt.Errorf("failed to create coupon with code %v", data.Code)
		return &res
	}
	// Basically it's not a good idea to expose internal id,
	// I assume this call is used by internal BE to create coupon
	res.CouponId = savedCoupon.ID
	return &res
}

func (c *Coupon) ListAll(ctx context.Context) *dto.ListResponse {
	res := dto.ListResponse{}
	coupons, err := c.repository.List(ctx)
	if err != nil {
		res.Error = fmt.Errorf("failed to fetch coupons %v", err)
		return &res
	}

	for _, coupon := range coupons {
		coup := dto.CouponData{
			Code:           coupon.Code,
			Discount:       coupon.Discount,
			MinBasketValue: coupon.MinBasketValue,
		}
		res.Coupons = append(res.Coupons, &coup)
	}
	return &res
}
