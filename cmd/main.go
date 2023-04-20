package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {

	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	ctx := context.Background()

	rdb.HSet(ctx, "asd123-asd123", "tries", "1")

	rdb.Expire(ctx, "asd123-asd123", time.Second*30)

	fmt.Println("hell l0l")
}
