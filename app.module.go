package main

import (
	"gomodule/libs"
	"gomodule/roles"
	"gomodule/users"
)

func AppModule() *libs.Module {
	module := libs.NewModule()

	module.Imports(users.UserModule(), roles.RoleModule())

	return module
}
