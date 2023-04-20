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

	//rdb := redis.NewClient(
	//	&redis.Options{
	//		Addr:     "localhost:6379",
	//		Password: "",
	//		DB:       0,
	//	},
	//)
	//
	//ctx := context.Background()
	//
	//rdb.HSet(ctx, "asd123-asd123", "tries", "1")
	//
	//rdb.Expire(ctx, "asd123-asd123", time.Second*30)
	//
	//fmt.Println("hell l0l")
}
