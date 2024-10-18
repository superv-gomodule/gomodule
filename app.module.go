package main

import (
	"gomodule/libs"
	"gomodule/users"
)

func AppModule() *libs.Module {
	module := libs.NewModule()

	module.Imports(users.UserModule())

	return module
}
