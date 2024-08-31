package success

type SuccessResponse struct {
	Code    uint   `json:"code"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
