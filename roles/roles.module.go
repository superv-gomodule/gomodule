package roles

import (
	"gomodule/libs"
	"gomodule/users"
)

func RoleModule() *libs.Module {
	module := libs.NewModule()

	module.Imports(users.UserModule())
	module.Controllers(NewRoleController(module))

	return module
}
