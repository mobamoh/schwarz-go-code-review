package usecase

import (
	"context"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/interactor/dto"
)

type CouponInteractor interface {
	Apply(context.Context, dto.ApplyRequest) *dto.ApplyResponse
	Create(context.Context, dto.CreateRequest) *dto.CreateResponse
	ListAll(context.Context) *dto.ListResponse
}
