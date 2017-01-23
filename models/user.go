package models

import (
	"encoding/hex"
	"fmt"
	"time"
	//	"fmt"
	//	"math/rand"
	//	"strconv"
	//	"time"
	"crypto/md5"

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
type UserList struct {
	Id       string `orm:"column(id)"`
	UserName string `orm:"column(userName)"`

	CreateTime    int `orm:"column(createTime)"`
	UpdateTime    int `orm:"column(updateTime)"`
	LastLoginTime int `orm:"column(lastLoginTime)"`
}

//添加用户
func AddUser(u User) error {
	_, e := o.Raw("insert into tb_user (userName,password,salt,createTime) values (?,?,?,?);", u.UserName, u.Password, u.Salt, u.CreateTime).Exec()
	if e != nil {
		return e
	}
	return nil
}

//查询用户是否可以登陆
func GetUserLoginInfo(u User) (bool, error) {
	var pw, salt string
	e := o.Raw("select `password`,`salt` from `tb_user` where `userName` = ?;", u.UserName).QueryRow(&pw, &salt)

	if e != nil {
		return false, e
	}
	//	md5salt := GetMd5Password(salt)
	fmt.Println(GetMd5Password(u.Password + salt))
	if pw != GetMd5Password(u.Password+salt) {
		return false, fmt.Errorf("wrong password")
	}
	//可以登陆，更新登陆时间
	_, e = o.Raw("update `tb_user` set `lastLoginTime` = ? where `userName` = ?", time.Now(), u.UserName).Exec()

	if e != nil {
		return false, e
	}
	return true, nil
}

//查询所有用户列表
func GetAllUser() ([]UserList, error) {
	var userList []UserList
	_, e := o.Raw("SELECT `id`,`userName`,`createTime`,`updateTime`,`lastLoginTime` FROM `tb_user`;").QueryRows(&userList)
	fmt.Println(userList)
	if e != nil {
		return nil, e
	}
	return userList, nil
}

//获取md5加密后的密码
func GetMd5Password(ps string) string {
	h := md5.New()
	h.Write([]byte(ps))

	return hex.EncodeToString(h.Sum(nil))
}
