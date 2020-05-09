package utils

// Response -> general response format
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
