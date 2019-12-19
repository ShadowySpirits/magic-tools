package controller

import (
	"bytes"
	"github.com/ShadowySpirits/magic-tools/xdxls2csv"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Xdxls2csv(c echo.Context) error {
	location, _ := time.LoadLocation("Asia/Shanghai")
	date, err := time.ParseInLocation(xdxls2csv.TimeLayout, c.FormValue("date"), location)
	if err != nil {
		c.String(http.StatusBadRequest, "日期格式错误")
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "文件错误")
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "文件错误")
	}
	buffer := bytes.NewBuffer(make([]byte, 0))
	xdxls2csv.ParseXlsFromFile(file, date, buffer)
	c.Response().Header().Add("content-disposition", `attachment;filename="课表.csv"`)
	return c.Stream(http.StatusOK, "text/csv", buffer)
}
