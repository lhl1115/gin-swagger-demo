package main

import (
	_ "gin-swagger-demo/docs/swag"
	control "gin-swagger-demo/internal/control/api"
	v1 "gin-swagger-demo/internal/v1/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Test API
// @version 1.0
// @description Test Full API v1.0
// @contact.name -Control API
// @contact.url http://localhost:8081/swagger/control/index.html
func main() {
	// New gin router
	router := gin.New()

	// Register v1 endpoints
	v1.Register(router)
	router.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))

	// Register control endpoints
	control.Register(router)
	router.GET("/swagger/control/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("control")))

	// Listen and Server in
	_ = router.Run(":8081")
}
