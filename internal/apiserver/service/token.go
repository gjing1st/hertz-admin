// Path: internal/apiserver/service
// FileName: token.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 17:49$

package service

import (
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	"github.com/gjing1st/hertz-admin/pkg/utils/rand"
	"strconv"
)

type TokenService struct {
}

const tokenPrefix = "admin:"

// GenerateToken
// @description: 生成token
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/4/15 14:37
// @success:
func (td TokenService) GenerateToken(u *entity.UserTokenInfo) (token string, errCode error) {
	data := map[string]interface{}{
		"id":      u.Id,
		"name":    u.Name,
		"role_id": u.RoleId,
		//"login_type":u.LoginType,
	}
	uuid := rand.GoogleUUID32()
	token = "L_admin_" + strconv.Itoa(int(u.Id)) + "_" + uuid
	//放入缓存
	hKey := tokenPrefix + token
	errCode = gCache.RemoveSet(hKey, data)

	return
}

// GetInfo
// @description: 获取token对应的信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 18:22
// @success:
func (td TokenService) GetInfo(token string) (u entity.UserTokenInfo, errCode error) {
	hKey := tokenPrefix + token
	v, errCode1 := gCache.Get(hKey)
	if errCode != nil || v == nil {
		errCode = errCode1
		return
	}
	m := v.(map[string]interface{})
	u.Id = uint(utils.Int(m["id"]))
	u.Name = utils.String(m["name"])
	u.RoleId = utils.Int(m["role_id"])
	return
}

func (td TokenService) RemoveToken(token string) bool {
	hKey := tokenPrefix + token
	return gCache.Remove(hKey)
}
