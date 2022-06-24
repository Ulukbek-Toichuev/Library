package handler

import (
	"github.com/Ulukbek-Toychuev/book_shop/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		book := api.Group("/book")
		{
			book.GET("/", h.GetAllBook)
			book.GET("/:id", h.GetBookByID)
			book.POST("/", h.AddBook)
			book.DELETE("/:id", h.DeleteBook)
			book.PUT("/:id", h.AddBook)
		}
	}

	return router
}
