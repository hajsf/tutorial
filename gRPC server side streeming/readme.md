Install `vscode-proto3` extenstion at your VS code
Install `protoc` binary from https://github.com/protocolbuffers/protobuf/releases
Install the `protoc` library:
https://developers.google.com/protocol-buffers/docs/reference/go-generated

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```
At the proto file define the Go package it will belongs to, as:
```bash
option go_package = "example.com/project/protos/fizz";
```
In our case, we have in top of the file this line:
```bash
option go_package = "/stream/server";
```
Be careful in the code above, The import path must contain at least one period ('.') or forward slash ('/') character.

To  create http gateway to the gRPC, (see gRPCgateway.png photo and https://grpc-ecosystem.github.io/grpc-gateway/)
add `import "google/api/annotations.proto";` to thr `proto` file, also you need to add the `option` block to the `rpc service`
```proto
    option (google.api.http) = {
      post: "/v1/api"
      body: "*"
    };
```
If `get` is used, then no need for `body`:
```proto
    option (google.api.http) = {
      get: "/v1/api"
    };
```
It is good practice to use `version` in the `API` for maintability.

`protoc` is required to generate the proto `message` file and the `grpc` connection as well as (if required) the `gateway` for http. It is recommended to generate 3 of them.

```bash
protoc -I ./proto \
  --go_out ./server/proto --go_opt paths=source_relative \
  --go_out ./client/proto --go_opt paths=source_relative \
  --go-grpc_out ./server/proto --go-grpc_opt paths=source_relative \
  --go-grpc_out ./client/proto --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./server/proto --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_out ./client/proto --grpc-gateway_opt paths=source_relative \
  ./proto/data.proto
  ```