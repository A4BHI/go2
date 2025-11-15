package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(ctx *gin.Context) {

		ctx.HTML(200, "index.html", nil)
	})

	r.GET("/pfp", func(ctx *gin.Context) {
		name := "Abhijith"
		ctx.HTML(200, "pfp.html", gin.H{
			"Data": name,
		})
	})
	r.Run()
}
