package libs

import (
	"net/http"
	"os"

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

type Context struct {
	*gin.Context
}

type Route struct {
	Method      HTTPMethod
	Path        string
	Handler     func(*Context) interface{}
	Body        interface{}
	Summary     string
	Description string
	Tags        []string
}

type Controller struct {
	Prefix string
	Routes []Route
}

type Param struct {
	Name string
}

func NewController(prefix string, routes []Route) *Controller {
	return &Controller{
		Prefix: prefix,
		Routes: routes,
	}
}

func WrapGinContext(c *gin.Context) *Context {
	return &Context{Context: c}
}

func RegisterController(r *gin.Engine, controller *Controller) {
	file, err := os.OpenFile("docs/docs.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, route := range controller.Routes {
		fullPath := controller.Prefix + route.Path

		handler := createGenericHandler(route.Handler)

		var body interface{} = route.Body

		switch route.Method {
		case GET:
			r.GET(fullPath, DynamicBindingMiddleware(body), handler)
		case POST:
			r.POST(fullPath, DynamicBindingMiddleware(body), GlobalValidationMiddleware(), handler)
		case PUT:
			r.PUT(fullPath, DynamicBindingMiddleware(body), handler)
		case DELETE:
			r.DELETE(fullPath, DynamicBindingMiddleware(body), handler)
		case PATCH:
			r.PATCH(fullPath, DynamicBindingMiddleware(body), handler)
		case OPTIONS:
			r.OPTIONS(fullPath, DynamicBindingMiddleware(body), handler)
		case HEAD:
			r.HEAD(fullPath, DynamicBindingMiddleware(body), handler)
		default:
			panic("Unsupported HTTP method: " + string(route.Method))
		}
	}
}

func createGenericHandler(handler func(*Context) interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := WrapGinContext(c)
		result := handler(ctx)
		c.JSON(http.StatusOK, result)
	}
}
