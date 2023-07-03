package dto

type CreateRequest struct {
}

type CreateResponse struct {
	Err error `json:"-"`
}
