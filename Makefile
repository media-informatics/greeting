.PHONY: build service server client clean dep run

service:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service/service.proto

bin/server.exe: service server/server.go
	go build -o bin/server.exe server/server.go

bin/client.exe: service client/client.go
	go build -o bin/client.exe client/client.go

build: bin/server.exe bin/client.exe


server: bin/server.exe
	./bin/server.exe

client: bin/client.exe
	./bin/client.exe

clean:
	go clean
	rm bin/client.exe
	rm bin/server.exe
	rm service/*.pb.go

dep:
	go mod download