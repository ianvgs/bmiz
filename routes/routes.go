package routes

import (
	"bmiz/controllers"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()

	if os.Getenv("GO_ENV") == "production" {

		// Get the absolute path to the executable
		executablePath, err := os.Executable()
		if err != nil {
			log.Fatalf("Error getting executable path: %s", err)
		}
		executableDir := filepath.Dir(executablePath)
		r.Static("/assets", filepath.Join(executableDir, "assets"))
		r.LoadHTMLGlob(filepath.Join(executableDir, "templates/*.html"))
		public := r.Group("/")
		publicRoutes(public)
		r.Run(":8080")
		return
	}

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")

	public := r.Group("/")
	publicRoutes(public)
	r.Run(":8080")
}

func publicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.IndexHandler())
	g.GET("/home", controllers.HomeHandler())
}
