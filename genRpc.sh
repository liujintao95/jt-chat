goctl rpc protoc ./apps/user/rpc/user.proto --go_out=./apps/user/rpc --go-grpc_out=./apps/user/rpc --zrpc_out=./apps/user/rpc

goctl rpc protoc ./apps/message/rpc/message.proto --go_out=./apps/message/rpc --go-grpc_out=./apps/message/rpc --zrpc_out=./apps/message/rpc