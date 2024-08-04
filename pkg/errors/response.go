package errors

type ErrorResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
