package models

// Result as a return json template.
type VerificateResult struct {
	Result int64  `json:"result"`
	Err    string `json:"err"`
}

//注册返回结果
type RegisterResult struct {
	Result int64  `json:"result"`
	Err    string `json:"err"`
}

//登陆结果
type LoginResult struct {
	Result int64  `json:"result"`
	Err    string `json:"err"`
}

//用户列表返回结果
type UserListResult struct {
	Result   int64      `json:"result"`
	Err      string     `json:"err"`
	Num      int        `json:"num"`
	UserList []UserList `json:"userList"`
}

//获取用户详情
type UserInfoResult struct {
	Result int64  `json:"result"`
	Err    string `json:"err"`
	User   User   `json:"user"`
}

//战队
type AddTeamResult struct {
	Result int64  `json:"result"`
	Err    string `json:"err"`
}

//加入战队结果
type JoinTeamResult struct {
	Result int64  `json:"result"`
	Err    string `json:"err"`
}

//获取战队详情
type TeamInfoResult struct {
	Result int64  `json:"result"`
	Err    string `json:"err"`
	Team   Team   `json:"team"`
}

//获取战队列表
type TeamListResult struct {
	Result   int64  `json:"result"`
	Err      string `json:"err"`
	TeamList []Team `json:"teamList"`
	Num      int    `json:"num"`
}
