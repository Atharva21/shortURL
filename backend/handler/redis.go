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
	ShortURL = "SHORTURL"
	LongURL  = "LONGURL"
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

func CheckIfShortURLAbsent(url string) (bool, error) {
	ok, err := redisClient.HExists(ctx, ShortURL, url).Result()
	if err != nil || ok {
		return false, fmt.Errorf("error while querying redis")
	}
	return true, nil
}

func CheckIfLongURLPresent(url string) (bool, error) {
	ok, err := redisClient.HExists(ctx, LongURL, url).Result()
	if err != nil {
		return false, fmt.Errorf("error while querying redis")
	}
	return ok, nil
}

func Link(shortURL, longURL string) error {
	_, err := redisClient.HSet(ctx, ShortURL, shortURL, longURL).Result()
	if err != nil {
		return err
	}
	_, err = redisClient.HSet(ctx, LongURL, longURL, shortURL).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetShortURL(longURL string) (string, error) {
	val, err := redisClient.HGet(ctx, LongURL, longURL).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func GetLongURL(shortURL string) (string, error) {
	val, err := redisClient.HGet(ctx, ShortURL, shortURL).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
