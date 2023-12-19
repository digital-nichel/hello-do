package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/samber/do"
	"hello-do/service"
	"hello-do/store"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	injector := do.NewWithOpts(&do.InjectorOpts{
		HookAfterRegistration: func(injector *do.Injector, serviceName string) {
			log.Info().
				Str("Service", serviceName).
				Msg("Service registered")
		},

		HookAfterShutdown: func(injector *do.Injector, serviceName string) {
			log.Info().
				Str("Service", serviceName).
				Msg("Service shutdown")
		},

		Logf: func(format string, args ...any) {
			log.Debug().Msgf(format, args...)
		},
	})

	do.Provide(injector, store.NewStore)
	do.Provide(injector, service.NewService)

	s, err := do.Invoke[service.Service](injector)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create service")
	}

	if err := s.HealthCheck(); err != nil {
		log.Fatal().Err(err).Msg("Health check failed")
	}

	s.Handle()

	if err := injector.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("Shutdown failed")
	}
}
