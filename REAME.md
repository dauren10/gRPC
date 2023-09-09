$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
protoc --go_out=. --go-grpc_out=. proto/greet.proto