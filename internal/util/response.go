package util

type ValidateResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string              `json:"message"`
	Errors  []*ValidateResponse `json:"errors"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
