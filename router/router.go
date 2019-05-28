package router

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/asset/*filepath",http.Dir("./asset"))
	registerView(router)
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

