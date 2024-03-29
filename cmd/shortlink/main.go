package main

import (
	"log"
	"math/rand"
	"os"
	"shortlink/internal/handler"
	"shortlink/internal/repository/inmemory"
	"shortlink/internal/repository/postgres"
	"shortlink/internal/service"
	conf "shortlink/pkg/config"
	"shortlink/pkg/server"
	"time"
)

func main() {
	services := service.Service{}
	handlers := handler.NewHandler(&services)
	//для рандомных чисел
	rand.Seed(time.Now().Unix())

	//Использовать кэш или бд
	if os.Getenv("DB") == "postgres" {
		config, err := conf.NewConfig()
		if err != nil {
			log.Fatal(err)
		}
		postgre := postgres.Postgres{}
		err = postgre.InitPostgres(config)
		if err != nil {
			log.Fatal(err)
		}
		services.DB = &postgre
	} else {
		services.DB = inmemory.NewInmemory()
	}

	srv := server.Server{}
	//брать порт из переменных
	log.Fatal(srv.Run("8000", handlers))
}
