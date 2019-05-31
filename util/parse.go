package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func Bind(r *http.Request, args interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(contentType), "application/json") { //json请求 解析贼简单
		return BindJson(r, args)
	}

	if strings.Contains(strings.ToLower(contentType), "application/x-www-form-urlencoded") {
		return BindForm(r, args)
	}


	return errors.New("当前方法不支持")
}

/**
动态解析json
 */
func BindJson(r *http.Request, form interface{}) error {
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(s,form)
	if err != nil {
		return err
	}

	return nil
}

func BindForm(r *http.Request, form interface{}) error {


	return nil
}