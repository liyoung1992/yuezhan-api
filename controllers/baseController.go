package controllers

import (
	//	"yuezhan-api/common"
	//	"yuezhan-api/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//api请求前验证
func (b *BaseController) Prepare() {
	type Message struct {
		Token string `orm:"column(token)"`
		Key   string `orm:"column(key)"`
	}

}
