package response

// Success -> general response success format
type Success struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []Error     `json:"errors"`
}

type Error struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}
