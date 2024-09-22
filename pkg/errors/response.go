package errors

type ErrorResponse struct {
	Code       uint   `json:"code"`
	Message    string `json:"message"`
	StackTrace string `json:"trace,omitempty"`
}

type ErrorRes struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StackTrace string `json:"trace,omitempty"`
}
