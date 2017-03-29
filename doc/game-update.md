### 添加比赛


| 更改描述 | 更改时间 | 版本 | 作者 |
|--------|--------|--------|--------|
| 新做成| 2017/03/29  |v1.0|张李杨|


API概述：更新比赛信息。

API接口：IP/v1/game/game-update




##### API请求详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  platformId |    int    |是|平台id（1.web；2.安卓，3.ios；4.微信）|
|  token |   string |是|api接入验证，由后台api提供|
|  key |   string |是|api接入验证，由后台api提供|
|  versionId |   string |是|api版本号|
|  name |   string |是|比赛名称|
|  team1 |   int |否|队伍1的id|
|  team2 |   int |否|队伍2的id|
|  userId |  int|是|创建人id|
|  createTime |   int |否|创建时间|
|  startTime |   int |否|比赛开始时间|
|  endTime |   int |否|创建时间|
|  regionId |   int |否|大区id|







##### API返回详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  result |    int    |是|返回结果id，1000表示成功，否则为错误|
|  message |   string |是|错误描述信息|
|  gameId |   int |是|比赛id|


##### 请求参数示例


```go

{

    platformId：1，
    versionId："v1.0"
    token: "token",
    key: "key",
    name: "扭曲-12-3比赛",
    team1: 11,
    team2: 12,
    userId: 121,
    createTime: 1486284225,
    startTime: "1212111111"
    endTime:12131121,
    regionId:14
}

```

##### 返回参数示例



```go

{
      result: 1000,//注册结果，1000表示成功，否则错误（显示对应的错误码）
      message: "success" //错误描述
      gameId:10034
}

```