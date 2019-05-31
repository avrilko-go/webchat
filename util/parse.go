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

	if strings.Contains(strings.ToLower(contentType), "application/x-www-form-urlencoded") { // 处理表单请求
		return BindForm(r, args)
	}


	return errors.New("当前方法不支持")
}

/**
动态解析json
 */
func BindJson(r *http.Request, args interface{}) error {
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(s,args)
	if err != nil {
		return err
	}

	return nil
}

/**
处理表单请求(使用反射库动态解析)
 */
func BindForm(r *http.Request, args interface{}) error {

	return nil
}