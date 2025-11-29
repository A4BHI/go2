package main

import (
	"context"
	"fmt"
	"go2/addurls"
	"go2/db"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Links struct {
	ID    int
	SHORT string
	ORG   string
}

func main() {

	db.Connect()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/images", "./images")
	data := []Links{}

	r.GET("/", func(ctx *gin.Context) {
		var id int
		var org string
		var short string
		rows, err := db.Pool.Query(context.Background(), "select id,short_code,org_link from links where username=$1", "admin")
		if err != nil {
			gin.LoggerWithWriter(os.Stdout, "Error in executing")
		}

		for rows.Next() {
			rows.Scan(&id, &org, &short)
			data = append(data, Links{ID: id, SHORT: short, ORG: org})
		}
		fmt.Println(data)
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

		ctx.Redirect(302, "https://"+orglink)
	})
	r.Run()

}
