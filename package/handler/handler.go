package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
