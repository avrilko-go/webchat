package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
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
	err := r.ParseForm() // 必须先运行解析函数才能得到表单
	if err != nil {
		return err
	}
	err = mapForm(args,r.Form) // r.Form 将内容解析到一个map[string][]string中 ，如果为单元素 ，则取值为form["test"][0], 取第0个下标，此处坑
	if err != nil {
		return err
	}
	return nil
}

func mapForm(ptr interface{}, form map[string][]string) error  {
	val := reflect.ValueOf(ptr).Elem()
	typ := reflect.TypeOf(ptr).Elem()
	// 我们默认传进来的是一个struct指针 可以反射出结构体内的元素数量
	num := typ.NumField()
	if num <= 0 { // 空的结构体不解析
		return errors.New("该结构体为空")
	}

	for i:=0;i<typ.NumField();i++ { // 循环结构体数据
		typeField := typ.Field(i)
		valeField := val.Field(i)
		if !valeField.CanSet() { // 如果不是一个可以赋值的指针则跳出循环
			continue
		}

		// 获取单个元素的类型
		typeOne := valeField.Kind() // 类型
		formName := typeField.Tag.Get("form") // 获取tag值

		if formName == "" { // 没有打表单的名称
			// 获取默认的字段名称
			formName = typeField.Name
			if typeOne == reflect.Struct { // 如果单个字段还是struct结构体则做一个递归转换
				err := mapForm(valeField.Addr().Interface(),form)
				if err != nil {
					return nil
				}
				continue // 跳出循环
			}
		}

		inputValue,exist := form[formName]
		if !exist { // 表单输入中没有此字段则不解析
			continue
		}

		numInputValue := len(inputValue)

		if typeOne == reflect.Slice && numInputValue > 0 { // 是一个切片类型
			// 获取切片的数据类型
			typeSlice := valeField.Type().Elem().Kind()
			slices := reflect.MakeSlice(valeField.Type(),numInputValue,numInputValue) //创建切片
			for j:= 0; j< numInputValue;j++ {
				err := setWithProperType(typeSlice, inputValue[j], slices.Index(j))
				if err != nil {
					return err
				}
			}
			val.Field(i).Set(slices)
		} else {
			if _,isTime := valeField.Interface().(time.Time); isTime { // 是时间类型
				if err := setTimeField(inputValue[0],typeField,valeField) ; err != nil {
					return err
				}
				continue
			} else {
				err := setWithProperType(typeOne,inputValue[0],valeField)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	switch valueKind {
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32, structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	default:
		return errors.New("Unknown type")
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return nil
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}


func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	//2018-01-02 01:02:03

	if timeFormat == "" {
		timeFormat = "2006-01-02 15:04:05"
		val = strings.Replace(val,"/","-",0)
		num := len(strings.Split(val," "))
		if num==1{
			val = val +" 00:00:00"
		}else{
			//2018-01-02 00
			num =len(strings.Split(val,":"))

			if num==1{
				val = val +":00:00"
			}else if num==2{
				val = val +":00"
			}
		}

	}

	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return err
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return nil
}





















