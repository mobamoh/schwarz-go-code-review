package dto

import "github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"

type ListRequest struct {
}

type ListResponse struct {
	Coupons []*domain.Coupon `json:"coupons"`
}
