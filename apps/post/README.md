# Post Service

Content management and post handling service using Domain-Driven Design and CQRS patterns.

## Overview

The Post Service is responsible for managing all post-related content, including historical posts, articles, and user-generated content. Built with Go using modern architectural patterns:

- **Domain-Driven Design (DDD)** - Clear business domain boundaries
- **Dependency Injection** - Using uber-go/fx framework
- **gRPC** - Internal service communication

## Technology Stack

- **Language:** Go 1.21+
- **Framework:** Custom Go (no heavy framework)
- **Database:** PostgreSQL
- **Communication:** gRPC, TCP, RabbitMQ
- **Logging:** Zap (structured logging)
- **Configuration:** Viper
- **Dependency Injection:** uber-go/fx
- **Validation:** go-playground/validator

## Quick Start

### Prerequisites

- Go 1.21+ installed
- PostgreSQL running
- Environment variables configured

### Installation

```bash
cd apps/post

# Download dependencies
go mod tidy

# Verify dependencies
go mod verify
```

### Configuration

Copy `.env.example` to `.env` and update values:

```bash
cp .env.example .env
```

### Running

```bash
# Development with hot reload (requires air)
make run

# Direct run
go run ./cmd/main.go

# Build
make build

# Production
./bin/server
```

Service will start on gRPC port `:6005`

## Architecture

### Project Structure

```
apps/post/
├── bin/                      # Compiled binaries
├── cmd
│   └── server
│       └── main.go           # Entry point
├── internal
│   ├── posts
│   │   ├── application       # Application layer
│   │   ├── domain            # Domain layer
│   │   ├── infrastructure    # Infrastructure layer
│   │   |   ├── cache         # Outbound Redis repo
│   │   |   ├── persistence   # Outbound PostgreSQL repo
│   │   |   ├── messaging     # Inbound/Outbound RabbitMQ consumer/producer
│   │   |   └── transport     # Inbound gRPC
│   │   └── posts.module.go
│   └── shared
│       ├── app               # Wire dependencies
│       ├── grpc/genproto     # Generated protobuf files
│       └── infrastructure    # Provide infrastructure modules
├── Makefile                  # Script shortcuts
├── .air.toml                 # Hot reload config
└── README.md
```

---

## Code Generation

### Generate gRPC Code

```bash
# Run on project root
# Generate Go proto files
# Uses protoc compiler
make generate-proto
```

---

## Related Documentation

- [Architecture Overview](../../ARCHITECTURE.md#3-post-service-go)
- [Getting Started](../../GETTING_STARTED.md)
- [Proto Definitions](../../api/proto/post_service/)
- [Go Documentation](https://golang.org/doc/)
- [gRPC Guide](https://grpc.io/docs/)
- [uber-go/fx](https://github.com/uber-go/fx)

---

## Contributing

When adding new features:

1. Define domain entities and aggregates
2. Create commands/queries
3. Implement handlers
4. Add tests
5. Generate gRPC code if proto changes
6. Update documentation

---

## Support

For issues:

1. Check logs
2. Check infrastructure: `docker ps`
3. Open GitHub issue
