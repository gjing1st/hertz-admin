package utils

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var ifc interface{}

func BenchmarkInt(b *testing.B) {
	ifc = "123577"
	for i := 0; i < b.N; i++ {
		Int(ifc)
	}
}

func BenchmarkString(b *testing.B) {
	ifc = "sdsd"
	for i := 0; i < b.N; i++ {
		String(ifc)
	}
}

func TestCode(t *testing.T) {
	// 假设接收到的字节切片是这样的
	temp := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38} // 代表 "12345678" 的字节切片

	// 将字节切片转换为十六进制字符串
	hexStr := hex.EncodeToString(temp)
	fmt.Println("十六进制字符串:", hexStr) // 输出: 3132333435363738

	// 现在我们要将这个十六进制字符串转换回原始字符串
	decodedBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("解码错误:", err)
		return
	}

	// 将解码后的字节转换为字符串
	originalStr := string(decodedBytes)
	fmt.Println("原始字符串:", originalStr) // 输出: 12345678
}
