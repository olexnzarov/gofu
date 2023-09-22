DAEMON_MAIN = cmd/gofudaemon/main.go
DAEMON_BIN = build/gofudaemon

run-tests:
	go test -v ./...

build-daemon:
	@(echo "Building gofu-daemon binary...")
	go build -o $(DAEMON_BIN) $(DAEMON_MAIN)

run-daemon:
	go run $(DAEMON_MAIN)

install-gofu:
	go install ./cmd/gofu

run-grpc-ui:
	grpcui \
		-plaintext \
		-proto ./pb/process_manager.proto \
		-port 50055 \
		localhost:50051

build-proto:
	rm -rf ./pb/*.pb.go
	protoc ./pb/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative 
	ls ./pb/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
