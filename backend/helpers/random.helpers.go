package helpers

import (
	"math/rand"
)

func RandStr(n uint) string {
	var str = "abcdefgh123456789"
	byt := make([]byte, n)

	for i := range byt {
		byt[i] = str[rand.Intn(int(len(str)))]
	}
	return string(byt)
}
