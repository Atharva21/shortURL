package util

import (
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type configType struct {
	GoEnv        string
	RedisHost    string
	RedisPort    string
	EncodeLength int
	TTL          time.Duration
}

var Config *configType

func init() {
	Config = &configType{}
	Config.GoEnv = os.Getenv("GOENV")
	Config.RedisHost = os.Getenv("REDIS_HOST")
	Config.RedisPort = os.Getenv("REDIS_PORT")
	encodeLength, err := strconv.Atoi(os.Getenv("ENCODE_LENGTH"))
	if err != nil {
		encodeLength = 8
	}
	Config.EncodeLength = encodeLength
	ttl, err := strconv.Atoi(os.Getenv("TTL"))
	if err != nil {
		ttl = 3600
	}
	Config.TTL = time.Duration(ttl) * time.Second
}
