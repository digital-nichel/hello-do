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
			log.Trace().Msgf(format, args...)
		},
	})

	do.Provide(injector, store.NewStore)
	do.ProvideNamed(injector, "Service-01", service.NewService)
	do.ProvideNamed(injector, "Service-02", service.NewService)

	s1, err := do.InvokeNamed[service.Service](injector, "Service-01")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create service")
	}

	if err := s1.HealthCheck(); err != nil {
		log.Fatal().Err(err).Msg("Health check failed")
	}

	items, err := s1.GetItems()
	if err != nil {
		log.Warn().Str("Service", "Service-01").Err(err).Msg("Get items failed")
	}

	log.Info().Str("Service", "Service-01").Strs("Items", items).Msg("Get items")

	_, err = s1.GetItems()
	if err != nil {
		log.Warn().Str("Service", "Service-01").Err(err).Msg("Get items failed")
	}

	s2 := do.MustInvokeNamed[service.Service](injector, "Service-02")
	_, err = s2.GetItems()
	if err != nil {
		log.Warn().Str("Service", "Service-02").Err(err).Msg("Get items failed")
	}

	log.Info().Str("Service", "Service-02").Strs("Items", items).Msg("Get items")

	if err := injector.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("Shutdown failed")
	}
}
