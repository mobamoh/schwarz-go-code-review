package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/pkg/coupon/usecase"
)

type CouponHandler struct {
	service usecase.CouponInteractor
}

func NewCouponHTTP(svc usecase.CouponInteractor) *CouponHandler {
	return &CouponHandler{service: svc}
}

func (h *CouponHandler) HandleApply(c *gin.Context) error {

	return nil
}

func (h *CouponHandler) HandleCreate(c *gin.Context) error {
	return nil
}
func (h *CouponHandler) HandleList(c *gin.Context) error {
	return nil
}
