package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --------------- CONTROLLER DECORATOR ----------------------

// Route is a struct that defines a route
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// Controller is a struct that will register routes
type Controller struct {
	Prefix string
	Routes []Route
}

// NewController creates a new controller with a given route prefix
func NewController(prefix string, routes []Route) *Controller {
	return &Controller{
		Prefix: prefix,
		Routes: routes,
	}
}

// RegisterController registers the controller's routes with the gin router
func RegisterController(r *gin.Engine, controller *Controller) {
	for _, route := range controller.Routes {
		fullPath := controller.Prefix + route.Path
		switch route.Method {
		case "GET":
			r.GET(fullPath, route.Handler)
		case "POST":
			r.POST(fullPath, route.Handler)
			// Add more methods as needed
		}
	}
}

// --------------- MODULE DECORATOR --------------------------

// Module represents a module with controllers and providers
type Module struct {
	Controllers []*Controller
	Providers   []interface{} // You can define providers as services or dependencies here
}

// NewModule creates a new module with given controllers and providers
func NewModule(controllers []*Controller, providers []interface{}) *Module {
	return &Module{
		Controllers: controllers,
		Providers:   providers,
	}
}

// RegisterModule registers all controllers in the module
func RegisterModule(r *gin.Engine, module *Module) {
	for _, controller := range module.Controllers {
		RegisterController(r, controller)
	}
}

// --------------- MAIN APP ---------------------------

// Sample route handlers
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Users",
	})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create User",
	})
}

func main() {
	// Create a gin router
	r := gin.Default()

	// Define routes for UserController
	userController := NewController("/users", []Route{
		{Method: "GET", Path: "/", Handler: getUsers},
		{Method: "POST", Path: "/", Handler: createUser},
	})

	// Create a module with UserController
	userModule := NewModule([]*Controller{userController}, nil)

	// Register the module
	RegisterModule(r, userModule)

	// Start the gin server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
