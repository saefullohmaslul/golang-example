package exception

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrorHandler to handling error
func ErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, _ := err.(*echo.HTTPError)

		response := Exception{}
		if err := json.Unmarshal([]byte(err.Error()), &response); err != nil {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())

			c.JSON(http.StatusInternalServerError, report)
			c.Logger().Error(report)
			return
		}

		if response.Status == 0 {
			c.JSON(http.StatusInternalServerError, response)
			c.Logger().Error(response)
			return
		}

		c.JSON(response.Status, response)
		c.Logger().Error(report)
	}
}
