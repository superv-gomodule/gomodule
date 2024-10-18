package libs

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Module struct {
	controllers []*Controller
	providers   map[reflect.Type]interface{}
	imports     []*Module
}

func NewModule() *Module {
	return &Module{
		controllers: []*Controller{},
		providers:   make(map[reflect.Type]interface{}),
	}
}

func RegisterModule(r *gin.Engine, module *Module) {
	for _, importedModule := range module.imports {
		RegisterModule(r, importedModule)
	}

	for _, controller := range module.controllers {
		RegisterController(r, controller)
	}

}

func (m *Module) Controllers(controller ...*Controller) {
	m.controllers = append(m.controllers, controller...)
}

func (m *Module) Imports(importedModule ...*Module) {
	m.imports = append(m.imports, importedModule...)
}

func (m *Module) Providers(provider ...interface{}) {
	for _, p := range provider {
		RegisterProvider(m, p)
	}
}

func RegisterProvider(m *Module, provider interface{}) {
	m.providers[reflect.TypeOf(provider)] = provider
}

func (m *Module) Inject(target interface{}) error {
	targetType := reflect.TypeOf(target).Elem()

	for providerType, provider := range m.providers {
		if providerType.Implements(targetType) || providerType == targetType {
			reflect.ValueOf(target).Elem().Set(reflect.ValueOf(provider))
			return nil
		}
	}

	for _, importedModule := range m.imports {
		err := importedModule.Inject(target)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("no provider found for type %s", targetType)
}
