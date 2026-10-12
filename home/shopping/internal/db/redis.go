package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx, Client = configureRedis()

func configureRedis() (context.Context, *redis.Client) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "admin",
	})
	return ctx, client
}

func GetNextIndex(table string) (int, error) {
	key := "index:" + table
	var nextIndex int

	err := Client.Watch(Ctx, func(tx *redis.Tx) error {
		nextIndex, err := tx.Get(Ctx, key).Int()
		if err != nil && err != redis.Nil {
			return GetError(key, err)
		}
		nextIndex++

		_, err = tx.TxPipelined(Ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(Ctx, key, nextIndex, 0)
			return nil
		})
		if err != nil {
			return PipeError(key, err)
		}
		return nil
	}, key)
	if err != nil {
		return 0, WatchError(key, err)
	}

	return nextIndex, nil
}
