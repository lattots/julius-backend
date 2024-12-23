run-grpc:
	go run cmd/server/server.go

run-client:
	go run cmd/client/client.go

proto-gen:
	protoc --proto_path=proto --go_out=proto --go-grpc_out=proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative event.proto
