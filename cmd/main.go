package main

import (
	"github.com/joho/godotenv"
	server "github.com/mazhaboy/rest-api-auth"
	"github.com/mazhaboy/rest-api-auth/configs"
	"github.com/mazhaboy/rest-api-auth/pkg/handler"
	"github.com/mazhaboy/rest-api-auth/pkg/repository"
	"github.com/mazhaboy/rest-api-auth/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	path  = "schema"
	dbURI = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfigs(); err != nil {
		logrus.Fatalf("error initializnig configs: %s", err.Error())
	}

	if err := configs.Migrate(path, dbURI); err != nil {
		logrus.Fatalf("unable to migrate: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("unable to load env : %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(server.Server)

	err = srv.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
