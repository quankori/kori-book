# Labs gRPC with Golang

## Install

Protoc CLI

```bash
https://grpc.io/docs/protoc-installation/
```

Install protoc-gen-go

```go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Maybe have error about path, need add in .zshrc

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Build package

```bash
protoc --go_out=services --go_opt=paths=source_relative \
    --go-grpc_out=services --go-grpc_opt=paths=source_relative \
    proto/message.proto
```

grpcurl test 

```bash
grpcurl -plaintext localhost:50051 list
```

```bash
grpcurl -plaintext localhost:50051 list services.StreamService
```

```bash
grpcurl -plaintext localhost:50051 describe services.StreamRequest
grpcurl -plaintext localhost:50051 describe services.StreamResponse
```

```bash
grpcurl -plaintext -d '{"message": "gopher"}' localhost:50051 services.StreamService/BidirectionalStream
```