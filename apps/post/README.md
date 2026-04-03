# Post Service

Content management and post handling service using Domain-Driven Design and CQRS patterns.

## Overview

The Post Service is responsible for managing all post-related content, including historical posts, articles, and user-generated content. Built with Go using modern architectural patterns:

- **Domain-Driven Design (DDD)** - Clear business domain boundaries
- **CQRS Pattern** - Separate read and write models
- **Dependency Injection** - Using uber-go/fx framework
- **Event Sourcing** - Event-driven state management
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

Key variables:

- `DATABASE_URL` - PostgreSQL connection string
- `GRPC_PORT` - gRPC server port (default: 50051)
- `RABBITMQ_URL` - RabbitMQ connection string

### Running

```bash
# Development with hot reload (requires air)
make run

# Direct run
go run ./cmd/main.go

# Build
make build

# Production
./bin/nien-su-viet
```

Service will start on gRPC port `:50051` and HTTP port `:8080`

## Architecture

### Project Structure

```
apps/post/
├── bin/                      # Compiled binaries
├── cmd/
│   └── main.go              # Entry point
├── internal/
│   ├── app/                 # Application layer
│   ├── domain/              # Domain logic (DDD core)
│   │   ├── aggregates/      # Business aggregates
│   │   ├── entities/        # Domain entities
│   │   ├── value_objects/   # Value objects
│   │   ├── events/          # Domain events
│   │   └── repositories/    # Repository interfaces
│   ├── application/         # Service layer
│   │   ├── commands/        # CQRS commands
│   │   ├── queries/         # CQRS queries
│   │   └── dto/             # Data transfer objects
│   ├── infrastructure/      # Implementation details
│   │   ├── persistence/     # Database access
│   │   ├── adapters/        # External service adapters
│   │   └── events/          # Event publisher
│   ├── interfaces/          # API handlers
│   │   ├── grpc/            # gRPC handlers
│   │   ├── http/            # HTTP handlers
│   │   └── tcp/             # TCP handlers
│   └── shared/
│       └── grpc/genproto/   # Generated gRPC code
├── pkg/                     # Public packages (if any)
├── test/                    # Integration tests
├── go.mod                   # Module definition
├── go.sum                   # Dependency checksums
├── Makefile                 # Build automation
├── .env                     # Environment variables
└── nodemon.json            # Hot reload config
```

---

## Code Generation

### Generate gRPC Code

```bash
cd apps/post

# Generate Go proto files
# Uses protoc compiler
make generate-proto

# Or manually
protoc \
  --go_out=. \
  --go-grpc_out=. \
  --proto_path=../../api/proto \
  ../../api/proto/post_service/*.proto
```

Generated files go to:

```
internal/shared/grpc/genproto/
├── post_service/
│   ├── posts.pb.go
│   └── posts_grpc.pb.go
└── common/
    ├── pagination.pb.go
    └── response.pb.go
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

1. Check logs: `docker logs post-service`
2. Verify environment: `echo $DATABASE_URL`
3. Check infrastructure: `docker ps`
4. Open GitHub issue
