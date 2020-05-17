package exception

import (
	"github.com/saefullohmaslul/golang-example/src/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// ErrorHandler -> handling error middleware
func ErrorHandler(c *gin.Context, err interface{}) {
	res := response.Response{}

	/**
	 * checking error formatting
	 * if error format does not match then
	 * an error has occurred to the code
	 */
	if err := mapstructure.Decode(err, &res); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": res.Message,
			"data":    nil,
			"errors": []map[string]interface{}{
				{"message": "An error occurred on our server", "flag": "ERROR_MAP_TO_STRUCT"},
			},
		})

		/**
		 * you can send error to stack driver in here
		 */
		return
	}

	/**
	 * format error result
	 */
	c.AbortWithStatusJSON(res.Status, gin.H{
		"status":  res.Status,
		"message": res.Message,
		"data":    nil,
		"errors":  res.Errors,
	})

	/**
	 * you can send error to stack driver in here
	 */
}

// Exception -> struct exception format
type Exception struct {
	Status  int         `json:"status"`
	Flag    string      `json:"flag"`
	Errors  ErrorDetail `json:"errors"`
	Message string      `json:"message"`
}

// ErrorDetail -> struct error detail format
type ErrorDetail struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}
