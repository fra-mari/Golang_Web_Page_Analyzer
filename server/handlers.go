package server

import (
	"github.com/gin-gonic/gin"
	"home24/analyzer"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	r.POST("/analyze", func(c *gin.Context) {
		url := c.PostForm("url")
		result, err := analyzer.AnalyzePage(url)
		if err != nil {
			result.ErrorMessage = err.Error()
		}
		c.HTML(http.StatusOK, "result.html", result)
	})
}
