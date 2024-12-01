package controllers

import "github.com/gin-gonic/gin"

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET orders",
		})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST order",
		})
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET order",
		})
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT order",
		})
	}
}
