// Path: internal/pkg/utils/gm
// FileName: sm_test.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/26$ 20:14$

package gm

import (
	"crypto/hmac"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestHmac(t *testing.T) {
	key := []byte("1231")
	msg := []byte("231312")
	mac := hmac.New(New, key)
	mac.Write(msg)
	macData := mac.Sum(nil)
	p := base64.StdEncoding.EncodeToString(macData)
	fmt.Println("mac", string(p))
	fmt.Println("len", len(p))
	ok := hmac.Equal(macData, macData)
	fmt.Println("ok", ok)
}

func TestCheckPasswd(t *testing.T) {
	user := "admin"
	pass := "123456"
	en := EncryptPasswd(user, pass)
	b := CheckPasswd(user, pass, en)
	fmt.Println("b", b)
}
