# Compile ProtoBuf & gRPC

Please do follow command in the root of this project.

### Compile Message
```
protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/*.proto
```

### Compile Service
```
protoc --go-grpc_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/*.proto
```