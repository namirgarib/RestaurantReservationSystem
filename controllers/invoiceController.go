package controllers

import "github.com/gin-gonic/gin"

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET invoices",
		})
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET invoice",
		})
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST invoice",
		})
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT invoice",
		})
	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE invoice",
		})
	}
}
