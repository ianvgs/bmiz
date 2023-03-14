package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "",
		})
	}
}

func HomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": "",
		})
	}
}
