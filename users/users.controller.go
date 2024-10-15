package users

import (
	"gomodule/libs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController() *libs.CtrlController {
	userController := libs.Controller("/users", []libs.Route{
		{Method: "GET", Path: "/", Handler: getUsers},
		{Method: "POST", Path: "/", Handler: createUser},
	})

	return userController
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Users 4",
	})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create User",
	})
}
