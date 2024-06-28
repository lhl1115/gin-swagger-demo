package api

import (
	"gin-swagger-demo/pkg/serializer"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	v1 := router.Group("api/v1")
	v1.GET("/books", GetBooks)
}

type Book struct {
	ID     int     `json:"id,omitempty"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   *uint16 `json:"year"`
}

// GetBooks
// @Summary Get a list of books in the the store
// @Tags Books
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.Response{data=[]Book}
// @Router /api/v1/books [get]
func GetBooks(ctx *gin.Context) {
	serializer.SendJSON(ctx, 200, []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1", Year: nil},
		{ID: 2, Title: "Book 2", Author: "Author 2", Year: nil},
	}, "success")
}
