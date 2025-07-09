package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", pong)
}
