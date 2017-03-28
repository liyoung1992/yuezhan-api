### 更新战队信息


| 更改描述 | 更改时间 | 版本 | 作者 |
|--------|--------|--------|--------|
| 新做成| 2017/03/28  |v1.0|张李杨|


API概述：更新战队信息

API接口：IP/v1/team/team-update




##### API请求详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  platformId |    int    |是|平台id（1.web；2.安卓，3.ios；4.微信）|
|  token |   string |是|api接入验证，由后台api提供|
|  key |   string |是|api接入验证，由后台api提供|
|  versionId |   string |是|api版本号|
|  regionId |   int |否|大区id|
|  userId |   int |否|创建用户id，默认是队长|
|  leaderId |   int |否|队长id|
|  createTime |   int |否|创建时间，时间戳格式|
|  imageUrl |   string |否|头像图片URL|


##### API返回详细参数信息：


| 字段名称 | 类型 |是否必须|描述|
|--------|--------|--------|--------|
|  result |    int    |是|返回结果id，1000表示成功，否则为错误|
|  message |   string |是|错误描述信息|
|  teamId |   int |是|更新后的战队id|


##### 请求参数示例


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
    imageUrl: "1212111111"
}

```

##### 返回参数示例


```go

{
      result: 1000,//注册结果，1000表示成功，否则错误（显示对应的错误码）
      message: "success" //错误描述
      teamId:10034
}

```