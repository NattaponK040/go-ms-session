package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Check(e *echo.Echo, checks []HealthCheck) {
	for _, c := range checks {
		e.POST(c.Route, c.Process)
		c.StatusCode = http.StatusOK
	}
}
