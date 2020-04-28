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
			c.JSON(response.Status, report)
			c.Logger().Error(report)
			return
		}

		if response.Status == 0 {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			c.JSON(response.Status, report)
			c.Logger().Error(report)
			return
		}

		c.JSON(response.Status, response)
		c.Logger().Error(report)
	}
}
