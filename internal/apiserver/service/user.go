// Path: internal/apiserver/service
// FileName: user.go
// Created by bestTeam
// Author: GJing
// Date: 2022/12/27$ 15:54$

package service

import (
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/entity"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/request"
	"github.com/gjing1st/hertz-admin/internal/apiserver/model/response"
	"github.com/gjing1st/hertz-admin/internal/apiserver/store/database"
	"github.com/gjing1st/hertz-admin/internal/pkg/config"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	"github.com/gjing1st/hertz-admin/pkg/errcode"
	"github.com/gjing1st/hertz-admin/pkg/utils/gm"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
}

var userDB = database.UserDB{}

// Create
// @description: 创建管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 15:58
// @success:
func (us *UserService) Create(req *request.UserCreate) (err error) {
	err = us.CreateUser(req)
	return
}

// CreateUser
// @description: 创建用户
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/17 14:26
// @success:
func (us *UserService) CreateUser(req *request.UserCreate) (errCode error) {
	user, errCode3 := userDB.GetByNameLoginType(req.Name, req.LoginType)
	if errCode3 != nil && errCode3 != errcode.UserNotFound {
		errCode = errCode3
		return
	}
	if user.ID != 0 {
		functions.AddErrLog(log.Fields{"msg": "创建管理员，该用户已存在", "userName": req.Name})
		errCode = errcode.UserNameExist
		return
	}
	var data entity.User
	data.Name = req.Name
	//data.NickName = req.Name
	//data.UserSerialNum = req.Serial
	data.Password = gm.EncryptPasswd(data.Name, req.Pin)
	data.LoginType = req.LoginType
	data.RoleId = req.RoleId
	_, errCode = userDB.Create(nil, &data)
	return
}

// Login
// @description: 管理员登录
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/27 17:05
// @success:
func (us *UserService) Login(req *request.UserLogin) (res response.UserLogin, err error) {
	user, err := userDB.GetByNameLoginType(req.Name, req.LoginType)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errcode.UserNotFound
		return
	}
	if len(req.Password) == 0 {
		req.Password = req.Pin
	}
	//验证密码
	ok := gm.CheckPasswd(req.Name, req.Password, user.Password)
	if !ok {
		if user.ErrNum >= config.Config.Base.PwdMaxErrNum {
			err = errcode.PwdMaxErr
			return
		}
		//记录密码错误次数
		retryCount := config.Config.Base.PwdMaxErrNum - user.ErrNum - 1
		if retryCount < 0 {
			res.RetryCount = 0
		} else {
			res.RetryCount = retryCount
			err = userDB.AddErrNum(int(user.ID))
		}
		err = errcode.PwdErr
		return
	} else {
		//密码错误次数归0
		_ = userDB.ClearErrNum(int(user.ID))
	}
	//删除之前的toekn，断点登录
	if user.Token != "" {
		TokenService{}.RemoveToken(user.Token)

	}

	var info entity.UserTokenInfo
	info.Id = user.ID
	info.Name = user.Name
	info.RoleId = user.RoleId

	token, err1 := TokenService{}.GenerateToken(&info)
	if err1 != nil {
		err = err1
		return
	}
	user.Token = token
	//更新数据库token

	err = userDB.UpdateToken(user.ID, user.Token)
	res.User = user
	//u, _ := TokenService{}.GetInfo(user.Token)
	//fmt.Println("u===", u)
	return
}

// List
// @description: 用户列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/28 17:25
// @success:
func (us *UserService) List(req *request.UserList) (list interface{}, total int64, err error) {
	list, total, err = userDB.List(req)
	return
}

// InfoByName
// @description: 操作员查询管理员列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/29 20:19
// @success:
func (us *UserService) InfoByName(name string) (list interface{}, total int64, err error) {
	user, err1 := userDB.GetByName(name)
	if err1 != nil {
		err = err1
		return
	}
	total = 1
	var list1 []*entity.User
	list = append(list1, user)
	return
}

// DeleteById
// @description: 删除指定id
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/12/30 11:26
// @success:
func (us *UserService) DeleteById(userid int) (err error) {
	err = userDB.DeleteById(nil, userid)
	return
}

// DeleteUser
// @description: 删除管理员
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/3/17 16:18
// @success:
func (us *UserService) DeleteUser(req *request.UserDelete) (err error) {
	err = userDB.DeleteById(nil, req.ID)
	return
}

// ChangePasswd
// @Description 修改用户名口令密码
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/6/8 9:39
func (us *UserService) ChangePasswd(req *request.ChangePasswd, username string, loginType int) error {
	user, err := userDB.GetByNameLoginType(username, loginType)
	if err != nil {
		return err
	}
	//验证密码是否正确
	ok := gm.CheckPasswd(username, req.OldPassword, user.Password)
	if !ok {
		err = errcode.PasswordErr
		return err
	}
	user.Password = gm.EncryptPasswd(username, req.NewPassword)
	err = userDB.UpdatePwd(user)
	//密码错误次数归0
	_ = userDB.ClearErrNum(int(user.ID))
	return err
}

func (us *UserService) Register(req *request.UserRegister) (err error) {
	user, err1 := userDB.GetByPhone(req.Phone)
	if err1 == nil && user.ID != 0 {
		err = errcode.UserNameExist
		return
	}
	user.Name = req.Username
	user.Phone = req.Phone
	user.NickName = req.Username
	user.Password = gm.EncryptPasswd(user.Name, req.Password)
	return userDB.Save(user)
}
