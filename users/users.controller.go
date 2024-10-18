package users

import (
	"gomodule/libs"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"required,email"`
}

type UserParams struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required,min=1"`
}

type UserController struct {
	userService UserService
}

func NewUserController(module *libs.Module) *libs.Controller {
	userController := &UserController{}
	module.Inject(&userController.userService)

	controller := libs.NewController("/users")
	controller.GET(libs.Route{
		Path:        "/",
		Handler:     userController.getUsers,
		Summary:     "Get user information",
		Description: "Retrieve user details based on the provided user ID.",
		Tags:        []string{"Users"},
	})

	controller.GET(libs.Route{
		Path:    "/:id",
		Handler: userController.getUser,
		Summary: "Get user by ID",
		Tags:    []string{"Users"},
	})

	controller.POST(libs.Route{
		Path:    "/",
		Handler: userController.createUser,
		Summary: "Create a new user",
		Tags:    []string{"Users"},
	})

	return controller
}

// UserController godoc
// @Summary Get user by ID
// @Description Get user details
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {string} string "Get User {id}"
// @Router /users/{id} [get]
func (uc *UserController) getUser(c *libs.Context) interface{} {
	return uc.userService.FindAll()

}

// getUsers godoc
// @Summary Get all users
// @Description Get a list of users
// @Tags users
// @Success 200 {object} map[string]interface{}
// @Router /users/ [get]
func (uc *UserController) getUsers(c *libs.Context) interface{} {
	var params UserParams
	libs.Query(c, &params)

	return map[string]interface{}{
		"message": "User created",
		"user":    params,
	}
}

// createUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Success 200 {string} string "Create User"
// @Router /users/ [post]
// @Param name body string true "Name"
// @Param email body string true "Email"
func (uc *UserController) createUser(c *libs.Context) interface{} {
	var user User
	libs.Body(c, &user)

	return map[string]interface{}{
		"message": "User created",
		"user":    user,
	}
}
