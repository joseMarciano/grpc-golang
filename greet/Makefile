

generate:
	@echo "Generating gRPC code from .proto files..."
	mkdir -p protogen
	protoc \
	  -I=prototypes \
	  --go_out=./protogen --go_opt=paths=source_relative \
	  --go-grpc_out=./protogen --go-grpc_opt=paths=source_relative \
	  greet.proto