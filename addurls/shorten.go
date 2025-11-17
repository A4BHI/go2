package addurls

import (
	"context"
	"crypto/rand"
	"go2/db"
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
	code := ShortCode()

	db.Pool.Exec(context.Background(), "insert into links (username ,short_code,org_link,) values (?,?,?)", code)
}
