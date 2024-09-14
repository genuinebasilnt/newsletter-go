package main

import (
	"fmt"
	"genuinebasilnt/newsletter-go/internal/config"
	"genuinebasilnt/newsletter-go/internal/env"
	"genuinebasilnt/newsletter-go/internal/router"
	"log"
	"net/http"
)

func main() {
	env, err := env.SetupEnv()
	if err != nil {
		log.Fatalln("Cannnot initialize environment")
	}

	defer env.Pool.Close()

	config, err := config.GetConfiguration("config")
	if err != nil {
		env.Logger.Fatal().Msgf("Failed to get configuration: %s", err)
	}

	r := router.Router(env)

	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", config.ApplicationSettings.Port),
		Handler: r,
	}

	env.Logger.Info().Msgf("Started server on port: %d", config.ApplicationSettings.Port)
	if err := server.ListenAndServe(); err != nil {
		env.Logger.Fatal().Msgf("Cannot listen on port: %v: %v\n", config.ApplicationSettings.Port, err)
	}

}
