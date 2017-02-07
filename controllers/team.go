package controllers

import (
	"encoding/json"
	//	"fmt"
	"yuezhan-api/common"
	"yuezhan-api/models"

	"github.com/astaxie/beego"
)

type TeamController struct {
	beego.Controller
}

//添加战队
func (u *TeamController) AddTeam() {
	type Message struct {
		Id         int    `orm:"column(id)"`
		TeamName   string `orm:"column(teamName)"`
		RegionId   int    `orm:"column(regionId)"`
		UserId     int    `orm:"column(userId)"`
		LeaderId   int    `orm:column(leaderId)`
		CreateTime int    `orm:column(createTime)`
		ImageUrl   string `orm:column(imageUrl)`
		Token      string `orm:column(token)`
		Key        string `orm:column(key)`
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)

	if common.Filter(m.Token, m.Key) {
		e := models.AddTeam(models.Team{TeamName: m.TeamName, RegionId: m.RegionId, UserId: m.UserId, LeaderId: m.LeaderId, CreateTime: m.CreateTime,
			ImageUrl: m.ImageUrl})
		if e != nil {
			u.Data["json"] = &models.AddTeamResult{Result: 1, Err: e.Error()}
		} else {
			u.Data["json"] = &models.AddTeamResult{Result: 0, Err: "ok"}
		}
	} else {
		u.Data["json"] = &models.AddTeamResult{Result: 1, Err: "token或key错误"}
	}
	u.ServeJSON()
}

//加入战队
func (u *TeamController) JoinTeam() {
	type Message struct {
		Token  string `orm:column(token)`
		Key    string `orm:column(key)`
		UserId int    `orm:column(userId)`
		TeamId int    `orm:column(teamId)`
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)

	if common.Filter(m.Token, m.Key) {
		e := models.JoinTeam(models.Team{Id: m.TeamId, UserId: m.UserId})
		if e != nil {
			u.Data["json"] = &models.JoinTeamResult{Result: 1, Err: e.Error()}
		} else {
			u.Data["json"] = &models.JoinTeamResult{Result: 0, Err: "ok"}
		}
	} else {
		u.Data["json"] = &models.JoinTeamResult{Result: 1, Err: "token或key错误"}
	}
	u.ServeJSON()
}

func (u *TeamController) TeamInfo() {
	type Message struct {
		Token  string `orm:column(token)`
		Key    string `orm:column(key)`
		TeamId int    `orm:column(teamId)`
	}
	var m Message
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)

	if common.Filter(m.Token, m.Key) {
		var t models.Team
		t, e := models.TeamInfo(models.Team{Id: m.TeamId})
		if e != nil {
			u.Data["json"] = &models.TeamInfoResult{Result: 1, Err: e.Error()}
		} else {
			u.Data["json"] = &models.TeamInfoResult{Result: 0, Err: "ok", Team: t}
		}
	} else {
		u.Data["json"] = &models.TeamInfoResult{Result: 1, Err: "token或key错误"}
	}
	u.ServeJSON()
}
func (t *TeamController) TeamList() {
	type Message struct {
		Token string `orm:column(token)`
		Key   string `orm:column(key)`
	}
	var m Message
	json.Unmarshal(t.Ctx.Input.RequestBody, &m)

	if common.Filter(m.Token, m.Key) {
		var teams []models.Team
		teams, e := models.TeamList()
		if e != nil {
			t.Data["json"] = &models.TeamListResult{Result: 1, Err: e.Error()}
		} else {
			t.Data["json"] = &models.TeamListResult{Result: 0, Err: "ok", TeamList: teams, Num: len(teams)}
		}
	} else {
		t.Data["json"] = &models.TeamListResult{Result: 1, Err: "token或key错误"}
	}
	t.ServeJSON()
}
