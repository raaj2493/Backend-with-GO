package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router * gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message" : "TODO API is Running",
			"Status": "Success",
		})
	})
	router.Run(":3000")
}