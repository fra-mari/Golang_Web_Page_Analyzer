package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	//r.POST("/analyze", func(c *gin.Context) {
	//	url := c.PostForm("url")
	//	result, err := analyzePage(url)
	//	if err != nil {
	//		result.ErrorMessage = err.Error()
	//	}
	//	c.HTML(http.StatusOK, "result.html", result)
	//})

	r.Run(":8080")
}
