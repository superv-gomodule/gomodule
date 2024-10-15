package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Users",
	})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create User",
	})
}
