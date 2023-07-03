package bind

import (
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/pkg/coupon/usecase"
)

type CouponHandler struct {
	service usecase.CouponInteractor
}

func NewCouponHTTP(svc usecase.CouponInteractor) *CouponHandler {
	return &CouponHandler{service: svc}
}
