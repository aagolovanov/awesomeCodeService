package main

import (
	"fmt"
	"github.com/aagolovanov/awesomeCodeService/api"
	"github.com/aagolovanov/awesomeCodeService/domain"
	"github.com/aagolovanov/awesomeCodeService/repository"
	"github.com/aagolovanov/awesomeCodeService/util"
	"log"
	"os"
)

func main() {

	config := util.LoadConfig()

	// loggers
	serverLogger := log.New(os.Stdout, "SERVER ", log.LstdFlags)
	domainLogger := log.New(os.Stdout, "DOMAIN ", log.LstdFlags)
	// todo repositoryLogger

	storage, err := repository.NewKeyDB(&config)
	if err != nil {
		log.Fatalf("Error while trying to connect to KeyDB: %v", err)
	}

	// TODO NewDomain
	dom := &domain.Domain{
		Storage: storage,
		Logg:    domainLogger,
		Config:  &config,
	}

	server := api.NewServer(&config, dom, serverLogger)

	err = server.Start()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
