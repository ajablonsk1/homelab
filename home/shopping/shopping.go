package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx, client = configureRedis()

func configureRedis() (context.Context, *redis.Client) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "admin",
	})
	return ctx, client
}

func main() {
	args := os.Args
	if len(args) > 1 {
		parseArguments(args)
	} else {
		getShoppingList()
	}
}

func getShoppingList() {
	iter := client.Scan(ctx, 0, "shopping:*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		val, err := client.HGet(ctx, key, "name").Result()
		if err != nil {
			log.Fatalf("Encountered error during retrieving shoppig list key %s: %s", key, err)
		}
		fmt.Println(val)
	}
}

func parseArguments(args []string) {
	cmd := args[1]
	switch cmd {
	case "add":
		addProductToShoppingList(args[2:])
	default:
		log.Fatal("Unrecognized command! Usage shopping (add) <product>")
	}
}

func addProductToShoppingList(product []string) {
	index := getNextIndex("shopping")
	key := "shopping:" + strconv.Itoa(index)
	_, err := client.HSet(ctx, key, product).Result()
	if err != nil {
		log.Fatalf("Encountered error during setting hkey %s: %s", key, err)
	}
}

func getNextIndex(table string) int {
	key := "index:" + table
	var indexInt int

	index, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		indexInt = 0
	} else if err != nil {
		log.Fatalf("Encountered error during retrieving key %s: %s", key, err)
	} else {
		indexInt, _ = strconv.Atoi(index)
	}

	nextIndex := indexInt + 1
	_, err = client.Set(ctx, key, nextIndex, 0).Result()
	if err != nil {
		log.Fatalf("Encountered error during seting key %s with value %d: %s", key, nextIndex, err)
	}

	return nextIndex
}
