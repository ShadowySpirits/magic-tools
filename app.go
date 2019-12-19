package main

import (
	"github.com/ShadowySpirits/magic-tools/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegWebRoutes(e)

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":80"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	switch code {
	case http.StatusNotFound:
		if err := c.Redirect(http.StatusFound, "https://blog.sspirits.top"); err != nil {
			c.Logger().Error(err)
		}
		c.Logger().Error(err)
	default:
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}
