package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("receiver")
		status := c.PostForm("status")
		alerts := c.PostForm("alerts")
		//这种格式只能获取网页提交表单的格式，如果用apipost提交请把header中"Content-Type": "application/x-www-form-urlencoded"

		c.JSON(200, gin.H{
			"message": message,
			"status":  status,
			"alerts":  alerts,
		})
	})
	router.Run(":8080")
}
