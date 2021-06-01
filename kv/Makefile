compile:
		protoc  .\kv\protos\*.proto  \
		--go_out=kv\api\ \
		--go_opt=paths=source_relative 

test:
	go test -race ./...		

