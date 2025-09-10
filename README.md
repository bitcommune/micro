# 安装和使用
1. 安装工具
   go install github.com/bitcommune/micro@latest
2. 初始化新项目
   micro init my-microservice
   cd my-microservice
3. 创建 proto 文件
   micro proto hello --package=myapp --gopackage=github.com/bitcommune/my-microservice/pkg/pb
4. 创建服务
   micro service hello
5. 编译 proto 文件（需要安装 protoc）


# 手动执行protoc命令编译proto文件
protoc --go_out=. --go-grpc_out=. api/proto/hello.proto

```
micro/
├── cmd/
│   ├── root.go
│   ├── init.go
│   ├── service.go
│   ├── proto.go
│   ├── docker.go
│   └── database.go
├── internal/
│   ├── generator/
│   │   ├── project.go
│   │   ├── proto.go
│   │   ├── service.go
│   │   ├── docker.go
│   │   ├── database.go
│   │   ├── config.go
│   │   ├── middleware.go
│   │   └── test.go
│   ├── utils/
│   │   ├── file.go
│   │   ├── template.go
│   │   └── protoc.go
│   └── types/
│       └── config.go
├── templates/
│   ├── project/
│   │   ├── main.go.tmpl
│   │   ├── config.yaml.tmpl
│   │   ├── dockerfile.tmpl
│   │   ├── docker-compose.yaml.tmpl
│   │   └── Makefile.tmpl
│   ├── proto/
│   │   └── service.proto.tmpl
│   ├── service/
│   │   ├── service.go.tmpl
│   │   ├── handler.go.tmpl
│   │   └── repository.go.tmpl
│   ├── middleware/
│   │   ├── logging.go.tmpl
│   │   ├── auth.go.tmpl
│   │   └── recovery.go.tmpl
│   ├── database/
│   │   ├── model.go.tmpl
│   │   ├── repository.go.tmpl
│   │   └── migration.sql.tmpl
│   └── test/
│       ├── service_test.go.tmpl
│       └── handler_test.go.tmpl
├── go.mod
├── go.sum
└── main.go
```


