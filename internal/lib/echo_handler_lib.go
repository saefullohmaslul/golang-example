package lib

import "github.com/labstack/echo/v4"

type EchoHandler struct {
	Echo *echo.Echo
}

func NewEchoHandler() EchoHandler {
	e := echo.New()
	return EchoHandler{Echo: e}
}
