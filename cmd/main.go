package main

import (
	"fmt"
	"net/http"

	"github.com/nadiannis/libry-api/internal/handler"
	"github.com/rs/zerolog/log"
)

type application struct {
	port     int
	handlers handler.Handlers
}

func main() {
	app := &application{
		port:     8080,
		handlers: handler.NewHandlers(),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.port),
		Handler: app.routes(),
	}

	log.Info().Msg("starting server on " + srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal().Msg(err.Error())
}
