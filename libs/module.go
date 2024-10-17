package libs

import "github.com/gin-gonic/gin"

type Module struct {
	Controllers []*Controller
	Providers   []interface{}
	Imports     []*Module
}

func NewModule(controllers []*Controller, providers []interface{}, imports []*Module) *Module {
	return &Module{
		Controllers: controllers,
		Providers:   providers,
		Imports:     imports,
	}
}

func RegisterModule(r *gin.Engine, module *Module) {
	for _, controller := range module.Controllers {
		RegisterController(r, controller)
	}

	for _, importedModule := range module.Imports {
		RegisterModule(r, importedModule)
	}
}
