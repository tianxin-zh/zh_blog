package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNotes(c echo.Context) error {
	id := c.Param("id")
	switch id {
	case "1":
		return c.String(http.StatusOK, fmt.Sprintf("我的id是%v,名字叫小豪，年龄是3岁", id))
	case "2":
		return c.String(http.StatusOK, fmt.Sprintf("我的id是%v,名字叫小明，年龄是5岁", id))
	default:
		return c.String(http.StatusOK, fmt.Sprintf("你的id是%v,但我没找到你的信息", id))
	}
}
