package domain

import "context"

type CouponRepository interface {
	FindByCode(context.Context, string) (*Coupon, error)
	Save(context.Context, *Coupon) error
}
