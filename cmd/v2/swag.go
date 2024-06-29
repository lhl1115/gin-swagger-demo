//go:build doc
// +build doc

package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	fmt.Println("init swag")
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v2"))
}
