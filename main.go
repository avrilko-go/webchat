package main

import (
	"net/http"
	"webchat/app/model"
	router2 "webchat/router"
)

func main() {
	model.Db.Query("select * from User")
	router := router2.RegisterRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic("http服务启动失败")
	}
}
