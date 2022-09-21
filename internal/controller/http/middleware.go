package http

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	excludeGzipPaths = []string{"docs", "metrics"}
)

var pingSkipper = func(ctx echo.Context) bool {
	return strings.HasSuffix(ctx.Request().URL.Path, "/ping")
}

type Middleware func(*Server)

func (s *Server) Middlewares(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		middleware(s)
	}
}

func WithLogger() Middleware {
	return func(s *Server) {
		s.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			CustomTimeFormat: "2006-01-02T15:04:05.1483386-00:00",
			Format: "[${time_custom}][INFO] [method=${method}] [uri=${uri}] [status=${status}]" +
				"[origin=${header:X-Application-ID}]\n",
			Skipper: pingSkipper,
		}))
	}
}

func WithRecover() Middleware {
	return func(s *Server) {
		s.server.Use(middleware.Recover())
	}
}

func WithGzip() Middleware {
	return func(s *Server) {
		s.server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				for _, path := range excludeGzipPaths {
					if strings.Contains(c.Request().URL.Path, path) {
						return true
					}
				}

				return false
			},
		}))
	}
}

func WithCORS() Middleware {
	return func(s *Server) {
		DefaultCORSConfig := middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders,
				echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccessControlAllowMethods, echo.HeaderAuthorization},
			AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
			AllowCredentials: true,
		}
		s.server.Use(middleware.CORSWithConfig(DefaultCORSConfig))
	}
}
