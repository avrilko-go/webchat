package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"webchat/app/model"
	"webchat/util"
)

type UserService struct {

}

/**
用户注册逻辑
 */
func (s *UserService)Register(mobile, password, nickname, avatar string, sex int8) (model.User, error) {
	user := model.User{}
	_,err :=Db.Where("Mobile=?",mobile).Get(&user)
	if  err != nil {
		return user, err
	}
	if user.UserId > 0 {
		return user, errors.New("该用户已经存在")
	}
	user.Mobile = mobile
	user.NickName = nickname
	user.Avatar = avatar
	user.Sex = sex
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	user.Salt = fmt.Sprintf("%06d",rand.Int31n(10000))
	user.Password = util.MakePassword(password,user.Salt)
	user.Token = fmt.Sprintf("%08d",rand.Int31())

	_ ,err =Db.InsertOne(&user)
	if err != nil {
		return user, errors.New("添加用户失败")
	}

	return user,nil
}

/**
登录逻辑
 */
func (s *UserService)Login(mobile,password string)(model.User, error)  {
	user := model.User{}
	//查询数据
	_,err := Db.Where("Mobile=?",mobile).Get(&user)
	if err != nil {
		return user,err
	}
	if user.UserId < 1 {
		return user,errors.New("该用户不存在")
	}
	// 验证密码
	if !util.ValidatePassword(password,user.Salt,user.Password) {
		return user,errors.New("密码不正确")
	}
	// 刷新token
	token := util.MD5Encode(fmt.Sprintf("%d",time.Now().Unix()))
	user.Token = token
	Db.Id(user.UserId).Cols("Token").Update(&user)

	return user, nil
}
