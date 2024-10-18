package users

import (
	"gomodule/libs"
)

func UserModule() *libs.Module {
	module := libs.NewModule()

	module.AddController(UserController())

	return module
}
