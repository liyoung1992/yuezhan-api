package common

import (
	"github.com/astaxie/beego/orm"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "gdzt:039.com@tcp(192.168.2.221:3306)/beego?charset=utf8&loc=Asia%2FShanghai")
}
