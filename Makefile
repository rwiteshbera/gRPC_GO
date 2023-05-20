gen:
	protoc --proto_path=proto --go_out=. --go-grpc_out=. ./proto/process.proto

clean:
	rm pb/*.go
