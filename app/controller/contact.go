package controller

import (
	"webchat/app/model"
	"webchat/util"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"webchat/app/format"
	"webchat/app/service"
	)

func AddFriend(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	contact := &model.Contact{}
	err := util.Bind(r, contact)
	if err != nil {
		format.Fail(w, err.Error())
		return
	}

	contactService := service.ContactService{}
	err = contactService.AddFriend(contact.UserId, contact.AddId)
	if err != nil {
		format.Fail(w, err.Error())
		return
	}
	format.Success(w,nil, "添加好友成功")
	return
}
