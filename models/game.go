package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	o = orm.NewOrm()
}

type Game struct {
	Id         int    `orm:"coloumn(id)"`
	Name       string `orm:"column(name)"`
	Team1      int    `orm:"column(team1)"`
	Team2      int    `orm:"column(team2)"`
	UserId     int    `orm:"column(userId)"`
	RegionId   int    `orm:"column(regionId)"`
	ImageUrl   string `orm:"column(imageUrl)"`
	StartTime  int    `orm:"column(startTime)"`
	CreateTime int    `orm:"column(CreateTime)"`
	UpdateTime int    `orm:"column(UpdateTime)"`
}

//添加比赛
//func AddGame(g Game) error {
//	_, e := o.Raw("insert into tb_game (name,team1,team2,userId,regionId,imageUrl,startTime,createTime,updateTime) values(?,?,?,?,?,?,?,?,?,?)",
//		g.Name, g.Team1, g.Team2, g.UserId, g.RegionId, g.ImageUrl, g.StartTime, g.CreateTime, g.UpdateTime)
//	if e != nil {
//		return e
//	}
//	return nil
//}

//获取比赛详情
func GameInfo(g Game) (Game, error) {
	var game Game
	e := o.Raw("select * from tb_game where id = ?", g.Id).QueryRow(&game)
	if e != nil {
		return game, e
	}
	return game, nil
}

//获取比赛列表
func GameList() ([]Game, error) {
	var games []Game
	_, e := o.Raw("select * from tb_game").QueryRows(&games)
	if e != nil {
		return nil, e
	}
	return games, nil
}
