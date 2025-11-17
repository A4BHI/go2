package main

import (
	"fmt"
	"go2/addurls"
	"go2/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/images", "./images")
	data := []gin.H{
		{
			"ID":    1,
			"SHORT": "localhost:8080/a1b2",
			"ORG":   "https://google.com",
		},
		{
			"ID":    2,
			"SHORT": "localhost:8080/x9y8",
			"ORG":   "https://youtube.com",
		},
	}
	r.GET("/", func(ctx *gin.Context) {

		ctx.HTML(200, "index.html", gin.H{
			"Urls": data,
		})
	})

	r.GET("/test", func(ctx *gin.Context) {
		ctx.HTML(200, "test.html", gin.H{
			"Urls": data,
		})
	})

	r.GET("/pfp", func(ctx *gin.Context) {
		// name := "Abhijith"
		ctx.HTML(200, "pfp.html", gin.H{
			"Name": "Abhijith",
		})
	})
	r.POST("/addurl", func(ctx *gin.Context) {
		url := ctx.Request.FormValue("url")
		fmt.Println(url)
		addurls.ShortCode()
	})
	r.Run()

}
