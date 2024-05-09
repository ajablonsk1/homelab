package main

import "github.com/redis/go-redis/v9"

func main() {
	client := redis.NewClient(&redis.Options{})
}
