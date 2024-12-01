package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all users",
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get a user",
		})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create a user",
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a user",
		})
	}
}

func HashPassword(password string) string {
	return password
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return userPassword == providePassword, "Verify password"
}
