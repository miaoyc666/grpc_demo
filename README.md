# grpc_demo
grpc demo


### 使用指南（golang）
#### protoc和protoc-gen-go
```bash
# protoc
protoc去github下载安装，自行配置环境变量
# protoc-gen-go、protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### build and run
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative string_service.proto
go mod init demo
go mod tidy
go run server.go
go run client.go
```