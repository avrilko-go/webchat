package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func UserLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	io.WriteString(w, "test")
}

