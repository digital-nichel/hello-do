package store

import (
	"github.com/samber/do"
	"time"
)

type Store interface {
	do.Healthcheckable
	do.Shutdownable

	GetItems() ([]string, error)
}

type store struct {
	items []string
}

func NewStore(_ *do.Injector) (Store, error) {
	// Simulate some work
	time.Sleep(3 * time.Second)

	return &store{
		items: []string{"a", "b", "c"},
	}, nil
}

func (s *store) Shutdown() error {
	return nil
}

func (s *store) HealthCheck() error {
	return nil
}

func (s *store) GetItems() ([]string, error) {
	// Simulate some work
	time.Sleep(1 * time.Second)

	return s.items, nil
}
