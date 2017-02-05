package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	o = orm.NewOrm()
}

type Team struct {
	Id         int    `orm:"column(id)"`
	TeamName   string `orm:"column(teamName)"`
	RegionId   int    `orm:"column(regionId)"`
	UserId     int    `orm:"column(userId)"`
	LeaderId   int    `orm:column(leaderId)`
	CreateTime int    `orm:column(createTime)`
	ImageUrl   string `orm:column(imageUrl)`
}

//添加战队
func AddTeam(t Team) error {
	_, e := o.Raw("insert into tb_team (teamName,regionId,userId,leaderId,createTime,imageUrl) values(?,?,?,?,?,?);",
		t.TeamName, t.RegionId, t.UserId, t.LeaderId, t.CreateTime, t.ImageUrl).Exec()
	if e != nil {
		return e
	}
	return nil
}
