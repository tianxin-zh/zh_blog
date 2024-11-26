package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	r.Static("/static", "templates")
	r.LoadHTMLGlob("templates/**/*")
	registerRouter(r)

	r.Run(":8081")
}
