### 用户登陆

URL：http://127.0.0.1:8080/v1/user/login

1. 请求参数

```go

{
	Token: "token",
    Key: "key",
    UserName: "liyoung",//用户名
    Password: "e10adc3949ba59abbe56e057f20f883e"//密码
}

```

2. 返回参数

```go

    "Result": 0,//注册结果，0表示成功，否则错误（显示对应的错误码）
    "Err": "ok" //错误描述

```

3. 登陆逻辑

对用户传过来的用户名加密码传到服务器端。根据返回的结果判断登陆结果。

