package dto

import (
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
)

type ApplyRequest struct {
	CouponCode string
	CartId     string
}

type ApplyResponse struct {
	Basket *domain.Basket `json:"basket"`
	Err    error          `json:"-"`
}
