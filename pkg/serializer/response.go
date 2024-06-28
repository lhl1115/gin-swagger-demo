package serializer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

func SendJSON(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
