package domain

type CouponRequestData struct {
	Code           string
	Discount       int
	MinBasketValue int
}
type Coupon struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"minBasketValue"`
}

func NewCoupon(couponData CouponRequestData) *Coupon {
	return &Coupon{
		Code:           couponData.Code,
		Discount:       couponData.Discount,
		MinBasketValue: couponData.MinBasketValue,
	}
}

func validate(couponData CouponRequestData) error {
	panic("todo: perform some validation")
}
