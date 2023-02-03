package ugen

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

var (
	digit = []byte("0123456789")
	lower = []byte("abcdefghijklmnopqrstuvwxyz")
	upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	digitLower = append(digit, lower...)
	digitUpper = append(digit, upper...)
	lowerUpper = append(lower, upper...)

	digitLowerUpper = append(digit, lowerUpper...)
)

func gen(l int, b, ch []byte) {
	if len(ch) == 0 {
		if _, err := io.ReadAtLeast(rand.Reader, b, l); err != nil {
			panic(fmt.Sprintf("Error reading random bytes: %v\n", err))
		}

		return
	}

	max := big.NewInt(int64(len(ch)))

	for i := range b {
		ri, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(fmt.Sprintf("Error producing random integer: %v\n", err))
		}
		b[i] = ch[int(ri.Int64())]
	}
}
