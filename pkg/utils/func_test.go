// Path: pkg/utils
// FileName: func_test.go
// Created by bestTeam
// Author: GJing
// Date: 2022/11/16$ 11:20$

package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestDiffNatureDays(t *testing.T) {
	t1 := time.Now().Unix()
	t2 := time.Now().Unix() + 60*60*13
	d := DiffNatureDays(t1, t2)
	fmt.Println(d)
}

func TestName(t *testing.T) {
	for i := 1; i < 100; i++ {
		is := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				is = true
				break
			}
		}
		if !is {
			fmt.Println(i)
		}
	}
}
