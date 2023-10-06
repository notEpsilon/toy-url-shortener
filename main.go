package main

import (
	"log"

	"github.com/notEpsilon/shorty/pkg/controller"
	"github.com/notEpsilon/shorty/pkg/repository"
	"github.com/notEpsilon/shorty/pkg/service"
)

func main() {
	cntrl := controller.NewController("127.0.0.1", 7000, service.NewServiceImpl(repository.NewInMemoryRepository()))
	log.Fatalln(cntrl.Start())
}
