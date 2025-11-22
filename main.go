package main

import (
	"context"
	"fmt"
	"go2/addurls"
	"go2/db"
	"net/http"

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

	r.GET("/pfp", func(ctx *gin.Context) {
		// name := "Abhijith"
		ctx.HTML(200, "pfp.html", gin.H{
			"Name": "Abhijith",
		})
	})

	r.POST("/addurl", func(ctx *gin.Context) {
		url := ctx.Request.FormValue("url")
		fmt.Println(url)
		code := addurls.Url(url)
		var id int
		db.Pool.QueryRow(context.Background(), "select id from links where short_code=$1", code).Scan(&id)

		data := gin.H{
			"ID":    id,
			"SHORT": "localhost:8080/" + code,
			"ORG":   url,
		}
		ctx.HTML(http.StatusOK, "row.html", data)

	})

	r.GET("/:code", func(ctx *gin.Context) {

		code := ctx.Param("code")
		var orglink string

		row := db.Pool.QueryRow(context.Background(), "select org_link from links where short_code = $1", code)
		err := row.Scan(&orglink)

		if err != nil {
			fmt.Println("Error occured in the redirecting function: ", err)
			ctx.String(404, "Invalid or Expired link")
			return
		}

		ctx.Redirect(302, orglink)
	})
	r.Run()

}
