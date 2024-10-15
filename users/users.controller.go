package users

import (
	"gomodule/libs"

	"github.com/gin-gonic/gin"
)

func UserController() *libs.CtrlController {
	userController := libs.Controller("/users", []libs.Route{
		{Method: libs.GET, Path: "/", Handler: getUsers},
		{Method: libs.GET, Path: "/:id", Handler: getUser},
		{Method: libs.POST, Path: "/", Handler: createUser},
	})

	return userController
}

func getUser(c *gin.Context) interface{} {
	id := c.Param("id")
	return "Get User " + id
}

func getUsers(c *gin.Context) interface{} {
	return map[string]interface{}{
		"data": 1,
	}
}

func createUser(c *gin.Context) interface{} {
	return "Create User"
}
