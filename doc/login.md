### 用户登陆



| 更改描述 | 更改时间 | 版本 | 作者 |
|--------|--------|--------|--------|
| 新做成| 2017/03/24  |v1.0|张李杨|


API概述：系统用户注册

API接口：IP/v1/user/login

##### API请求详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  platformId |    int    |是|平台id（1.web；2.安卓，3.ios；4.微信）|
|  token |   string |是|api接入验证，由后台api提供|
|  key |   string |是|api接入验证，由后台api提供|
|  versionId |   string |是|api版本号|
|  qq |   string |否|qq|
|  phone |   string |否|电话号码|
|  email |   string |是|邮箱|
|  userName |   string |否|用户名称|
|  password |   string| 是|用户密码|
|  lastLoginTime |   int |否|当前登陆时间|
|  lastLoginIp |   string |否|当前登陆ip|



##### API返回详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  result |    int    |是|返回结果id，1000表示成功，否则为错误|
|  message |   string |是|错误描述信息|


##### 请求参数示例

```go

{
      platformId：1，
      versionId："v1.0"
      token: "token",
      key: "key",
      userName: "liyoung",
      email:"2587401690@qq.com"
      password: "e10adc3949ba59abbe56e057f20f883e"//密码
      lastLoginTime:"12123123"
      lastLoginIp:"192.168.2.1"
      qq:"258701690"
      phone:"18355519881"
}

```

##### 返回参数示例

```go
    {
       result: 1000,//注册结果，1000表示成功，否则错误（显示对应的错误码）
       message: "success" //错误描述
    }
```

##### 登陆逻辑

对用户传过来的用户名加密码传到服务器端。根据返回的结果判断登陆结果。

