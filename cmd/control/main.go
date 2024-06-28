package main

import (
	_ "gin-swagger-demo/docs/swag"
	control "gin-swagger-demo/internal/control/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Control API
// @version 1.0
// @description Control Full API v1.0
// @contact.name -Full API
// @contact.url http://localhost:8081/swagger/v1/index.html
func main() {
	// New gin router
	router := gin.New()

	// Register api/v1 endpoints
	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
	control.Register(router)

	// Listen and Server in
	_ = router.Run(":8081")
}
