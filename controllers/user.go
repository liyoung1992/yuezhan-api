package controllers

import (
	"encoding/json"
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
		UserName, Password, Salt, Token, Key string
		CreateTime                           int
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

//// @Title CreateUser
//// @Description create users
//// @Param	body		body 	models.User	true		"body for user content"
//// @Success 200 {int} models.User.Id
//// @Failure 403 body is empty
//// @router / [post]
//func (u *UserController) Post() {
//	var user models.User
//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//	uid := models.AddUser(user)
//	u.Data["json"] = map[string]string{"uid": uid}
//	u.ServeJSON()
//}

//// @Title GetAll
//// @Description get all Users
//// @Success 200 {object} models.User
//// @router / [get]
//func (u *UserController) GetAll() {
//	users := models.GetAllUsers()
//	u.Data["json"] = users
//	u.ServeJSON()
//}

//// @Title Get
//// @Description get user by uid
//// @Param	uid		path 	string	true		"The key for staticblock"
//// @Success 200 {object} models.User
//// @Failure 403 :uid is empty
//// @router /:uid [get]
//func (u *UserController) Get() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		user, err := models.GetUser(uid)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = user
//		}
//	}
//	u.ServeJSON()
//}

//// @Title Update
//// @Description update the user
//// @Param	uid		path 	string	true		"The uid you want to update"
//// @Param	body		body 	models.User	true		"body for user content"
//// @Success 200 {object} models.User
//// @Failure 403 :uid is not int
//// @router /:uid [put]
//func (u *UserController) Put() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		var user models.User
//		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//		uu, err := models.UpdateUser(uid, &user)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = uu
//		}
//	}
//	u.ServeJSON()
//}

//// @Title Delete
//// @Description delete the user
//// @Param	uid		path 	string	true		"The uid you want to delete"
//// @Success 200 {string} delete success!
//// @Failure 403 uid is empty
//// @router /:uid [delete]
//func (u *UserController) Delete() {
//	uid := u.GetString(":uid")
//	models.DeleteUser(uid)
//	u.Data["json"] = "delete success!"
//	u.ServeJSON()
//}

//// @Title Login
//// @Description Logs user into the system
//// @Param	username		query 	string	true		"The username for login"
//// @Param	password		query 	string	true		"The password for login"
//// @Success 200 {string} login success
//// @Failure 403 user not exist
//// @router /login [get]
//func (u *UserController) Login() {
//	username := u.GetString("username")
//	password := u.GetString("password")
//	if models.Login(username, password) {
//		u.Data["json"] = "login success"
//	} else {
//		u.Data["json"] = "user not exist"
//	}
//	u.ServeJSON()
//}

//// @Title logout
//// @Description Logs out current logged in user session
//// @Success 200 {string} logout success
//// @router /logout [get]
//func (u *UserController) Logout() {
//	u.Data["json"] = "logout success"
//	u.ServeJSON()
//}
