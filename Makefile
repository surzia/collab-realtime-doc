proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/*.proto

server:
	cd cmd/server && go build -o server . && ./server

client:
	cd cmd/client && go build -o client . && ./client