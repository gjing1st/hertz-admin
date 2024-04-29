// rand_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/11$ 14:10$

package rand

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	s := Str("1234567890", 5)
	fmt.Println(s)
	n := Intn(2)
	fmt.Println(n)
	aa := S(10, false)
	fmt.Println(aa)
}

func BenchmarkGoogleUUID20(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
func TestGenerateRuleName(t *testing.T) {
	fmt.Println(GenerateRuleName())
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRuleName()
	}
}
