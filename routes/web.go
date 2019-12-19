package routes

import (
	"github.com/ShadowySpirits/magic-tools/app/controller"
	"github.com/labstack/echo/v4"
)

func RegWebRoutes(e *echo.Echo) {
	xdxls2csv(e)
}

func xdxls2csv(e *echo.Echo) {
	e.Static("/static", "public/assets")

	e.File("/", "public/index.html")
	e.POST("/xls2csv", controller.Xdxls2csv)
}
