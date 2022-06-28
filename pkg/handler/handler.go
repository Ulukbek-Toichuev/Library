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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)

	}

	api := router.Group("/api", h.userIdentity)
	{
		book := api.Group("/book")
		{
			book.GET("/", h.GetAllBook)
			book.GET("/:id", h.GetBookByID)
			book.DELETE("/:id", h.DeleteBook)
			book.PUT("/:id", h.AddBook)
			admin := book.Group("/admin", h.userIdentity)
			{
				admin.POST("/", h.AddBook)
			}
		}
	}

	return router
}
