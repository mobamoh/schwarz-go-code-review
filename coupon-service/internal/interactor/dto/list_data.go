package dto

type ListRequest struct {
}

type ListResponse struct {
	Coupons []*CouponData `json:"coupons"`
	Error   error         `json:"error"`
}

type CouponData struct {
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"minBasketValue"`
}
