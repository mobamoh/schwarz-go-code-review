package domain

type Coupon struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"minBasketValue"`
}

func validate(coupon Coupon) error {
	panic("todo: perform business validation rules")
}
