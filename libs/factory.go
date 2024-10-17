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

func DynamicBindingMiddleware(body interface{}, query interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if body != nil {
			var bodyPtr interface{}
			if reflect.TypeOf(body).Kind() == reflect.Ptr {
				bodyPtr = reflect.New(reflect.TypeOf(body).Elem()).Interface()
			} else {
				bodyPtr = reflect.New(reflect.TypeOf(body)).Interface()
			}

			if err := c.ShouldBindJSON(bodyPtr); err != nil {
				c.JSON(400, gin.H{
					"error":  "Invalid request payload",
					"detail": err.Error(),
				})
				c.Abort()
				return
			}

			c.Set(boundBody, bodyPtr)
		}

		if query != nil {
			var queryPtr interface{}
			if reflect.TypeOf(query).Kind() == reflect.Ptr {
				queryPtr = reflect.New(reflect.TypeOf(query).Elem()).Interface()
			} else {
				queryPtr = reflect.New(reflect.TypeOf(query)).Interface()
			}

			if err := c.ShouldBindQuery(queryPtr); err != nil {
				c.JSON(400, gin.H{
					"error":  "Invalid query parameters",
					"detail": err.Error(),
				})
				c.Abort()
				return
			}

			c.Set(boundQuery, queryPtr)
		}

		c.Next()
	}
}

func Body(c *Context, obj interface{}) error {
	boundData, exists := c.Get(boundBody)
	if !exists {
		return fmt.Errorf("no bound object found in context")
	}

	boundValue := reflect.ValueOf(boundData)
	if boundValue.Kind() != reflect.Ptr {
		return fmt.Errorf("bound object must be a pointer")
	}

	if reflect.TypeOf(boundData).Elem() != reflect.TypeOf(obj).Elem() {
		return fmt.Errorf("bound object type does not match expected type")
	}

	reflect.ValueOf(obj).Elem().Set(boundValue.Elem())

	return nil
}

func Query(c *Context, obj interface{}) error {
	boundData, exists := c.Get(boundQuery)
	if !exists {
		return fmt.Errorf("no bound object found in context")
	}

	boundValue := reflect.ValueOf(boundData)
	if boundValue.Kind() != reflect.Ptr {
		return fmt.Errorf("bound object must be a pointer")
	}

	if reflect.TypeOf(boundData).Elem() != reflect.TypeOf(obj).Elem() {
		return fmt.Errorf("bound object type does not match expected type")
	}

	reflect.ValueOf(obj).Elem().Set(boundValue.Elem())

	return nil
}

type PipeFunc func(obj interface{}) error

func (c *CustomEngine) UseGlobalPipes(pipes ...PipeFunc) {
	globalPipes = append(globalPipes, pipes...)
}

func GlobalPipes() gin.HandlerFunc {
	return func(c *gin.Context) {
		boundBody, bodyExists := c.Get(boundBody)
		if bodyExists {
			for _, pipe := range globalPipes {
				if err := pipe(boundBody); err != nil {
					c.JSON(400, gin.H{
						"error":  "Processing failed on body",
						"detail": err.Error(),
					})
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}
