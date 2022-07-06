package encoder

import (
	"math/rand"
	"time"
)

var URLEncoder encoder

type encoder interface {
	Encode(string) (string, error)
}

func init() {
	rand.Seed(time.Now().UnixNano())
	URLEncoder = &randomEncoder{} // ◀ can attach any implementation here
}
