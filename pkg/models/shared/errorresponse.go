package shared

type ErrorResponse struct {
	Error      *string `json:"error,omitempty"`
	Message    *string `json:"message,omitempty"`
	StatusCode *int64  `json:"statusCode,omitempty"`
}
