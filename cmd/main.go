package main

import (
	"os"

	"github.com/aburtasov/todo-app"
	"github.com/aburtasov/todo-app/pkg/handler"
	"github.com/aburtasov/todo-app/pkg/repository"
	"github.com/aburtasov/todo-app/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	///////////////////////Init ConfigFile//////////////////////////////////////////////////////

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s ", err.Error())
	}

	/////////////////////////////////////////////////////////////////////////////////////

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	///////////////////////Init DB//////////////////////////////////////////////////////

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Failed to initialized db: %s", err.Error())
	}

	/////////////////////////////////////////////////////////////////////////////////////

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
