### 获取用户详情

URL：http://127.0.0.1:8080/v1/user/user-info

1. 请求参数

```go

{
    token: "token",
    key: "key",
    userId: 12
}

```

2. 返回参数

```go

{
    Result: 0,
    Err: "ok",
    User: {
      Id: "12",
      UserName: "2587401690@qq.com",
      Password: "",
      Salt: "",
      CreateTime: 0,
      UpdateTime: 0,
      TeamId: 0,
      LastLoginTime: 0
    }
}

```