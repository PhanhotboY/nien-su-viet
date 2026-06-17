# Billing Service

Subscription & Payment management service using Domain-Driven Design and CQRS patterns.

## Overview

The Billing Service is responsible for managing all aspects of subscription plans, billing cycles, and payment processing. Built with Go using modern architectural patterns:

- **Domain-Driven Design (DDD)** - Clear business domain boundaries
- **Dependency Injection** - Using uber-go/fx framework
- **gRPC** - Internal service communication

## Technology Stack

- **Language:** Go 1.25+
- **Framework:** Custom Go (no heavy framework)
- **Database:** PostgreSQL
- **Communication:** gRPC, RabbitMQ
- **Logging:** Zap (structured logging)
- **Configuration:** Viper
- **Dependency Injection:** uber-go/fx
- **Validation:** go-playground/validator

## Quick Start

### Prerequisites

- Go 1.25+ installed
- PostgreSQL running
- Environment variables configured

### Installation

```bash
cd apps/billing

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

Service will start on gRPC port `:8085`

## Architecture

### Project Structure

```
apps/billing/
├── bin/                      # Compiled binaries
├── cmd
│   └── server
│       └── main.go           # Entry point
├── internal
│   ├── subscriptions         # Subscription module
│   │   ├── application       # Application layer
│   │   ├── domain            # Domain layer
│   │   ├── infrastructure    # Infrastructure layer
│   │   |   ├── cache         # Outbound Redis repo
│   │   |   ├── persistence   # Outbound PostgreSQL repo
│   │   |   ├── messaging     # Inbound/Outbound RabbitMQ consumer/producer
│   │   |   └── transport     # Inbound gRPC
│   │   └── subscriptions.module.go
│   ├── outbox_events         # Outbox event module
│   ├── payment_attempts      # Payment attempt module
│   ├── payment_transactions  # Payment transaction module
│   ├── plans                 # Plan module
│   ├── processed_events      # Processed event module
│   ├── purchases             # Purchase module
│   ├── subscription_events   # Subscription event module
│   └── shared
│       ├── app               # Wire dependencies
│       ├── config            # Service configuration
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
# Generate Go proto files. Uses protoc compiler
make generate-proto
```

---

## Checklist

### Purchase module

  - [x] Define purchase domain entity
  - [x] Define purchase DTOs
  - [x] Create purchase command and handler
  - [x] Create purchase query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers

### Payment Attempt module

  - [x] Define payment attempt domain entity
  - [x] Create payment attempt command and handler
  - [x] Create payment attempt query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers
  
### Payment Transaction module

  - [x] Define payment transaction domain entity
  - [x] Define payment transaction DTOs
  - [x] Create payment transaction command and handler
  - [x] Create payment transaction query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers

### Inbox Event module

  - [x] Define inbox event domain entity
  - [x] Define inbox event DTOs
  - [x] Create inbox event command and handler
  - [x] Create inbox event query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers
  
### Outbox Event module

  - [x] Define outbox event domain entity
  - [x] Define outbox event DTOs
  - [x] Create outbox event command and handler
  - [x] Create outbox event query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers
  
### Processed Event module

  - [x] Define processed event domain entity
  - [x] Define processed event DTOs
  - [x] Create processed event command and handler
  - [x] Create processed event query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers
  
### Plan module

  - [x] Define plan domain entity
  - [x] Define plan DTOs
  - [x] Create plan command and handler
  - [x] Create plan query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers
  
### Subscription module

  - [x] Define subscription domain entity
  - [x] Define subscription DTOs
  - [x] Create subscription command and handler
  - [x] Create subscription query and handler
  - [ ] Consume/produce events
  - [x] Wire gRPC handlers

  ### Payment webhook module
  - [x] Handle payment callback

  ### Worker
  - [ ] Implement worker to process outbox events
  - [x] Implement worker to process inbox events

  ### Testing
  - [ ] Integration test purchase (in progress)
  - [ ] Unit test purchase
  - [ ] Integration test payment attempt
  - [ ] Unit test payment attempt
  - [ ] Integration test payment transaction
  - [ ] Unit test payment transaction
  - [ ] Integration test outbox event
  - [ ] Unit test outbox event
  - [ ] Integration test processed event
  - [ ] Unit test processed event
  - [ ] Integration test plan
  - [ ] Unit test plan
  - [ ] Integration test subscription
  - [ ] Unit test subscription
  - [ ] Test payment webhook

---

## Related Documentation

- [Architecture Overview](../../ARCHITECTURE.md#3-billing-service-go)
- [Getting Started](../../GETTING_STARTED.md)
- [Proto Definitions](../../api/proto/billing_service/)
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
