# Architecture Documentation

## System Architecture

This project is a **microservices-based monorepo** designed to provide an interactive platform for learning Vietnamese history. The system follows modern cloud-native patterns including DDD, CQRS, and event-driven architecture.

```
┌─────────────────────────────────────────────────────────────────┐
│                     Client Layer                                 │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  Next.js Frontend (React + TypeScript)                   │   │
│  │  - Responsive UI and historical content display          │   │
│  │  - User authentication & session management              │   │
│  └──────────────────────────────────────────────────────────┘   │
└────────────────┬────────────────────────────────────────────────┘
                 │ HTTP/HTTPS
                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                   Gateway API Layer                              │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  API Gateway (NestJS + Express)                          │   │
│  │  - Request routing & authentication proxy                │   │
│  │  - Rate limiting & CORS handling                         │   │
│  │  - Response serialization & compression                  │   │
│  │  - Input validation                                      │   │
│  └──────────────────────────────────────────────────────────┘   │
└────────────────┬────────────────────────────────────────────────┘
         ┌───────┴──────┬──────────┬─────────────┐
         │              │          │             │
         ▼              ▼          ▼             ▼
    ┌────────┐    ┌────────┐  ┌────────┐  ┌────────────┐
    │  Auth  │    │ Post   │  │History │  │  Other     │
    │Service │    │Service │  │Service │  │  Services  │
    │        │    │        │  │        │  │            │
    │NestJS  │    │  Go    │  │NestJS  │  │            │
    │(JWT)   │    │(DDD+  │  │(TCP)   │  │            │
    │        │    │CQRS)  │  │Events  │  │            │
    └────────┘    └────────┘  └────────┘  └────────────┘
         │              │          │
         └──────────────┼──────────────┘
                        │
         ┌──────────────┴──────────────┐
         │   Internal Communication   │
         │  (gRPC, REST, TCP, Events) │
         └──────────────┬──────────────┘
                        │
        ┌───────────────┴───────────────┐
        ▼               ▼               ▼
    ┌────────┐   ┌────────┐   ┌────────────┐
    │Database│   │ Cache  │   │  Message   │
    │       │   │ Layer  │   │   Queue    │
    │PosgSQL│   │ Redis  │   │ RabbitMQ   │
    │(3x)   │   │        │   │            │
    └────────┘   └────────┘   └────────────┘
```

## Service Descriptions

### 1. Gateway API (NestJS/TypeScript)

**Purpose:** Central entry point for all client requests

**Key Responsibilities:**

- Route requests to appropriate microservices
- Authenticate requests and proxy to Auth Service
- Apply rate limiting and compression
- Validate input data
- Serialize responses
- Handle CORS policies
- Graceful shutdown handling

**Technologies:** NestJS, Express, TypeScript

**Location:** [apps/gateway/](./apps/gateway)

---

### 2. Auth Service (NestJS/TypeScript)

**Purpose:** Handle authentication, authorization, and user management

**Key Features:**

- JWT token generation and validation
- User registration and login
- Role-based access control (RBAC)
- Permission management
- Session management

**Technologies:** NestJS, Prisma ORM, PostgreSQL

**Location:** [apps/auth/](./apps/auth)

**Database:** PostgreSQL (dedicated schema)

---

### 3. Post Service (Go)

**Purpose:** Manage historical posts and content

**Architectural Patterns:**

- **Domain-Driven Design (DDD):** Clear domain boundaries and business logic separation
- **CQRS (Command Query Responsibility Segregation):** Separate read and write operations
- **Dependency Injection:** Using uber-go/fx for IoC container
- **Event Sourcing:** Event-driven state management

**Key Features:**

- CRUD operations for posts
- Full-text search capabilities
- Event publishing for state changes
- gRPC communication with other services
- Comprehensive input validation
- Structured logging (Zap)
- Environment-based configuration (Viper)

**Technologies:** Go 1.21+, gRPC, Zap, Viper, Fx

**Location:** [apps/post/](./apps/post)

**Database:** PostgreSQL (dedicated schema)

**Proto Definitions:** [api/proto/post_service/](./api/proto/post_service)

---

### 4. Historical Event Service (NestJS/TypeScript)

**Purpose:** Manage historical events and timelines

**Key Features:**

- TCP event handling for real-time updates
- Event storage and retrieval
- Timeline management
- Event categorization

**Technologies:** NestJS, TCP sockets, Prisma ORM, PostgreSQL

**Location:** [apps/historical-event/](./apps/historical-event)

**Database:** PostgreSQL (dedicated schema)

---

### 5. Frontend (Next.js)

**Purpose:** User-facing web application

**Key Features:**

- Interactive historical content display
- User authentication interface
- Responsive design
- Real-time updates

**Technologies:** Next.js 14+, React, TypeScript

**Location:** [apps/client/](./apps/client)

---

## Data Flow

### Authentication Flow

1. Client submits credentials to Gateway API
2. Gateway proxies to Auth Service
3. Auth Service validates credentials and returns JWT token
4. Client includes JWT in subsequent requests
5. Gateway validates token and may refresh if needed

### Content Retrieval Flow

