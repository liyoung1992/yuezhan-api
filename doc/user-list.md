### 用户列表

| 更改描述 | 更改时间 | 版本 | 作者 |
|--------|--------|--------|--------|
| 新做成| 2017/03/24  |v1.0|张李杨|


API概述：系统用户注册

API接口：IP/v1/user/user-list

| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  platformId |    int    |是|平台id（1.web；2.安卓，3.ios；4.微信）|
|  token |   string |是|api接入验证，由后台api提供|
|  key |   string |是|api接入验证，由后台api提供|
|  versionId |   string |是|api版本号|
|  qq |   string |否|qq|
|  phone |   string |否|电话号码|
|  email |   string |否|邮箱|
|  userName |   string |否|用户名称|
|  userId |   string |否|用户名称|
|  shift |   int |否|偏移量|
|  point |   int |否|请求开始点|
|  order |   int |否|0：时间倒序，1时间正序|



##### API返回详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  result |    int    |是|返回结果id，1000表示成功，否则为错误|
|  message |   string |是|错误描述信息|
|  num |   int |否|本次返回数量|
|  total |   int |是|符合搜索条件的总数，分页用|
|  userList |  array |否|用户信息|

###### userList 详细参数

| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  qq |   string |否|qq|
|  phone |   string |否|电话号码|
|  email |   string |是|邮箱|
|  userName |   string |否|用户名称|
|  createTime |   int |否|创建时间，时间戳格式|
|  updateTime |   int |否|创建时间，时间戳格式|
|  teamId |   int |否|队伍id|
|  imageUrl |   string |否|头像图片URL|
|  lastLoginTime |   int |否|当前登陆时间|
|  lastLoginIp |   string |否|当前登陆ip|


##### 请求参数示例


```go

{
	platformId：1，
    versionId："v1.0"
    token: "token",
    key: "key",
    userId: 12
    userName:"liyoung",
    qq:"2587401690",
    phone:"18355519881",
    email:"liyoung_1992@163.com"
    shift:15,
    point:0,
    order:0
}

```

##### 返回参数示例

```go

{
    result: 1000,//注册结果，1000表示成功，否则错误（显示对应的错误码）
    message: "success", //错误描述
    num:12,
    total:102,
    userList: [
    {
        id: "12",
        userName:"liyoung",
        qq:"2587401690",
        phone:"18355519881",
        email:"liyoung_1992@163.com",
        createTime:1484893394,
        updateTime:14848933922,
        teamId:10023,
        imageUrl:"https://beego.me/static/img/brands/360.png"
        lastLoginTime:"12123123"
        lastLoginIp:"192.168.2.1"
    },
    {
        id: "13",
        userName:"liyoung",
        qq:"2587401690",
        phone:"18355519881",
        email:"liyoung_1992@163.com",
        createTime:1484893394,
        updateTime:14848933922,
        teamId:10023,
        imageUrl:"https://beego.me/static/img/brands/360.png"
        lastLoginTime:"12123123"
        lastLoginIp:"192.168.2.1"
    }
    ]
}

```