package main

import (
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/cmd/server/dep"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/infrastructure"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/internal/interactor"
	"github.com/mobamoh/schwarz-go-code-review/coupon-service/pkg/transport/bind"
)

func main() {
	app, err := BuildAppCompileTime()
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		panic(err)
	}
}

func BuildAppCompileTime() (*dep.App, error) {
	listener, err := dep.NewListener()
	if err != nil {
		return nil, err
	}
	memoryRepository := infrastructure.NewMemoryRepository()
	coupon := interactor.NewCoupon(memoryRepository)
	couponHandler := bind.NewCouponHandler(coupon)
	handler := dep.NewRouter(couponHandler)
	app := dep.NewApp(listener, handler)
	return app, nil
}
