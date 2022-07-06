package util

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type configType struct {
	GoEnv        string
	RedisHost    string
	RedisPort    string
	EncodeLength int
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
}
