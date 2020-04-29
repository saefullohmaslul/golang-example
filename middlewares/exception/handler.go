package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// ErrorHandler to handling error
func ErrorHandler(c *gin.Context, err interface{}) {
	res := Exception{}
	mapstructure.Decode(err, &res)

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status": res.Status,
		"flag":   res.Flag,
		"errors": map[string]interface{}{
			"message": res.Errors.Message, "flag": res.Errors.Flag,
		},
	})
	return
}

// Exception type
type Exception struct {
	Status int64       `json:"status"`
	Flag   string      `json:"flag"`
	Errors ErrorDetail `json:"errors"`
}

// ErrorDetail type
type ErrorDetail struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}
