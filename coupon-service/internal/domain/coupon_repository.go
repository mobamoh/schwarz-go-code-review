package domain

import "context"

type CouponRepository interface {
	Save(context.Context, *Coupon) (*Coupon, error)
	FindByCode(context.Context, string) (*Coupon, error)
	List(context.Context) ([]*Coupon, error)
	FindCartByCode(context.Context, string) (*Basket, error)
}
