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
	Summary     string
	Description string
	Tags        []string
}

type CtrlController struct {
	Prefix string
	Routes []Route
}

type Param struct {
	Name string
}

func Controller(prefix string, routes []Route) *CtrlController {
	return &CtrlController{
		Prefix: prefix,
		Routes: routes,
	}
}

func WrapGinContext(c *gin.Context) *Context {
	return &Context{Context: c}
}

func RegisterController(r *gin.Engine, controller *CtrlController) {
	file, err := os.OpenFile("docs/docs.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, route := range controller.Routes {
		fullPath := controller.Prefix + route.Path
		// swaggerDoc := fmt.Sprintf(`
		// 	// @Summary %s
		// 	// @Description %s
		// 	// @Tags %s
		// 	// @Router %s [%s]`,
		// 	route.Summary, route.Description,
		// 	strings.Join(route.Tags, ","),
		// 	fullPath,
		// 	route.Method)

		// if _, err := file.WriteString(swaggerDoc + "\n"); err != nil {
		// 	panic(err)
		// }

		switch route.Method {
		case GET:
			r.GET(fullPath, func(c *gin.Context) {
				ctx := WrapGinContext(c)
				result := route.Handler(ctx)
				c.JSON(http.StatusOK, result)
			})
		case POST:
			r.POST(fullPath, func(c *gin.Context) {
				ctx := WrapGinContext(c)
				result := route.Handler(ctx)
				c.JSON(http.StatusOK, result)
			})
		case PUT:
			r.PUT(fullPath, func(c *gin.Context) {
				ctx := WrapGinContext(c)
				result := route.Handler(ctx)
				c.JSON(http.StatusOK, result)
			})
		case DELETE:
			r.DELETE(fullPath, func(c *gin.Context) {
				ctx := WrapGinContext(c)
				result := route.Handler(ctx)
				c.JSON(http.StatusOK, result)
			})
		case PATCH:
			r.PATCH(fullPath, func(c *gin.Context) {
				ctx := WrapGinContext(c)
				result := route.Handler(ctx)
				c.JSON(http.StatusOK, result)
			})
		case OPTIONS:
			r.OPTIONS(fullPath, func(c *gin.Context) {
				ctx := WrapGinContext(c)
				result := route.Handler(ctx)
				c.JSON(http.StatusOK, result)
			})
		case HEAD:
			r.HEAD(fullPath, func(c *gin.Context) {
				ctx := WrapGinContext(c)
				result := route.Handler(ctx)
				c.JSON(http.StatusOK, result)
			})
		default:
			panic("Unsupported HTTP method: " + string(route.Method))
		}
	}
}
