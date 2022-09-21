package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(server *echo.Echo) {
	server.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}
