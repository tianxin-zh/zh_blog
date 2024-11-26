package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	id := c.Param("id")
	switch id {
	case "1":
		c.String(http.StatusOK, fmt.Sprintf("我的id是%v,名字叫小豪，年龄是3岁", id))
	case "2":
		c.String(http.StatusOK, fmt.Sprintf("我的id是%v,名字叫小明，年龄是5岁", id))
	default:
		c.String(http.StatusOK, fmt.Sprintf("你的id是%v,但我没找到你的信息", id))
	}
}
