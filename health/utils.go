package health

import "github.com/labstack/echo/v4"

type HealthCheck struct {
	Process    func(c echo.Context) error
	StatusCode int
	Route      string
}
