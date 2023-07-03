package domain

import "context"

const (
	CouponCreated = "COUPON_CREATED" // Produced
	CouponApplied = "COUPON_APPLIED" // Produced
)

type CouponSAGAEventBus interface {
	Applied(ctx context.Context, service string) error
	Failed(ctx context.Context, service, msg string) error
	Created(ctx context.Context, coupon Coupon) error
}
