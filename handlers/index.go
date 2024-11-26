package handlers

import (
	"net/http"

	"code/zh_blog/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := map[string]any{}
	res["notes"] = []model.Note{
		{Title: "hello", TitleDesc: "hello desc"},
		{Title: "hello1", TitleDesc: "hello desc1"},
	}
	res["labels"] = []model.NoteLabel{
		{ID: 1, Name: "读后感"},
		{ID: 2, Name: "电影观后感"},
	}
	c.HTML(http.StatusOK, "index.html", res)
}
