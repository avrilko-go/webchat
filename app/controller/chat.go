package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"webchat/app/service"
	"github.com/gorilla/websocket"
)

func Chat(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	// 先鉴权
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId,_ := strconv.Atoi(id)
	checkPass:= service.UserService{}.CheckToken(userId, token)
	// 然后将http请求升级为WebSocket请求
	conn, err :=websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return checkPass
		},
	}.Upgrade(w,r,nil)


}
