package main

import (
	"github.com/sinyavcev/todoGolang"
	"github.com/sinyavcev/todoGolang/handlers"
	"github.com/sinyavcev/todoGolang/repository"
	"github.com/sinyavcev/todoGolang/service"
	"log"
)

func main() {
	db, err := repository.Init()

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(todoGolang.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
