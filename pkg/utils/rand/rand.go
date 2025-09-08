package rand

import (
	"math/rand/v2"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	// 创建一个新的随机数生成器
	//r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano())+getGoroutineID(), 0))
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r.IntN(len(letters))]
	}
	return string(b)
}

func getGoroutineID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stack := string(buf[:n])
	idField := strings.Fields(strings.TrimPrefix(stack, "goroutine "))[0]
	id, _ := strconv.ParseUint(idField, 10, 64)
	return id
}
