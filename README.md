<h1 align="center">Nien Su Viet</h1>

<p align="center">Provide a more interesting way to learn the heroic history of Vietnam.</p>

<p align="center">
  <a href="#quick-start">Quick Start</a> •
  <a href="#documentation">Documentation</a> •
  <a href="#architecture">Architecture</a> •
  <a href="#features">Features</a> •
  <a href="#contribution">Contribution</a>
</p>

## Quick Start

**Prerequisites:** Bun (or Node.js), Go, Make, Docker, Protoc

```sh
# 1. Clone and setup repository
git clone <repository-url>
cd nien-su-viet

# 2. Give execute permissions to CLI scripts
chmod +x ./cli/*.sh

# 3. Start infrastructure services
./cli/dcp.sh dev start

# 4. Start all services (in separate terminals)
bun i
bun dev gateway
bun dev auth
bun dev historical-event

cd apps/post && go mod tidy && make run

# 5. Start the frontend
cd apps/client && bun dev
```

👉 **[Read the complete setup guide →](./GETTING_STARTED.md)**

## Documentation

- **[Getting Started](./GETTING_STARTED.md)** - Detailed setup, troubleshooting, prerequisites
- **[Architecture](./ARCHITECTURE.md)** - System design, service relationships, technology stack
- **[Auth Service](./apps/auth/README.md)** - Authentication & authorization service
- **[Post Service](./apps/post/README.md)** - Posts & content management service
- **[Historical Event Service](./apps/historical-event/README.md)** - Historical events service
- **[Gateway API](./apps/gateway/README.md)** - API Gateway & routing service
- **[Frontend](./apps/client/README.md)** - Next.js client application

## Architecture

This is a **microservices-based monorepo** with:

- **Gateway API** (NestJS/TypeScript) - Entry point, auth proxy, rate limiting, response serialization
- **Auth Service** (NestJS/TypeScript) - JWT/OAuth2, user registration, RBAC
- **Post Service** (Go) - DDD + CQRS pattern, content management
- **Historical Event Service** (NestJS/TypeScript) - Event management
- **Frontend** (Next.js) - React-based client application

👉 **[View detailed architecture diagram →](./ARCHITECTURE.md#system-architecture)**

## Features

### Gateway API

- [x] Request authorization & authentication
- [x] Role-based access control (RBAC)
- [x] Rate limiting
- [x] Response compression & serialization
- [x] CORS support
- [x] Response caching
- [x] Input validation
- [x] Graceful shutdown
- [x] Auth service request proxying

### Auth Service

- [x] JWT token management
- [x] User authentication & registration
- [x] Role-based authorization

### Post Service (Go)

- [x] Domain-driven design (DDD)
- [x] CQRS pattern
- [x] Dependency injection (uber-go/fx)
- [x] gRPC internal communication
- [x] Input validation
- [x] Structured logging (Zap)
- [x] Configuration management (Viper)
- [ ] Distributed tracing (OpenTelemetry + Jaeger/Zipkin)
- [ ] Metrics collection (Prometheus + Grafana)
- [ ] Unit tests with mocking
- [ ] Integration & E2E tests

### Historical Event Service

- [x] TCP event handling
- [x] Event storage & retrieval

### Frontend

- [x] Next.js 14+ with modern React patterns
- [x] Type-safe development with TypeScript
- [x] Responsive UI components

## Development

### Infrastructure Services

Required services (started via `./cli/dcp.sh dev start`):

- 3 PostgreSQL instances
- 1 Redis instance
- 1 RabbitMQ instance

### Running Services

```sh
# Terminal 1:
# Start Infrastructure
./cli/dcp.sh dev start
# Gateway API
bun i && bun dev gateway

# Terminal 3: Auth Service
bun dev auth

# Terminal 4: Historical Event Service
bun dev historical-event

# Terminal 5: Post Service
cd apps/post && go mod tidy && make run

# Terminal 6: Frontend
cd apps/client && bun dev
```

### Code Generation

Generate protobuf files:

```sh
./cli/gengoproto.sh   # Go proto files
./cli/gentsproto.sh   # TypeScript proto files
```

See [proto documentation](./api/proto/) for schema definitions.

## Tech Stack

| Layer             | Technology                    | Services                              |
| ----------------- | ----------------------------- | ------------------------------------- |
| **Frontend**      | Next.js 14, React, TypeScript | Client App                            |
| **Backend**       | NestJS, Go, Express           | Gateway, Auth, Post, Historical Event |
| **Communication** | gRPC, REST, TCP               | Internal & External                   |
| **Database**      | PostgreSQL                    | Primary data store                    |
| **Cache**         | Redis                         | Session, cache layer                  |
| **Message Queue** | RabbitMQ                      | Async messaging                       |
| **Logging**       | Zap, Winston                  | Structured logging                    |

## Contribution

The application is in **active development**. Contributions are welcome! Please:

1. Create an issue for feature discussions
2. Submit a pull request with clear descriptions
3. Follow code style conventions

## License

The project is under [MIT license](./LICENSE).
