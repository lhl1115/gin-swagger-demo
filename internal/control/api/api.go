package api

import (
	"gin-swagger-demo/pkg/serializer"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	v1 := router.Group("api/control")
	v1.GET("/devices", GetDevices)
}

type Device struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GetDevices
// @Summary Get devices
// @Tags Devices
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response{data=[]Device}
// @Router /api/control/devices [get]
func GetDevices(ctx *gin.Context) {
	serializer.SendJSON(ctx, 200, []Device{
		{ID: 1, Name: "AirSwitch", Address: "0x01"},
	}, "success")
}
