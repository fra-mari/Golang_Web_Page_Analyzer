package server

import (
	"github.com/gin-gonic/gin"
	"home24/analyzer"
	"html/template"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	// load html and css files
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// set the routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	r.POST("/analyze", func(c *gin.Context) {
		url := c.PostForm("url")
		a := analyzer.NewAnalyzer()
		result, err := a.AnalyzePage(url)
		if err != nil {
			result.ErrorMessage = template.HTML(err.Error())
		}
		c.HTML(http.StatusOK, "result.html", result)
	})
}
