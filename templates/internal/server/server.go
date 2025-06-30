package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from LupettoGo starter 👋",
		})
	})

	r.Run(":8080")
}
