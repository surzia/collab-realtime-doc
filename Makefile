proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/protobuf/*.proto

grpcserver:
	go build -o server . && ./server

client:
	cd cmd/client && go build -o client . && ./client

webserver:
	cd web && npm run start

protojs:
	protoc server/protobuf/*.proto --js_out=import_style=commonjs:./web/src/newspb --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web/src/newspb