### 用户列表

URL：http://127.0.0.1:8080/v1/user/user-list

1. 请求参数


```go

{
	Token: "token",
    Key: "key",
}

```

2. 返回参数

```go

{
    Result: 0,
    Num: 2,
    Err: "ok",
    UserList: [
    {
    Id: "10",
    UserName: "2587401690@qq.com",
    CreateTime: 1485053045,
    UpdateTime: 0,
    LastLoginTime: 2017
    },
    {
    Id: "11",
    UserName: "121",
    CreateTime: 12,
    UpdateTime: 0,
    LastLoginTime: 0
    }
    ]
}

```