# WeChatApi 本项目 基于 gin框架
### 前端项目：
##### https://github.com/lwytry/WeChat

## 高仿微信计划：
### 已经实现功能
1. 登录控制器
    * 登录、注册、发送验证码(未接入)、更新token
2. 联系人控制器
    * 获取联系人列表
3. 聊天控制器
    * 发送消息、拉取消息
4. wesocket服务
    * 用户登录、发送消息、接收消息
### 待实现功能
根据前端项目实现对应功能支持
## 项目主要使用的第三方库
* [Jwt-go](https://github.com/dgrijalva/jwt-go)：jwt存储用户登录信息
* [Redis](https://github.com/gomodule/redigo)：redis数据库
* [Gorm](https://github.com/jinzhu/gorm)：数据库orm框架
* [Websocket](https://github.com/gorilla/websocket)：socket通信