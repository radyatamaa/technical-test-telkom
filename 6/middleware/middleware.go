package middleware

import (
	"strconv"
	"technical-test-telkom/helper/logger"

	"github.com/labstack/echo/v4"
)

const (
	USERNAME = "telkom"
	PASSWORD = "telkom147"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// LOG will handle the LOG middleware
func (m *GoMiddleware) Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		l := logger.L
		l.Info("Accepted")

		next(c)
		l.Info("[" + strconv.Itoa(c.Response().Status) + "] " + "[" + c.Request().Method + "] " + c.Request().Host + c.Request().URL.String())

		l.Info("Closing")
		return nil
	}
}

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
