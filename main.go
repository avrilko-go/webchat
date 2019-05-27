package main

import (
	"net/http"
	router2 "webchat/router"
)

func main()  {
	router := router2.RegisterRouter()
	err := http.ListenAndServe(":8080", router)
	if  err != nil {
		panic("http服务启动失败")
	}
}

