package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
        "time"

	"github.com/Ulukbek-Toychuev/book_shop/internal/server"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/handler"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/repository"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/service"
	
        "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Library API
// @version 1.0
// @description REST API server for Library

// @BasePath /

// @license.name  MIT License
// @license.url   https://opensource.org/licenses/MIT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
        time.Sleep(5 * time.Second)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occured read config file %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		Password: viper.GetString("db.password"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server %s", err.Error())
		}
	}()

	logrus.Print("Library app started!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Library app shutdown!")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Error("Error occured on server shutting down %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("Error occured on db connection closed %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
