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
		auth.POST("/sign-in", h.SignIn) // Обработчик для аутентификации
		auth.POST("/sign-up", h.SignUp) // Обработчик для регистрации

	}

	api := router.Group("/api", h.userIdentity)
	{
		book := api.Group("/book")
		{
			book.GET("/:id", h.GetBookByID)   // Получение книги по ID
			book.GET("/my", h.GetAllMyBooks)  // Получение всех книг пользователя
			book.POST("/", h.AddingUsersBook) // Добавление книг в список желаемых пользователя
			book.DELETE("/:id", h.DeleteBook) //
			//book.DELETE("/my/:id", h.DeleteOneOfTheUsersBook)
			admin := book.Group("/admin", h.userIdentity)
			{
				admin.POST("/", h.AddBook) // Добавление книг в библиотеку
				admin.PUT("/:id", h.AddBook)
			}

			all := book.Group("/all")
			{
				all.GET("/", h.GetAllBooks) // Получение всех книг
			}
		}
	}

	return router
}
