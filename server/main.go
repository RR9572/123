package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./view")
	r.GET("ping", func(context *gin.Context) {
		id := context.Query("id")
		context.HTML(200, "index", gin.H{
			"code": "200",
			"msg":  id,
		})
	})
}
