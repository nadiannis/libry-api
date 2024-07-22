package main

import (
	"fmt"
	"net/http"

	"github.com/nadiannis/libry-api/internal/handler"
	"github.com/nadiannis/libry-api/internal/repository"
	"github.com/nadiannis/libry-api/internal/usecase"
	"github.com/rs/zerolog/log"
)

type application struct {
	port     int
	handlers handler.Handlers
}

func main() {
	repos := repository.NewRepositories()
	usecases := usecase.NewUsecases(repos)
	handlers := handler.NewHandlers(usecases)

	app := &application{
		port:     8080,
		handlers: handlers,
	}

	log.Info().Msg("add books")
	prepopulateBooks(usecases.Books)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.port),
		Handler: app.routes(),
	}

	log.Info().Msg("starting server on " + srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal().Msg(err.Error())
}
