package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET orderItems",
		})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET orderItem",
		})
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET orderItems by order",
		})
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {

}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST orderItem",
		})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT orderItem",
		})
	}
}
