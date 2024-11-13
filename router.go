package main

import (
	"code/zh_blog/handlers"

	"github.com/labstack/echo/v4"
)

func register(e *echo.Echo) {
	e.GET("/ping", handlers.Ping)

	{
		_api := e.Group("/api", _api_mws()...)
		{
			_zh_blog := _api.Group("/zh_blog", _zh_blog_mws()...)
			{
				_notes := _zh_blog.Group("/notes", _notes_mws()...)
				_notes.GET("/get/:id", handlers.GetNotes, _get_notes_mws()...)
			}
		}
	}
}

func _api_mws() []echo.MiddlewareFunc {
	return nil
}

func _zh_blog_mws() []echo.MiddlewareFunc {
	return nil
}

func _notes_mws() []echo.MiddlewareFunc {
	return nil
}

func _get_notes_mws() []echo.MiddlewareFunc {
	return nil
}
