package users

import (
	"gomodule/libs"
)

func UserModule() *libs.CtrlModule {
	return libs.Module([]*libs.CtrlController{UserController()}, nil, nil)
}
