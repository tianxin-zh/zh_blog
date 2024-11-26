package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListNotes(c *gin.Context) {
	res := map[string]any{}
	c.HTML(http.StatusOK, "index.html", res)
}
