package libs

import "github.com/gin-gonic/gin"

type Module struct {
	Controllers []*Controller
	Providers   []interface{}
	Imports     []*Module
}

func NewModule() *Module {
	return &Module{
		Controllers: []*Controller{},
		Providers:   []interface{}{},
		Imports:     []*Module{},
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

func (m *Module) AddController(controller ...*Controller) {
	m.Controllers = append(m.Controllers, controller...)
}

func (m *Module) AddModule(importedModule ...*Module) {
	m.Imports = append(m.Imports, importedModule...)
}