1. Client requests posts/events via Gateway API
2. Gateway routes to appropriate service (Post or Historical Event Service)
3. Service queries PostgreSQL and/or Redis cache
4. Response is serialized and returned to client

### Event Processing Flow

1. Service publishes event to RabbitMQ
2. Other services listen for specific events
3. Services update their views/caches accordingly
4. Eventual consistency achieved

---

## External Communication Protocols

### REST API

- **Used by:** Frontend ↔ Gateway API
- **Format:** JSON
- **Authentication:** JWT Bearer tokens

### gRPC

- **Used by:** Service-to-service communication
- **Format:** Protocol Buffers
- **Location:** [api/proto/](./api/proto)

### TCP

- **Used by:** Historical Event Service for real-time events
- **Format:** Binary protocol

---

## Data Persistence

### PostgreSQL Databases

- **3 separate PostgreSQL instances** (dev/staging/prod equivalent)
- **Schema 1:** Auth Service
- **Schema 2:** Post Service
- **Schema 3:** Historical Event Service

### Redis Cache

- **Purpose:** Session storage, response caching, rate limit counters
- **TTL:** Service-specific (configured per use case)

### RabbitMQ Message Queue

- **Purpose:** Async communication between services
- **Exchanges:** Topic-based for flexible routing
- **Dead Letter Queue:** For failed message handling

---

## Technology Stack

| Component          | Technology       | Version      |
| ------------------ | ---------------- | ------------ |
| Runtime            | Node.js / Go     | 20+ / 1.21+  |
| Backend Framework  | NestJS / Go      | 10+ / latest |
| Frontend Framework | Next.js          | 14+          |
| Language           | TypeScript / Go  | 5.0+ / 1.21+ |
| ORM                | Prisma           | 5.0+         |
| gRPC               | Protocol Buffers | 3.0+         |
| Database           | PostgreSQL       | 15+          |
| Cache              | Redis            | 7.0+         |
| Message Queue      | RabbitMQ         | 3.10+        |
| Logging            | Zap / Winston    | latest       |
| Containerization   | Docker           | 20.10+       |

---

## Deployment Architecture

### Development Environment

- Docker Compose for local services
- File watching and hot reload enabled
- Debug logging enabled

### Production Environment

- Containerized microservices
- Kubernetes orchestration (recommended)
- Centralized logging and monitoring
- Health checks and auto-recovery
- Distributed tracing (under development)

---

## Design Patterns Used

### 1. Microservices Pattern

- Independent deployment and scaling
- Service-specific technology choices
- Bounded contexts (DDD)

### 2. API Gateway Pattern

- Single entry point for clients
- Cross-cutting concerns (auth, rate limiting, compression)
- Service discovery and routing

### 3. Database per Service

- Independent data schemas
- Data consistency via events
- Easy to scale individual services

### 4. Event-Driven Architecture

- Loose coupling between services
- Asynchronous processing
- Event sourcing for audit trails

### 5. Domain-Driven Design (Post Service)

- Core domain logic separated from infrastructure
- Bounded contexts and aggregates
- Ubiquitous language

### 6. CQRS (Post Service)

- Separate models for read and write
- Optimized queries
- Event sourcing benefits

---

## Running the System

### Quick Start

```sh
# 1. Start infrastructure
./cli/dcp.sh dev start

# 2. Start services (each in separate terminal)
bun dev gateway
bun dev auth
bun dev historical-event
cd apps/post && make run
cd apps/client && bun dev
```

### Service Dependencies

```
Frontend (Client)
  ↓ HTTP/REST
Gateway API
  ↓ Proxy/Route
  ├→ Auth Service
  ├→ Post Service
  └→ Historical Event Service
  ↓ gRPC/TCP

Infrastructure
  ├→ PostgreSQL (3 instances)
  ├→ Redis
  └→ RabbitMQ
```

---

## Performance Considerations

1. **Response Caching:** Redis caching layer for frequently accessed data
2. **Request Compression:** GZIP compression enabled in Gateway
3. **Connection Pooling:** Database connection pools configured per service
4. **Rate Limiting:** Token-bucket algorithm in Gateway
5. **Async Processing:** Long-running tasks via RabbitMQ

---

## Security Architecture

1. **Authentication:** JWT tokens with RS256 signing
2. **Authorization:** Role-based access control (RBAC) at Gateway
3. **CORS:** Configured per environment
4. **Input Validation:** All inputs validated before processing
5. **HTTPS:** Enforced in production
6. **Environment Variables:** Sensitive configuration externalized

---

## Monitoring & Observability (Planned)

- **Distributed Tracing:** OpenTelemetry + Jaeger/Zipkin
- **Metrics:** Prometheus + Grafana
- **Logging:** Structured logs aggregated to central location
- **Health Checks:** Service health endpoints

---

## Future Architecture Improvements

1. Implement full OpenTelemetry integration
2. Add Kubernetes-native patterns
3. Implement service mesh (Istio)
4. Event sourcing for all services
5. CQRS pattern for all services
6. API versioning strategy
7. Contract testing between services

For more information, see individual service documentation in [apps/](./apps)
