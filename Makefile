build:
	rm -rf ./out
	go build -o ./out/enzoctl ./cmd/enzoctl
	go build -o ./out/enzoitem ./cmd/enzoitem
	go build -o ./out/enzod ./cmd/enzod
