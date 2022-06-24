package main

import (
	"log"

	"github.com/Ulukbek-Toychuev/book_shop/internal/server"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/handler"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/repository"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error occured read config file: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
