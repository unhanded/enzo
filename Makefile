build: pb go


go:
	echo "building go application..."
	go build -o build/server ./app/server

pb: go proto-clean
	echo "building protobuf assets..."
	mkdir ./internal/enzo_proto
	protoc --plugin=$(shell which protoc-gen-go) --go_out=./ ./proto/workitem.proto

clean:
	rm -rf ./build
	rm -rf ./internal/enzo_proto

proto-clean:
	echo "Cleaning up proto artifacts..."
	rm -rf ./internal/enzo_proto
