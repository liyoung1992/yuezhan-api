### 1.用户注册


1.    请求参数

   ```go
    {
       "UserName":"liyoung",//用户名
       "Salt":"4b9iWOU6qeECN8h",//salt
       "Password":"574aafa475e4a00ba95196dae7860089",//密码
       "CreateTime":1484893394 //创建时间（unix时间戳//注册结果）
       }
   ```

2.    返回参数

   ```go
    {
    "Result": 0,//注册结果，0表示成功，否则错误（显示对应的错误码）
    "Err": "ok" //错误描述
    }
   ```

3.    密码设计逻辑:

   1.加密逻辑
    Password=md5(password+salt)
   2.解密逻辑
    根据传过来的用户密码和salt在进行md5一次，然后和数据password进行对比。


