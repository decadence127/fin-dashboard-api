package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {

	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok!",
		})
	})
	return router
}
