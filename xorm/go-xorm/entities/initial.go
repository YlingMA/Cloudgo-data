package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	// Init xorm engine
	Engine, err := xorm.NewEngine("mysql", "root:myl.rl.0676@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	err = Engine.Sync2(new(UserInfo))
	checkErr(err)
	engine = Engine
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
