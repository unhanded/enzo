build: pb go


go:
	go build -o bin/server ./app/server

pb: go
	mkdir ./internal/enzo_proto
	protoc --plugin=$(shell which protoc-gen-go) --go_out=./ ./proto/workitem.proto

clean:
	rm -rf ./bin
	rm -rf ./internal/enzo_proto
