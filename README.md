# Chat
是一款使用 Golang 实现的简易 IM 服务器，主要特性：
1. 支持 websocket 接入
2. 单聊、群聊
3. 离线消息同步
4. 支持服务水平扩展

### 扩展安装
``` shell
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
go get -u github.com/golang-jwt/jwt/v4
go get github.com/satori/go.uuid
go get -u github.com/go-sql-driver/mysql
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/spf13/viper
go get github.com/go-redis/redis/v8
 go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
服务端启动：
1. 连接 MySQL，创建 go_chat 库，进入执行 sql/table.sql 文件中 SQL 代码
2. app.yaml 修改配置文件信息
3. main.go 启动服务端

### 模块
1. 登录注册模块
2. 使用Gin搭建Websocket服务
3. 发送接受消息
4. protobuf的使用
