package libs

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type CtrlController struct {
	Prefix string
	Routes []Route
}

func Controller(prefix string, routes []Route) *CtrlController {
	return &CtrlController{
		Prefix: prefix,
		Routes: routes,
	}
}

func RegisterController(r *gin.Engine, controller *CtrlController) {
	for _, route := range controller.Routes {
		fullPath := controller.Prefix + route.Path
		switch route.Method {
		case "GET":
			r.GET(fullPath, route.Handler)
		case "POST":
			r.POST(fullPath, route.Handler)
		}
	}
}
