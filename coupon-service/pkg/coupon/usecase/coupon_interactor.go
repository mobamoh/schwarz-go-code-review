package usecase

import (
	"context"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/interactor/dto"
)

type CouponInteractor interface {
	Apply(context.Context, dto.ApplyRequest) (*dto.ApplyResponse, error)
	Create(context.Context, int, string, int) any
	List(context.Context) ([]domain.Coupon, error)
}
