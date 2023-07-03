package action

import "github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"

type ApplyRequest struct {
}

type ApplyResponse struct {
	Author *domain.Basket `json:"basket"`
	Err    error          `json:"-"`
}
