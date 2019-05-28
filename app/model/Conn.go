package model

import "github.com/go-xorm/xorm"
import  _ "github.com/go-sql-driver/mysql"

var Db *xorm.Engine

func init()  {
	Db,_ = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3304)/quxiaotao?charset=utf8")
}

