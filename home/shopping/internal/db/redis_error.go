package db

import "fmt"

func GetError(key string, err error) error {
	return fmt.Errorf("encountered an error with retrieving key { %s }: %w", key, err)
}

func HGetError(key string, err error) error {
	return fmt.Errorf("encountered an error with retrieving hkey { %s }: %w", key, err)
}

func SetError(key string, err error) error {
	return fmt.Errorf("encountered an error with setting key { %s }: %w", key, err)
}

func HSetError(key string, err error) error {
	return fmt.Errorf("encountered an error with setting hkey { %s }: %w", key, err)
}

func PipeError(key string, err error) error {
	return fmt.Errorf("encountered an error when running pipe key { %s }: %w", key, err)
}

func WatchError(key string, err error) error {
	return fmt.Errorf("encountered an error with watch key { %s }: %w", key, err)
}

func DelError(key string, err error) error {
	return fmt.Errorf("encountered an error with deleting key { %s }: %w", key, err)
}
