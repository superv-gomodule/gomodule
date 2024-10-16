package libs

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CustomEngine struct {
	*gin.Engine
}

func Create(module *CtrlModule) *CustomEngine {
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
