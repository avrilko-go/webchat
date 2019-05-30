package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"webchat/app/format"
	"webchat/app/service"
)

/**
注册控制器
 */
func UserRegister(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	r.ParseForm()
	password := r.Form.Get("password")
	nickName := r.Form.Get("nick_name")
	sex := r.Form.Get("sex")
	sexint, _ := strconv.Atoi(sex)
	mobile := r.Form.Get("mobile")
	avatar := ""
	userService :=service.UserService{}
	user,err := userService.Register(mobile,password,nickName,avatar,int8(sexint))
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
	r.ParseForm()
	password := r.Form.Get("password")
	mobile := r.Form.Get("mobile")
	userService :=service.UserService{}
	user,err :=userService.Login(mobile,password)
	if err != nil {
		format.Fail(w,err.Error())
	} else {
		format.Success(w,user,"")
	}
}


