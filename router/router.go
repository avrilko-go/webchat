package router

import (
	"github.com/julienschmidt/httprouter"
	"webchat/app/controller"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/login", controller.UserLogin)
	return router
}

