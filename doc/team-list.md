### 战队列表

URL：http://127.0.0.1:8080/v1/team/team-list


| 更改描述 | 更改时间 | 版本 | 作者 |
|--------|--------|--------|--------|
| 新做成| 2017/03/28  |v1.0|张李杨|


API概述：系统用户注册

API接口：IP/v1/team/team-info


##### API请求详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  platformId |    int    |是|平台id（1.web；2.安卓，3.ios；4.微信）|
|  token |   string |是|api接入验证，由后台api提供|
|  key |   string |是|api接入验证，由后台api提供|
|  versionId |   string |是|api版本号|
|  teamName |   string |否|队伍名字|
|  leaderId |   int |否|队伍名字|
|  createTime |   int |否|队伍名字|
|  regionId |   int |否|队伍名字|
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
|  teamList |   array |否|战队信息|

###### teamInfo详细参数信息

| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  id |    int    |是|战队id|
|  teamName |   string |是|错误描述信息|
|  regionId |   int |否|战队信息|
|  userId |   int |否|战队信息|
|  leaderId |   int |否|战队信息|
|  createTime |   int |否|战队信息|
|  updateTime |   int |否|战队信息|
|  imageUrl |   string |否|战队信息|



1.请求参数

```go

{
    platformId：1，
    versionId："v1.0"
    token: "token",
    key: "key",
    teamName: "无敌战队",
    regionId: 11,
    userId: 12,
    leaderId: 121,
    createTime: 1486284225,
    shift:15,
    point:0,
    order:0
}

```

2.返回参数

```go

{
    result: 1000,//注册结果，1000表示成功，否则错误（显示对应的错误码）
    message: "success" //错误描述
    num:12,
    total:102,
    TeamList: [
    {
       id: 2,
       teamName: "无敌战队",
       regionId: 11,
       userId: 12121,
       leaderId: 12401,
       createTime: 12121312,
       updateTime：12122343,
       imageUrl: ""
    },
    {
       id: 3,
       teamName: "无敌战队",
       regionId: 11,
       userId: 12121,
       leaderId: 12401,
       createTime: 12121312,
       updateTime：12122343,
       imageUrl: ""
    },
    {
       id: 4,
       teamName: "无敌战队",
       regionId: 11,
       userId: 12121,
       leaderId: 12401,
       createTime: 12121312,
       updateTime：12122343,
       imageUrl: ""
    }
   ]
 }

```