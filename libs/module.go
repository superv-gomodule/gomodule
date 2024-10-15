package libs

import "github.com/gin-gonic/gin"

type CtrlModule struct {
	Controllers []*CtrlController
	Providers   []interface{}
	Imports     []*CtrlModule
}

func Module(controllers []*CtrlController, providers []interface{}, imports []*CtrlModule) *CtrlModule {
	return &CtrlModule{
		Controllers: controllers,
		Providers:   providers,
		Imports:     imports,
	}
}

func RegisterModule(r *gin.Engine, module *CtrlModule) {
	for _, controller := range module.Controllers {
		RegisterController(r, controller)
	}

	for _, importedModule := range module.Imports {
		RegisterModule(r, importedModule)
	}
}
