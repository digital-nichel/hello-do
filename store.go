package main

import "github.com/samber/do"

type Store interface {
	Start()
	Shutdown() error
}

type store struct {
}

func NewStore(i *do.Injector) (Store, error) {
	return &store{}, nil
}

func (s *store) Start() {
	println("store starting")
}

func (s *store) Shutdown() error {
	println("store stopped")
	return nil
}
