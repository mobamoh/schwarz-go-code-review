package domain

import "github.com/google/uuid"

type Coupon struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"minBasketValue"`
}

func NewCoupon(code string, discount, minBasketValue int) *Coupon {
	return &Coupon{
		ID:             uuid.NewString(),
		Code:           code,
		Discount:       discount,
		MinBasketValue: minBasketValue,
	}

}
func validate(coupon Coupon) error {
	panic("todo: perform business validation rules")
}
