package handler

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Atharva21/shortURL/util"
	"github.com/go-redis/redis/v8"
)

const (
	shortURLPrefix = "SHORTURL"
	longURLPrefix  = "LONGURL"
)

var redisClient *redis.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", util.Config.RedisHost, util.Config.RedisPort),
	})
	if redisClient == nil {
		log.Fatal("Error connecting to redis")
		panic(errors.New("Error connecting to redis server"))
	}
	log.Println("Connected to redis")
}

func CheckIfShortURLAbsent(shortURL string) (bool, error) {
	shortKey := fmt.Sprintf("%s:%s", shortURLPrefix, shortURL)
	_, err := redisClient.Get(ctx, shortKey).Result()
	if err != nil {
		return true, nil
	}
	return false, nil
}

func CheckIfLongURLPresent(longURL string) (bool, error) {
	longKey := fmt.Sprintf("%s:%s", longURLPrefix, longURL)
	_, err := redisClient.Get(ctx, longKey).Result()
	if err != nil {
		return false, nil
	}
	return true, nil
}

func Link(shortURL, longURL string) error {
	shortKey := fmt.Sprintf("%s:%s", shortURLPrefix, shortURL)
	log.Println("ttl is:", util.Config.TTL)
	_, err := redisClient.Set(ctx, shortKey, longURL, util.Config.TTL).Result()
	if err != nil {
		return err
	}
	longKey := fmt.Sprintf("%s:%s", longURLPrefix, longURL)
	_, err = redisClient.Set(ctx, longKey, shortURL, util.Config.TTL).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetShortURL(longURL string) (string, error) {
	longKey := fmt.Sprintf("%s:%s", longURLPrefix, longURL)
	shortURL, err := redisClient.Get(ctx, longKey).Result()
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func GetLongURL(shortURL string) (string, error) {
	shortKey := fmt.Sprintf("%s:%s", shortURLPrefix, shortURL)
	longURL, err := redisClient.Get(ctx, shortKey).Result()
	if err != nil {
		return "", err
	}
	return longURL, nil
}
