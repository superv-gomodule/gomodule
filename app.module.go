package main

import (
	"gomodule/libs"
	"gomodule/users"
)

func AppModule() *libs.Module {
	return libs.NewModule(
		nil,
		nil,
		[]*libs.Module{users.UserModule()},
	)
}
