package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/interactor/dto"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/pkg/coupon/usecase"
	"net/http"
)

type CouponHandler struct {
	service usecase.CouponInteractor
}

func NewCouponHandler(svc usecase.CouponInteractor) CouponHandler {
	return CouponHandler{service: svc}
}
func (h *CouponHandler) HandleApply(c *gin.Context) {
	req := dto.ApplyRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	res := h.service.Apply(c, req)
	c.JSON(statusCode(res.Error), res)
}

func (h *CouponHandler) HandleCreate(c *gin.Context) {
	req := dto.CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	res := h.service.Create(c, req)
	c.JSON(statusCode(res.Error), res)
}
func (h *CouponHandler) HandleList(c *gin.Context) {
	res := h.service.ListAll(c)
	c.JSON(statusCode(res.Error), res)
}

func statusCode(err error) int {
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
