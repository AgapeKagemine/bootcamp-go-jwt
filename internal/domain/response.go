package domain

type Response struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
	Payload any    `json:"data,omitempty"`
}
