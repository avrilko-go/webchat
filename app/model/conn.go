package model

import "github.com/go-xorm/xorm"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *xorm.Engine

func init()  {
	db,err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3304)/quxiaotao?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	Db = db
}

