package main

import (
	"fmt"

	"gomodule/libs"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router
	r := gin.Default()

	// Register the user module
	libs.RegisterModule(r, AppModule())

	// Start the gin server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
