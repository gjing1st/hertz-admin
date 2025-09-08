// rand_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/11$ 14:10$

package rand

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestStr(t *testing.T) {
	t1 := time.Now()
	for range 10 {
		//time.Sleep(time.Nanosecond)
		s := GoogleUUID32()
		fmt.Println(s)
	}
	var wg sync.WaitGroup
	for range 10 {

		wg.Go(func() {
			//time.Sleep(time.Nanosecond)
			s1 := generateRandomString(20)
			fmt.Println(s1)
		})

	}
	wg.Wait()
	//time.Sleep(time.Nanosecond)
	fmt.Println(time.Since(t1))
}

func BenchmarkGoogleUUID20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//S(20, false)
		generateRandomString(20)
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoogleUUID20()
	}
}
