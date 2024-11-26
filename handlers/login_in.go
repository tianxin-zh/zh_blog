package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginIn(c *gin.Context) {
	res := map[string]any{}
	c.HTML(http.StatusOK, "login.html", res)
}
