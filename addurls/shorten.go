package addurls

import (
	"crypto/rand"
	"math/big"
)

func ShortCode() string {

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		number, _ := rand.Int(rand.Reader, big.NewInt(int64(6)))
		code[i] = letters[number.Int64()]
	}

	return string(code)

}

func Url() {

}
