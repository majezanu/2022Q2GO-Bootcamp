package custom_error

type ErrorResponse struct {
	Error string `json:"error" example:"message"`
	Code  int    `json:"code" example:"400"`
}
