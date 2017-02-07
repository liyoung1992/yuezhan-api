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

//获取用户详情
type UserInfoResult struct {
	Result int64
	Err    string
	User   User
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
type TeamInfoResult struct {
	Result int64
	Err    string
	Team   Team
}

//获取战队列表
type TeamListResult struct {
	Result   int64
	Num      int
	Err      string
	TeamList []Team
}
