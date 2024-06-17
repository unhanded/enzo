build: protobuf cmd

cmd:
	echo "building go application..."
	go build -o out/enzosrv ./cmd/enzosrv
	go build -o out/enzoctl ./cmd/enzoctl
	go build -o out/enzoitem ./cmd/enzoitem

protobuf: proto-clean
	echo "building protobuf assets..."
	mkdir ./internal/enzo_proto
	protoc --plugin=$(shell which protoc-gen-go) --go_out=./ ./proto/workitem.proto

clean:
	rm -rf ./out
	rm -rf ./internal/enzo_proto

proto-clean:
	echo "Cleaning up proto artifacts..."
	rm -rf ./internal/enzo_proto
