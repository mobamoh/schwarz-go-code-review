package dto

type CreateRequest struct {
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"minBasketValue"`
}

type CreateResponse struct {
	CouponId string `json:"couponId"`
	Error    error  `json:"error"`
}
