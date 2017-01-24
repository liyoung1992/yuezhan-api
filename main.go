package main

import (
	//	"fmt"
	_ "yuezhan-api/common"
	_ "yuezhan-api/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//	fmt.Println(beego.AppConfig.String("token"))
	//	fmt.Println(beego.AppConfig.String("key"))
	beego.Run()
}
