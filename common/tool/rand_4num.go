package tool

import (
	"math/rand"
)

func RandCode(n int) string {
	bytes := make([]byte, n, n)
	for i := 0; i < n; i++ {
		bytes[i] = '0' + byte(rand.Intn(10))
	}

	return string(bytes)
}
