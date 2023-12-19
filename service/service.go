package service

import (
	"fmt"
	"github.com/samber/do"
	"hello-do/store"
)

type Service interface {
	do.Healthcheckable
	do.Shutdownable

	GetItems() ([]string, error)
}

type service struct {
	store store.Store
}

func NewService(i *do.Injector) (Service, error) {
	s, err := do.Invoke[store.Store](i)
	if err != nil {
		return nil, err
	}

	return &service{
		store: s,
	}, nil
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

func (s *service) GetItems() ([]string, error) {
	return s.store.GetItems()
}
