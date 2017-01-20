package models

import (
	//	"fmt"
	//	"math/rand"
	//	"strconv"
	//	"time"

	"github.com/astaxie/beego/orm"
)

var (
	o orm.Ormer
)

func init() {
	o = orm.NewOrm()
}

type User struct {
	Id            string
	UserName      string
	Password      string
	Salt          string
	CreateTime    int
	UpdateTime    int
	LastLoginTime int
}

func AddUser(u User) error {
	_, e := o.Raw("insert into tb_user (userName,password,salt,createTime) values (?,?,?,?);", u.UserName, u.Password, u.Salt, u.CreateTime).Exec()
	if e != nil {
		return e
	}
	return nil
}
