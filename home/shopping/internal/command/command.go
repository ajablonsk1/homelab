package command

import (
	"fmt"
	"strconv"

	"github.com/ajablonsk1/homelab/home/shopping/internal/db"
	"github.com/ajablonsk1/homelab/home/shopping/internal/parse"
	"github.com/redis/go-redis/v9"
)

func Ls() error {
	iter := db.Client.Scan(db.Ctx, 0, "shopping:*", 0).Iterator()
	for iter.Next(db.Ctx) {
		key := iter.Val()
		val, err := db.Client.HGet(db.Ctx, key, "name").Result()
		if err != nil {
			return db.HGetError(key, err)
		}
		fmt.Println(val)
	}
	return nil
}

func Add(product *parse.Product) error {
	_, err := db.Client.Get(db.Ctx, "shopping_name:"+product.Name).Result()
	if err != redis.Nil {
		return fmt.Errorf("shopping item with provided name already exists: %s", product.Name)
	}

	index, err := db.GetNextIndex("shopping")
	if err != nil {
		return err
	}
	key := "product:" + strconv.Itoa(index)

	pipe := db.Client.Pipeline()
	pipe.HSet(db.Ctx, key, product)
	pipe.Set(db.Ctx, "shopping_name:"+product.Name, index, 0)
	_, err = pipe.Exec(db.Ctx)
	if err != nil {
		return db.PipeError(key, err)
	}

	return nil
}

func Update(name string, product *parse.Product) error {
	key := "shopping_name:" + product.Name
	idx, err := db.Client.Get(db.Ctx, key).Result()
	if err != nil {
		return db.GetError(key, err)
	}
	key = "shopping:" + idx

	err = db.Client.Watch(db.Ctx, func(tx *redis.Tx) error {
		dbProductMap, err := db.Client.HGetAll(db.Ctx, key).Result()
		if err != nil {
			return db.HGetError(key, err)
		}

		dbProduct, err := parse.ProductFromMap(dbProductMap)
		if err != nil {
			return err
		}

		indexKey := "shopping_name:" + dbProduct.Name
		if product.Name != "" && dbProduct.Name != product.Name {
			tx.Del(db.Ctx, indexKey)
			indexKey = "shopping_name:" + product.Name
		}

		dbProduct.Update(product)

		_, err = tx.TxPipelined(db.Ctx, func(pipe redis.Pipeliner) error {
			pipe.HSet(db.Ctx, key, dbProduct)
			pipe.Set(db.Ctx, indexKey, idx, 0)
			return nil
		})
		if err != nil {
			return db.PipeError(key, err)
		}
		return nil
	}, key)

	return err
}

func Rm(name string) error {
	key := "shopping_name:" + name
	idx, err := db.Client.Get(db.Ctx, key).Result()
	if err != nil {
		return db.GetError(key, err)
	}

	key = "shopping:" + idx
	err = db.Client.Del(db.Ctx, key).Err()
	return db.DelError(key, err)
}
