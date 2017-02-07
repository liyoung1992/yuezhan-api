### 战队列表

URL：http://127.0.0.1:8080/v1/team/team-list

1. 请求参数

```go

{
    token: "token",
    key: "key"
}

```

2. 返回参数

```go

{
    Result: 0,
    Num: 3,
    Err: "ok",
    TeamList: [
    {
    Id: 4,
    TeamName: "无敌战队",
    RegionId: 11,
    UserId: 12,
    LeaderId: 0,
    CreateTime: 0,
    ImageUrl: ""
    },
    {
    Id: 5,
    TeamName: "无敌战队2",
    RegionId: 11,
    UserId: 12,
    LeaderId: 0,
    CreateTime: 0,
    ImageUrl: ""
    },
    {
    Id: 6,
    TeamName: "无敌战队3",
    RegionId: 0,
    UserId: 0,
    LeaderId: 0,
    CreateTime: 0,
    ImageUrl: ""
    }
   ]
 }

```