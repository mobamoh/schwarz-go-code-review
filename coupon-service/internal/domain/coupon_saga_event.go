package domain

import "context"

const (
	CouponCreated = "COUPON_CREATED" // Produced
	CouponApplied = "COUPON_APPLIED" // Produced
)

type AuthorSAGAEventBus interface {
	Applied(ctx context.Context, service string) error
	Failed(ctx context.Context, service, msg string) error
	Created(ctx context.Context, coupon Coupon) error
}
