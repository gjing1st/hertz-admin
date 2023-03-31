// Path: pkg/utils
// FileName: func_test.go
// Created by dkedTeam
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
