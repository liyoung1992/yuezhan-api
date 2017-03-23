package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"yuezhan-api/common"
	"yuezhan-api/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

//用户注册
func (u *UserController) Register() {
	type Message struct {
		UserName   string `orm:"column(userName)"`
		Password   string `orm:"column(password)"`
		Salt       string `orm:"column(salt)"`
		Token      string `orm:"column(token)"`
		Key        string `orm:"column(key)"`
		CreateTime int    `orm:"column(createTime)"`
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)
	if common.Filter(m.Token, m.Key) {
		e := models.AddUser(models.User{UserName: m.UserName, Password: m.Password, Salt: m.Salt, CreateTime: m.CreateTime})
		if e != nil {
			u.Data["json"] = &models.RegisterResult{Result: 1, Err: e.Error()}
		} else {
			u.Data["json"] = &models.RegisterResult{Result: 0, Err: "ok"}
		}
	} else {
		u.Data["json"] = &models.RegisterResult{Result: 1, Err: "token或key错误"}
	}

	u.ServeJSON()
}

//用户登陆
func (u *UserController) Login() {
	type Message struct {
		UserName, Password, Token, Key string
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)
	if common.Filter(m.Token, m.Key) {
		login_flag, e := models.GetUserLoginInfo(models.User{UserName: m.UserName, Password: m.Password})
		if login_flag != true {
			u.Data["json"] = &models.LoginResult{Result: 1, Err: e.Error()}
		} else {
			u.Data["json"] = &models.LoginResult{Result: 0, Err: "ok"}
		}
	} else {
		u.Data["json"] = &models.RegisterResult{Result: 1, Err: "token或key错误"}
	}
	u.ServeJSON()
}

//获取用户列表
func (u *UserController) UserList() {
	type Message struct {
		Token, Key string
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)

	if common.Filter(m.Token, m.Key) {
		var users []models.UserList
		users, e := models.GetAllUser()
		if e != nil {
			u.Data["json"] = &models.UserListResult{Result: 1, Err: e.Error()}
		} else {
			u.Data["json"] = &models.UserListResult{Result: 0, Err: "ok", UserList: users, Num: len(users)}
		}
	} else {
		u.Data["json"] = &models.RegisterResult{Result: 1, Err: "token或key错误"}
	}

	u.ServeJSON()
}

//获取用户详情
func (u *UserController) UserInfo() {

	type Message struct {
		Token  string `orm:column(token)`
		Key    string `orm:column(key)`
		UserId int    `orm:column(userId)`
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)
	fmt.Println(m)
	if common.Filter(m.Token, m.Key) {
		var user models.User
		user, e := models.UserInfo(models.User{Id: strconv.Itoa(m.UserId)})
		if e != nil {
			u.Data["json"] = &models.UserInfoResult{Result: 1, Err: e.Error()}
		} else {
			u.Data["json"] = &models.UserInfoResult{Result: 0, Err: "ok", User: user}
		}
	} else {
		u.Data["json"] = &models.UserInfoResult{Result: 1, Err: "token或key错误"}
	}
	u.ServeJSON()
}
