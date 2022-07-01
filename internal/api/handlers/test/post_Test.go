package test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trino-network/pay-srv/internal/api"
)

func PostTestRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Test.POST("/test", geTestHandler(s))
}

// Returns the version and build date baked into the binary.
func geTestHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		sig := c.Request().Header.Get("trino-pay-Sig")
		fmt.Println("sigggg gggg: ", sig)

		return nil
	}
}
