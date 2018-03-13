# @see https://libetrada.com/grpc_golang/
all:

set_up_server:
	go run server/main.go
exec:
	go run client/main.go
