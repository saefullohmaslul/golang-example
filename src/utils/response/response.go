package response

// Success -> general response success format
type Success struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
