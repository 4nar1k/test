package middleware

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
)

// Определяем ключ для хранения echo.Context в context.Context
type contextKey string

const EchoContextKey contextKey = "echo-context"

// Middleware для связывания Echo с context.Context
func AttachEchoContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), EchoContextKey, c)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

// Функция для извлечения echo.Context из context.Context
func GetEchoContext(ctx context.Context) (echo.Context, error) {
	c, ok := ctx.Value(EchoContextKey).(echo.Context)
	if !ok {
		return nil, errors.New("echo.Context not found in context.Context")
	}
	return c, nil
}
