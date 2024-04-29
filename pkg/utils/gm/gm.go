// Path: internal/pkg/utils/gm
// FileName: gm.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/26$ 20:22$

package gm

import (
	"crypto/hmac"
	"encoding/base64"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	log "github.com/sirupsen/logrus"
)

// EncryptPasswd
// @description: 加密密码。使用用户名作为key进行加密
// @param: username string 用户名。这里以用户名作为key
// @param: password string 原始明文密码
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/26 20:24
// @success:
func EncryptPasswd(username, password string) (data string) {
	mac := hmac.New(New, []byte(username))
	mac.Write([]byte(password))
	macData := mac.Sum(nil)
	data = base64.StdEncoding.EncodeToString(macData)
	return data
}

// CheckPasswd
// @description: 验证密码
// @param: user string 用户名。这里以用户名作为key
// @param: password string 原始明文密码
// @param: encryptPasswd string 加密后的密码(数据库中的密码)
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 8:58
// @success:
func CheckPasswd(user, password, encryptPasswd string) (res bool) {
	mac := hmac.New(New, []byte(user))
	mac.Write([]byte(password))
	macData := mac.Sum(nil)
	pb, err := base64.StdEncoding.DecodeString(encryptPasswd)
	if err != nil {
		functions.AddErrLog(log.Fields{"msg": "密码校验，数据库密码转byte失败", "err": err})
	}
	res = hmac.Equal(macData, pb)
	return
}
