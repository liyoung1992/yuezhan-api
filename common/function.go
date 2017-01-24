package common

import (
	"fmt"
	"strings"

	// MySQL driver
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func Filter(token string, key string) bool {

	fmt.Println(beego.AppConfig.String("token"))
	fmt.Println(beego.AppConfig.String("key"))
	fmt.Println(strings.EqualFold(token, beego.AppConfig.String("token")))
	if strings.EqualFold(token, beego.AppConfig.String("token")) && strings.EqualFold(key, beego.AppConfig.String("key")) {
		fmt.Println("true")
		return true
	} else {
		fmt.Println("false")
		return false
	}
}
