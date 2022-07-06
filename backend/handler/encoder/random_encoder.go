package encoder

import (
	"math/rand"

	"github.com/Atharva21/shortURL/handler"
	"github.com/Atharva21/shortURL/util"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type randomEncoder struct {
	// encoder that randomly encodes the longURL, till its available.
}

func (encoder *randomEncoder) Encode(longURL string) (string, error) {
	var encodedURL []byte
	for {
		encodedURL = []byte{}
		for i := 0; i < util.Config.EncodeLength; i++ {
			encodedURL = append(encodedURL, characters[rand.Intn(len(characters))])
		}
		absent, err := handler.CheckIfShortURLAbsent(string(encodedURL)) // ◀ check if encoding already exists for current encodedURL
		if err != nil {
			return "", err
		}
		if absent {
			break
		}
	}
	return string(encodedURL), nil
}
