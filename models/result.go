package models

// Result as a return json template.
type RegisterResult struct {
	Result int64
	Err    string
}

type LoginResult struct {
	Result int64
	Err    string
}
type UserListResult struct {
	Result   int64
	Num      int
	Err      string
	UserList []UserList
}

//战队
type AddTeamResult struct {
	Result int64
	Err    string
}

//加入战队结果
type JoinTeamResult struct {
	Result int64
	Err    string
}

//获取战队详情
type TeamInofoResult struct {
	Id         int    `orm:"column(id)"`
	TeamName   string `orm:"column(teamName)"`
	RegionId   int    `orm:"column(regionId)"`
	UserId     int    `orm:"column(userId)"`
	LeaderId   int    `orm:column(leaderId)`
	CreateTime int    `orm:column(createTime)`
	ImageUrl   string `orm:column(imageUrl)`
}
