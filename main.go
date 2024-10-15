package main

import (
	"fmt"

	"gomodule/libs"

	"github.com/gin-gonic/gin"

	_ "gomodule/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User API
// @version 1.0
// @description This is a sample server for managing users.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// Create a gin router
	r := gin.Default()

	// Register the user module
	libs.RegisterModule(r, AppModule())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the gin server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
