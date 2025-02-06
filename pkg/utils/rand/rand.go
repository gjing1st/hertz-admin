package rand

import (
	"math/rand/v2"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	// 创建一个新的随机数生成器
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r.IntN(len(letters))]
	}
	return string(b)
}
