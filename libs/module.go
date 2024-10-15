package libs

import "github.com/gin-gonic/gin"

type Module struct {
	Controllers []*Controller
	Providers   []interface{}
}

func NewModule(controllers []*Controller, providers []interface{}) *Module {
	return &Module{
		Controllers: controllers,
		Providers:   providers,
	}
}

func RegisterModule(r *gin.Engine, module *Module) {
	for _, controller := range module.Controllers {
		RegisterController(r, controller)
	}
}
