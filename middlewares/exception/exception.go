package exception

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Exception to return error exception
type Exception struct {
	Status int    `json:"status"`
	Flag   string `json:"flag"`
	Errors Errors `json:"errors"`
}

// Errors to format errors detail
type Errors struct {
	Flag    string `json:"flag"`
	Message string `json:"message"`
}

// BadRequest handler
func BadRequest(message string, flag string) error {
	e := Exception{
		Status: http.StatusBadRequest,
		Flag:   "BAD_REQUEST",
		Errors: Errors{
			Flag:    flag,
			Message: message,
		},
	}

	return errorHandler(e)
}

func errorHandler(e Exception) error {
	out, err := json.Marshal(&e)
	if err != nil {
		return errors.New("Failed parse response error")
	}

	return errors.New(string(out))
}
