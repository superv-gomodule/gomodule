package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	PATCH   HTTPMethod = "PATCH"
	OPTIONS HTTPMethod = "OPTIONS"
	HEAD    HTTPMethod = "HEAD"
)

type Route struct {
	Method  HTTPMethod
	Path    string
	Handler func(*gin.Context) interface{}
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
		case GET:
			r.GET(fullPath, func(c *gin.Context) {
				result := route.Handler(c)
				c.JSON(http.StatusOK, result)
			})
		case POST:
			r.POST(fullPath, func(c *gin.Context) {
				result := route.Handler(c)
				c.JSON(http.StatusOK, result)
			})
		case PUT:
			r.PUT(fullPath, func(c *gin.Context) {
				result := route.Handler(c)
				c.JSON(http.StatusOK, result)
			})
		case DELETE:
			r.DELETE(fullPath, func(c *gin.Context) {
				result := route.Handler(c)
				c.JSON(http.StatusOK, result)
			})
		case PATCH:
			r.PATCH(fullPath, func(c *gin.Context) {
				result := route.Handler(c)
				c.JSON(http.StatusOK, result)
			})
		case OPTIONS:
			r.OPTIONS(fullPath, func(c *gin.Context) {
				result := route.Handler(c)
				c.JSON(http.StatusOK, result)
			})
		case HEAD:
			r.HEAD(fullPath, func(c *gin.Context) {
				result := route.Handler(c)
				c.JSON(http.StatusOK, result)
			})
		default:
			panic("Unsupported HTTP method: " + string(route.Method))

		}
	}
}
