package main

import (
	_ "gin-swagger-demo/docs/swag"
	"github.com/gin-gonic/gin"
	v2 "github.com/swaggo/gin-swagger/example/multiple/api/v2"
)

var swagHandler gin.HandlerFunc

// @title Test API
// @version 2.0
// @description Test API v2.0
// @contact.name -V1 API
// @contact.url http://localhost:8081/swagger/v1/index.html
func main() {
	// New gin router
	router := gin.New()

	// Register v2 endpoints
	v2.Register(router)
	if swagHandler != nil {
		router.GET("/swagger/v2/*any", swagHandler)
	}

	// Listen and Server in
	_ = router.Run(":8081")
}
