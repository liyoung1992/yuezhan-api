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

//用户加入战队
func JoinTeam(t Team) error {
	_, e := o.Raw("update `tb_user` set `teamId` = ? where id = ?", t.Id, t.UserId).Exec()
	if e != nil {
		return e
	}
	return nil
}

//获取战队详情

func TeamInfo(t Team) (Team, error) {
	var team Team
	e := o.Raw("select * from tb_team where id = ?;", t.Id).QueryRow(&team)
	if e != nil {
		return team, e
	}
	return team, nil
}

//获取战队列表
func TeamList() ([]Team, error) {
	var teams []Team
	_, e := o.Raw("select * from tb_team").QueryRows(&teams)
	if e != nil {
		return nil, e
	}
	return teams, nil
}
