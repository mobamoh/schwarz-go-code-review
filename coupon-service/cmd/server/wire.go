//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/cmd/server/dep"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/domain"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/infrastructure"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/interactor"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/pkg/coupon/usecase"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/pkg/transport/bind"
)

func BuildAppCompileTime() (*dep.App, error) {
	wire.Build(
		wire.Bind(new(domain.CouponRepository), new(*infrastructure.MemoryRepository)),
		infrastructure.NewMemoryRepository,
		wire.Bind(new(usecase.CouponInteractor), new(*interactor.Coupon)),
		interactor.NewCoupon,
		bind.NewCouponHandler,
		dep.NewRouter,
		dep.NewListener,
		dep.NewApp)
	return &dep.App{}, nil
}
