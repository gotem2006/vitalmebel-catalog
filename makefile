
export GO111MODULE=on

SERVICE_NAME=catalog
SERVICE_PATH=gotem2006/vitalmebel/catalog


.PHONY:
generate:
	protoc --go_out=pkg/catalog \
		--go-grpc_out pkg/catalog --go-grpc_opt paths=source_relative  \
		--grpc-gateway_out  pkg/catalog  --grpc-gateway_opt paths=source_relative \
		--openapiv2 swagger 
		proto/catalog/catalog.proto  
	mv pkg/catalog/proto/catalog/catalog_grpc.pb.go pkg/catalog/
	mv pkg/catalog/proto/catalog/catalog.pb.gw.go pkg/catalog/
	rmdir -p pkg/catalog/proto/catalog 
	
.PHONY:
generatemocks:
	mockgen -source=internal/repo/repo.go -package service -destination internal/service/service_mocks_test.go Repo

.PHONY: deps-go
deps-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest


.PHONY: build-go
build-go: .build

.build:
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/app$(shell go env GOEXE) ./cmd/app/main.go