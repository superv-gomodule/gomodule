package libs

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type CustomEngine struct {
	*gin.Engine
}

func Create(module *Module) *CustomEngine {
	r := gin.Default()

	RegisterModule(r, module)

	return &CustomEngine{
		Engine: r,
	}
}

func (c *CustomEngine) Listen(port int) {
	addr := fmt.Sprintf(":%d", port)

	if err := c.Engine.Run(addr); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

type ValidatorFunc func(obj interface{}) error

var globalValidator ValidatorFunc

func (c *CustomEngine) UseGlobalValidator(v ValidatorFunc) {
	globalValidator = v
}

func DynamicBindingMiddleware(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if obj == nil {
			c.Next()
			return
		}

		var objPtr interface{}
		if reflect.TypeOf(obj).Kind() == reflect.Ptr {
			objPtr = reflect.New(reflect.TypeOf(obj).Elem()).Interface()
		} else {
			objPtr = reflect.New(reflect.TypeOf(obj)).Interface()
		}

		if err := c.ShouldBindJSON(objPtr); err != nil {
			c.JSON(400, gin.H{
				"error":  "Invalid request payload",
				"detail": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("boundObject", objPtr)
		c.Next()
	}
}

func GlobalValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		boundObject, exists := c.Get("boundObject")
		if !exists {
			c.JSON(400, gin.H{
				"error": "No object found for validation",
			})
			c.Abort()
			return
		}

		if globalValidator != nil {
			if err := globalValidator(boundObject); err != nil {
				c.JSON(400, gin.H{
					"error":  "Validation failed",
					"detail": err.Error(),
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func Body(c *Context, obj interface{}) error {
	boundObject, exists := c.Get("boundObject")
	if !exists {
		return fmt.Errorf("no bound object found in context")
	}

	boundValue := reflect.ValueOf(boundObject)
	if boundValue.Kind() != reflect.Ptr {
		return fmt.Errorf("bound object must be a pointer")
	}

	if reflect.TypeOf(boundObject).Elem() != reflect.TypeOf(obj).Elem() {
		return fmt.Errorf("bound object type does not match expected type")
	}

	reflect.ValueOf(obj).Elem().Set(boundValue.Elem())

	return nil
}
