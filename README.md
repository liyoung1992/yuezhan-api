## 系统后台API列表




| 更改描述 | 更改时间 | 版本 | 作者 |
|--------|--------|--------|--------|
| 新做成| 2017/03/24  |v1.0|张李杨|







* * *

系统api接口列表，调用接口的所有api都必须携带platformId,token,key,versionId.所有的API都走http协议。post方式传递数据。数据格式为json。

-  platformId 平台id（1.web；2.安卓，3.ios；4.微信）；
-  token api接入验证，由后台api提供；
-  key api接入验证，由后台api提供；
-  versionId 请求api的当前版本。



#### 用户模块

1. [用户注册](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/register.md)
2. [用户登陆](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/login.md)
3. [获取用户详情](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/user-info.md)
4. [查询用户列表](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/user-list.md)
5. [更新用户信息](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/user-update.md)
6. [删除用户](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/user-delete.md)

#### 战队模块


1. [添加战队](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/add-team.md)
2. [加入战队](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/join-team.md)
3. [获取战队详情](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/team-info.md)
4. [查询战队列表](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/team-list.md)
5. [更新战队信息](http://)
6. [删除战队](http://)



#### 比赛模块

1. [添加比赛](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/add-team.md)
2. [邀请战队加入比赛](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/join-team.md)
3. [获取比赛详情](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/team-info.md)
4. [查询比赛列表](https://github.com/liyoung1992/yuezhan-api/blob/master/doc/team-list.md)
5. [更新比赛信息](http://)
6. [删除比赛](http://)



