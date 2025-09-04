# gRPC Go Examples

Collection of gRPC service implementations in Go demonstrating various patterns and features.

## Prerequisites

```bash
# Install Go (1.19+)
go version

# Install protoc
brew install protobuf  # macOS

# Install Go protobuf plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Setup

```bash
go mod init github.com/joseMarciano/grpc
go mod tidy
```

## Services

### Greet Service

Simple greeting service example.

**Generate protobuf code:**
```bash
cd greet && make generate
```

**Run server:**
```bash
cd greet/server && go run *.go
```

**Run client:**
```bash
cd greet/client && go run *.go
```

## Troubleshooting

- **protoc not found**: Install Protocol Buffers compiler
- **protoc-gen-go not found**: Ensure `$GOPATH/bin` is in PATH
- **Connection refused**: Start server before client
- **Port in use**: Kill existing process or change port