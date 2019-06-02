package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"webchat/app/format"
	"webchat/app/model"
	"webchat/app/service"
	"webchat/util"
)

/**
注册控制器
 */
func UserRegister(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	postUser := &model.User{}
	err := util.Bind(r, postUser)
	if err != nil {
		format.Fail(w,err.Error())
	}

	password := postUser.Password
	nickName := postUser.NickName
	sex := postUser.Sex
	mobile := postUser.Mobile
	avatar := ""
	userService :=service.UserService{}
	user,err := userService.Register(mobile,password,nickName,avatar,sex)
	if err != nil {
		format.Fail(w,err.Error())
	} else {
		format.Success(w,user,"")
	}
}

/**
登录控制器
 */
func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	postUser := &model.User{}
	err := util.Bind(r,postUser)
	if err != nil {
		format.Fail(w,err.Error())
	}

	password := postUser.Password
	mobile := postUser.Mobile
	userService :=service.UserService{}
	user,err :=userService.Login(mobile,password)
	if err != nil {
		format.Fail(w,err.Error())
	} else {
		format.Success(w,user,"")
	}
}


