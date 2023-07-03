package usecase

import (
	"context"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
)

type CouponInteractor interface {
	ApplyCoupon(context.Context, domain.Basket, string) (*domain.Basket, error)
	CreateCoupon(context.Context, int, string, int) any
	GetCoupons(context.Context, []string) ([]domain.Coupon, error)
}
