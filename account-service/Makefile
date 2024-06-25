.PHONY: run app local
local:
	@echo "Running the app locally"
	@cd cmd && go run main.go

.PHONY generate proto:
protoc:
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.0
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@echo "Generating protobuf files..."
	@protoc -I ./proto --go_out=./ \
			--go-grpc_out=require_unimplemented_servers=false:./ \
			--grpc-gateway_out . --grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt generate_unbound_methods=true \
			./proto/account/*.proto ./proto/google/*.proto
	@echo "Done"