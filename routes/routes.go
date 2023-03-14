package routes

import (
	"bmiz/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")

	public := r.Group("/")
	publicRoutes(public)
	r.Run()
}

func publicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.IndexHandler())
}
