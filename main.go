package main

import (
	"github.com/samber/do"
	"log"
)

func main() {
	injector := do.New()

	do.Provide(injector, NewStore)
	do.Provide(injector, NewService)

	s := do.MustInvoke[Service](injector)
	s.Start()

	//if err := do.HealthCheck[Service](injector); err != nil {
	//	log.Fatal(err)
	//}

	if err := injector.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
