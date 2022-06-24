package main

import (
	"log"

	"github.com/Ulukbek-Toychuev/book_shop/internal/server"
	"github.com/Ulukbek-Toychuev/book_shop/package/handler"
)

func main() {
	srv := new(server.Server)

	handlers := new(handler.Handler)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
