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
	Query       interface{}
	Summary     string
	Description string
	Tags        []string
}

type Param struct {
	Name string
}

type Controller struct {
	Prefix string

	routes []Route
}

func NewController(prefix string) *Controller {
	return &Controller{
		Prefix: prefix,
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

	for _, route := range controller.routes {
		fullPath := controller.Prefix + route.Path

		handler := createGenericHandler(route.Handler)

		var body interface{} = route.Body
		var query interface{} = route.Query

		switch route.Method {
		case GET:
			r.GET(fullPath, DynamicBindingMiddleware(body, query), GlobalPipes(), handler)
		case POST:
			r.POST(fullPath, DynamicBindingMiddleware(body, query), GlobalPipes(), handler)
		case PUT:
			r.PUT(fullPath, DynamicBindingMiddleware(body, query), GlobalPipes(), handler)
		case DELETE:
			r.DELETE(fullPath, DynamicBindingMiddleware(body, query), GlobalPipes(), handler)
		case PATCH:
			r.PATCH(fullPath, DynamicBindingMiddleware(body, query), GlobalPipes(), handler)
		case OPTIONS:
			r.OPTIONS(fullPath, DynamicBindingMiddleware(body, query), GlobalPipes(), handler)
		case HEAD:
			r.HEAD(fullPath, DynamicBindingMiddleware(body, query), GlobalPipes(), handler)
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
