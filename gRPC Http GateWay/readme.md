Install `vscode-proto3` extenstion at your VS code
Install `protoc` binary from https://github.com/protocolbuffers/protobuf/releases
Install the `protoc` library:
https://developers.google.com/protocol-buffers/docs/reference/go-generated
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```
At thr proto file define the Go package it will belongs to, as:
```bash
option go_package = "example.com/project/protos/fizz";
```
In our case, we have in top of the file this line:
```bash
option go_package = "/stream/server";
```
Be careful in the code above, The import path must contain at least one period ('.') or forward slash ('/') character.

SO, in our case, our compilation line will become:
Generate the messages
```bash
protoc --proto_path=proto --go_out=./server --go_opt=Mproto/data.proto=proto --go_out=./client --go_opt=Mproto/data.proto=proto data.proto --go-grpc_out=./server
```
Generate the services
```bash
protoc --proto_path=proto --go-grpc_out=./server --go-grpc_out=./client data.proto 
```
Generate the Gateway
```bash
protoc -I ./proto \
  --go_out ./proto --go_opt paths=source_relative \
  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
  ./proto/helloworld/hello_world.proto
  ```
Some details about the compilation line are below:
Using the `M` argument `M${PROTO_FILE}=${GO_IMPORT_PATH}` in the command line, as:
```bash
--go_opt=Mprotos/data.proto=stream/server
```
The `M` option override whatever defined in the proto file itself
The `M` option can be used multiple times, so that multiple versions of the `.pb.go` file are created
Convert the `data.proto` to :
```bash
protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
```
that is:
```bash
protoc --proto_path=protos --go_out=protos --go_opt=paths=source_relative data.proto
```

protoc --proto_path=protos --go_out=plugins=grpc:. --go_opt=paths=source_relative /protos/data.proto

protoc --proto_path=protos --go-grpc_out=./server --go_opt=Mprotos/data.proto=protos --go-grpc_out=./client --go_opt=Mprotos/data.proto=protos data.proto 

**********
To  create http gateway to thr gRPC, https://grpc-ecosystem.github.io/grpc-gateway/
add `import "google/api/annotations.proto";` to thr `proto` file