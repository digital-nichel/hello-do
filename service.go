package main

import (
	"fmt"
	"github.com/samber/do"
)

type Service interface {
	Start()
	Shutdown() error
	HealthCheck() error
}

type service struct {
	store Store
}

func NewService(i *do.Injector) (Service, error) {
	s := do.MustInvoke[Store](i)

	return &service{
		store: s,
	}, nil
}

func (s *service) Start() {
	println("service starting")

	s.store.Start()
}

func (s *service) Shutdown() error {
	println("service stopped")
	return nil
}

func (s *service) HealthCheck() error {
	return fmt.Errorf("service broken")
}
