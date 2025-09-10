```
microgen/
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