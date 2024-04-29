// Path: pkg/errcode
// FileName: code_test.go
// Created by bestTeam
// Author: GJing
// Date: 2023/3/28$ 21:29$

package errcode

import (
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

func TestCodeString(t *testing.T) {
	e := JsonMarshal()
	fmt.Println(e.Error())
}

func JsonMarshal() (e Err) {
	var s = `{name:"11",age:`
	b, err := sonic.Marshal(s)
	if err == nil {
		//操作成功
		//return
	}
	var un map[string]interface{}
	err = sonic.Unmarshal(b, &un)
	if err != nil {
		return HaSysJsonUnMarshalErr
	}
	return
}
