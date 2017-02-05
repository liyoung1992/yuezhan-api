package controllers

import (
	"encoding/json"
	"fmt"
	"yuezhan-api/common"
	"yuezhan-api/models"

	"github.com/astaxie/beego"
)

type TeamController struct {
	beego.Controller
}

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
	fmt.Println(m)
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
