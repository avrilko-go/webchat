package model

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *xorm.Engine

func init() {
	var err error
	Db, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/hb?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	Db.SetMapper(core.SameMapper{})
	Db.ShowSQL(true)
	Db.SetMaxOpenConns(2)
	Db.Sync2(new(User))
}
