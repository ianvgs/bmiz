package routes

import (
	"bmiz/controllers"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Erro ao carregar path")
	}

	r := gin.Default()

	r.Static("/assets", dir+"/assets")
	r.LoadHTMLGlob(dir + "/templates/*.html")

	public := r.Group("/")
	publicRoutes(public)
	r.Run(":8080")
}

func publicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.IndexHandler())
	g.GET("/home", controllers.HomeHandler())
}
