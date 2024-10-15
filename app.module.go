package main

import (
	"gomodule/libs"
	"gomodule/users"
)

func AppModule() *libs.CtrlModule {
	return libs.Module(nil, nil, []*libs.CtrlModule{users.UserModule()})
}
