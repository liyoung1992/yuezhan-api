package controllers

import (
	"encoding/json"
	"yuezhan-api/common"
	"yuezhan-api/models"

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
	var m Message
	json.Unmarshal(b.Ctx.Input.RequestBody, &m)
	if !common.Filter(m.Token, m.Key) {
		u.Data["json"] = &models.VerificateResult{Result: 1001, Message: "token或key错误!"}
	}
}
