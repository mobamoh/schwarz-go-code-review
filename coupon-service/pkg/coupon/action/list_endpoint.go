package action

import "github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"

type ListRequest struct {
}

type ListResponse struct {
	Authors []*domain.Coupon `json:"coupons"`
}
