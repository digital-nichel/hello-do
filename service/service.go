package service

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/samber/do"
	"hello-do/store"
)

type Service interface {
	do.Healthcheckable
	do.Shutdownable

	Handle()
}

type service struct {
	store store.Store
}

func NewService(i *do.Injector) (Service, error) {
	s := do.MustInvoke[store.Store](i)

	return &service{
		store: s,
	}, nil
}

func (s *service) Handle() {
	items, err := s.store.GetItems()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get items")
	}

	log.Info().Strs("Items", items).Msg("Items")
}

func (s *service) Shutdown() error {
	return nil
}

func (s *service) HealthCheck() error {
	if s.store == nil {
		return fmt.Errorf("store is nil")
	}

	return s.store.HealthCheck()
}
