package format

import (
	json2 "encoding/json"
	"log"
	"net/http"
)

type HttpResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func SendHttpResponse(w http.ResponseWriter,code int, data interface{}, msg string)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := HttpResponse{
		Code:code,
		Msg:msg,
		Data:data,
	}
	json, err := json2.Marshal(h)
	if err != nil {
		log.Panicln(err.Error())
	}
	w.Write(json)
}