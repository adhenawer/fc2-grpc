# fc2-grpc

#### Install protoc-gen-go && protoc-gen-go-grpc

- https://developers.google.com/protocol-buffers/docs/gotutorial
- https://github.com/golang/protobuf/issues/1053#issuecomment-687834258

#### Build

Follow commands to generate protobuf files 

```bash
$ protoc --proto_path=proto proto/*.proto --go_out=pb
$ protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

#### Run

```bash
$ go run cmd/server/server.go
```