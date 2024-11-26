package main

import (
	"code/zh_blog/handlers"

	"github.com/gin-gonic/gin"
)

func registerRouter(g *gin.Engine) {
	g.GET("/ping", handlers.Ping)

	{
		_api := g.Group("/api", _api_mws)
		{
			_zh_blog := _api.Group("/zh_blog", _zh_blog_mws)
			{
				_zh_blog.GET("/", handlers.Index)
				_zh_blog.GET("/homepage", handlers.Index)
				_zh_blog.GET("/login", handlers.LoginIn)
				_notes := _zh_blog.Group("/notes", _notes_mws)
				_notes.GET("/get/:id", _get_notes_mws, handlers.GetNotes)
				_notes.GET("/list", _list_notes_mws, handlers.ListNotes)
			}
		}
	}
}

func _api_mws(c *gin.Context) {
	return
}

func _zh_blog_mws(c *gin.Context) {
	return
}

func _notes_mws(c *gin.Context) {
	return
}

func _get_notes_mws(c *gin.Context) {
	return
}

func _list_notes_mws(c *gin.Context) {
	return
}
