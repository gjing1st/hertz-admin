// Path: internal/apiserver/store
// FileName: crypto.go
// Created by bestTeam
// Author: GJing
// Date: 2023/10/24$ 14:47$

package store

import (
	"crypto/hmac"
	"encoding/base64"
	"fmt"
	"github.com/gjing1st/hertz-admin/pkg/utils/gm"
)

var dbkey = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
var dbiv = []byte{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func init() {
	_ = gm.SetIV(dbiv)
}
func EncryptString(plain string) (string, error) {
	plainBytes := []byte(plain)
	//cipherBytes, err := hard.DevCard.Encrypt(plainBytes, gmdevice.SGD_SM4_CBC)
	cipherBytes, err := gm.Sm4Cbc(dbkey, plainBytes, true)
	if err != nil {
		return "", fmt.Errorf("Sm4Cbc encrypt error: %s", err)
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

func DecryptString(cipher string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return "", fmt.Errorf("invalid base64 encoded cipher string: %s", err)
	}
	if len(cipherBytes) == 0 || len(cipherBytes)%gm.BlockSize != 0 {
		return "", fmt.Errorf("invalid cipher bytes length")
	}
	//plainBytes, err := hard.DevCard.Decrypt(cipherBytes, gmdevice.SGD_SM4_CBC)
	plainBytes, err := gm.Sm4Cbc(dbkey, cipherBytes, false)
	if err != nil {
		return "", fmt.Errorf("Sm4Cbc decrypt error: %s", err)
	}
	return string(plainBytes), nil
}

func ComputeCheckData(msg []byte) (value []byte) {
	mac := hmac.New(gm.New, dbkey)
	mac.Write(msg)
	return mac.Sum(nil)
}
