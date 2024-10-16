package users

import (
	"gomodule/libs"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func UserController() *libs.CtrlController {
	userController := libs.Controller("/users", []libs.Route{
		{
			Method:      libs.GET,
			Path:        "/",
			Handler:     getUsers,
			Summary:     "Get all users",
			Description: "Get a list of users",
			Tags:        []string{"users"},
		},
		{
			Method:      libs.GET,
			Path:        "/:id",
			Handler:     getUser,
			Summary:     "Get user by ID",
			Description: "Get user details",
			Tags:        []string{"users"},
		},
		{
			Method:      libs.POST,
			Path:        "/",
			Handler:     createUser,
			Summary:     "Create a new user",
			Description: "Create a new user",
			Tags:        []string{"users"},
		},
	})

	return userController
}

// UserController godoc
// @Summary Get user by ID
// @Description Get user details
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {string} string "Get User {id}"
// @Router /users/{id} [get]
func getUser(c *gin.Context) interface{} {
	id := c.Param("id")
	return "Get User " + id
}

// getUsers godoc
// @Summary Get all users
// @Description Get a list of users
// @Tags users
// @Success 200 {object} map[string]interface{}
// @Router /users/ [get]
func getUsers(c *gin.Context) interface{} {
	return map[string]interface{}{
		"data": 1,
	}
}

// createUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Success 200 {string} string "Create User"
// @Router /users/ [post]
func createUser(c *gin.Context) interface{} {
	return "Create User"
}
