package router

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"strings"
	"webchat/app/controller"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/asset/*filepath",http.Dir("./asset"))
	registerView(router)
	registerLogic(router)
	return router
}

func registerView(router *httprouter.Router)  {
	tpl,err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _,v := range tpl.Templates() {
		tplName := v.Name()
		if strings.HasPrefix(tplName, "/") {
			router.GET(tplName, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
				err = v.ExecuteTemplate(writer,tplName,nil)
				if err != nil {
					log.Fatal(err.Error())
				}
			})
		}
	}
}

func registerLogic(router *httprouter.Router)  {
	router.POST("/user/register",controller.UserRegister)
	router.POST("/user/login",controller.UserLogin)
	router.POST("/contact/addFriend",controller.AddFriend)
	router.POST("/contact/loadFriend",controller.LoadFriend)
	router.POST("/contact/loadGroup",controller.LoadGroup)
	router.POST("/contact/createGroup",controller.CreateGroup)
	router.POST("/contact/joinGroup",controller.JoinGroup)
}

