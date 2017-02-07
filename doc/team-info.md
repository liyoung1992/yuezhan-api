### 获取战队详情

URL：http://127.0.0.1:8080/v1/team/team-info

1. 请求参数

```go

{
    token: "token",
    key: "key",
    teamId: 3
}

```

2. 返回参数

```go

{
    Result: 0,
    Err: "ok",
    Team: {
    Id: 3,
    TeamName: "无敌战队",
    RegionId: 11,
    UserId: 12,
    LeaderId: 0,
    CreateTime: 0,
    ImageUrl: ""
    }
}

```