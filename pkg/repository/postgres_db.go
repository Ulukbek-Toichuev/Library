package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	BookTable      = "book"
	UsersTable     = "users"
	UsersBookTable = "users_book"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBname   string
	Password string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		logrus.Fatalf("Error, cant connect to database: %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	} else {
		logrus.Println("We succesfully connected to db!")
	}

	return db, nil
}
