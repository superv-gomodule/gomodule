package users

import (
	"gomodule/libs"
)

func UserModule() *libs.Module {
	return libs.NewModule(
		[]*libs.Controller{UserController()},
		nil,
		nil,
	)
}
