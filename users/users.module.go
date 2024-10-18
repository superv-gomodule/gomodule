package users

import (
	"gomodule/libs"
)

func UserModule() *libs.Module {
	module := libs.NewModule()
	userService := NewUserService()

	module.Providers(userService)
	module.Controllers(NewUserController(module))

	return module
}
