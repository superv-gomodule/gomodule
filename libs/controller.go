package libs

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type Controller struct {
	Prefix string
	Routes []Route
}

func NewController(prefix string, routes []Route) *Controller {
	return &Controller{
		Prefix: prefix,
		Routes: routes,
	}
}

func RegisterController(r *gin.Engine, controller *Controller) {
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
