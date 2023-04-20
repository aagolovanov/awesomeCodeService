package main

import (
	"fmt"
	"github.com/aagolovanov/awesomeCodeService/api"
	"github.com/aagolovanov/awesomeCodeService/domain"
	"github.com/aagolovanov/awesomeCodeService/util"
	"log"
	"os"
)

func main() {

	// temporary config
	config := util.Config{
		Port:   8080,
		DBAddr: "localhost",
		DBPass: "6379",
		TTL:    30,
	}

	// TODO NewDomain
	dom := &domain.Domain{
		Storage: nil,
		Logg:    nil,
		Config:  nil,
	}

	server := api.NewServer(&config, dom, log.New(os.Stdout, "SERVER", log.LstdFlags))

	err := server.Start()
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
