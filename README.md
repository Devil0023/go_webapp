# go_webapp

一个依照MVC模式封装的Golang Web Api框架。



## 运行

```$xslt
# 安装govendor
go get -u github.com/kardianos/govendor

# 拉取依赖
govendor sync

# 本地运行
go run main.go

```

## 项目结构

```$xslt
.
├── Dockerfile         // Dockerfile
├── Makefile           // Makefile
├── README.md
├── conf               // 配置文件
│   ├── app.dev.ini
│   └── app.ini
├── controller         // Controller层
│   └── api
├── cron               // Crontab服务
│   └── cron.go
├── main.go            // main
├── middleware         // 中间件
│   └── jwt
├── models             // Model层
│   ├── auth.go
│   ├── models.go
│   └── tag.go
├── pkg                // 工具库
│   ├── app
│   ├── code
│   ├── csv
│   ├── gfile
│   ├── ghttp
│   ├── gredis
│   ├── logging
│   ├── setting
│   └── util
├── routers            // 路由层
│   └── router.go
├── rpc                // RPC服务
│   ├── proto
│   ├── protofile
│   ├── rpc.go
│   └── services
├── server             // Server服务
│   └── server.go
├── storage            // 文件存放
│   └── logs
└── vendor             // vendor
    └── vendor.json

```