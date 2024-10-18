package roles

import (
	"gomodule/libs"
	"gomodule/users"
)

type RoleController struct {
	userService users.UserService
}

func NewRoleController(module *libs.Module) *libs.Controller {
	roleController := &RoleController{}
	module.Inject(&roleController.userService)

	controller := libs.NewController("/roles")
	controller.GET(libs.Route{
		Path:        "/",
		Handler:     roleController.getRoles,
		Summary:     "Get role information",
		Description: "Retrieve role details based on the provided role ID.",
		Tags:        []string{"Roles"},
	})

	return controller
}

func (rc *RoleController) getRoles(c *libs.Context) interface{} {
	return rc.userService.FindAll()
}
