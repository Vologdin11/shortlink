package main

import (
	"log"
	"shortlink/internal/handler"
	"shortlink/internal/repository"
	"shortlink/internal/service"
	conf "shortlink/pkg/config"
	"shortlink/pkg/server"
)

func main() {
	//Переделать на CLI
	//Проверка использовать кэш или бд
	services := service.Service{}
	// if os.Args[1] == "--cache" {
	// 	log.Println("Use cache")
	// } else {
	//брать значение из переменных окружения
	config := conf.NewConfig("localhost", "5432", "postgres", "qweasd", "shortlink", 5, 4)
	postgres := repository.Postgres{}
	err := postgres.InitPostgres(config)
	if err != nil {
		log.Fatal(err)
		// }
	}
	services.DB = &postgres
	handlers := handler.NewHandler(&services)
	srv := server.Server{}
	log.Fatal(srv.Run("8000", handlers.InitRouter()))
}
