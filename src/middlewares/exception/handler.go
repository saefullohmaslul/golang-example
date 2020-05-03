package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// ErrorHandler to handling error
func ErrorHandler(c *gin.Context, err interface{}) {
	res := Exception{}
	if err := mapstructure.Decode(err, &res); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusInternalServerError,
			"flag":   "INTERNAL_SERVER_ERROR",
			"errors": map[string]interface{}{
				"message": "An error occurred on our server", "flag": "ERROR_MAP_TO_STRUCT",
			},
		})
		return
	}

	if res.Status == 200 {
		c.AbortWithStatusJSON(res.Status, gin.H{
			"status":  res.Status,
			"message": res.Message,
			"errors": map[string]interface{}{
				"message": res.Errors.Message, "flag": res.Errors.Flag,
			},
		})
		return
	}

	c.AbortWithStatusJSON(res.Status, gin.H{
		"status": res.Status,
		"flag":   res.Flag,
		"errors": map[string]interface{}{
			"message": res.Errors.Message, "flag": res.Errors.Flag,
		},
	})
}

// Exception type
type Exception struct {
	Status  int         `json:"status"`
	Flag    string      `json:"flag"`
	Errors  ErrorDetail `json:"errors"`
	Message string      `json:"message"`
}

// ErrorDetail type
type ErrorDetail struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}
