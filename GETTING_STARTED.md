# Getting Started Guide

Complete setup guide for development environment and running the Nien Su Viet project.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Initial Setup](#initial-setup)
- [Running Services](#running-services)

---

## Prerequisites

Ensure you have the following installed on your system:

### Required Tools

| Tool                        | Version       | Purpose                           | Link                                                            |
| --------------------------- | ------------- | --------------------------------- | --------------------------------------------------------------- |
| **Node.js or Bun**          | 18+ or latest | TypeScript/JavaScript runtime     | [bun.sh](https://bun.sh) or [nodejs.org](https://nodejs.org)    |
| **Go**                      | 1.21+         | Post Service runtime              | [golang.org](https://golang.org)                                |
| **Docker & Docker Compose** | 20.10+        | Container & service orchestration | [docker.com](https://docker.com)                                |
| **GNU Make**                | 4.0+          | Build automation                  | Pre-installed on macOS/Linux                                    |
| **Protoc**                  | 3.0+          | Protocol Buffers compiler         | [installation guide](https://grpc.io/docs/protoc-installation/) |
| **Git**                     | 2.30+         | Version control                   | Pre-installed on most systems                                   |

---

## Initial Setup

### 1. Clone the Repository

```bash
git clone https://github.com/phanhotboy/nien-su-viet.git
cd nien-su-viet
```

### 2. Make CLI Scripts Executable

```bash
chmod +x ./cli/*.sh
```

### Service Credentials (Development)

| Service            | Host                   | Port | Username | Password |
| ------------------ | ---------------------- | ---- | -------- | -------- |
| PostgreSQL (Auth)  | localhost              | 5432 | postgres | password |
| PostgreSQL (Post)  | localhost              | 5433 | postgres | password |
| PostgreSQL (Event) | localhost              | 5434 | postgres | password |
| Redis              | localhost              | 6379 | -        | -        |
| RabbitMQ           | localhost              | 5672 | guest    | guest    |
| RabbitMQ UI        | http://localhost:15672 | -    | guest    | guest    |

### Stop Infrastructure

```bash
./cli/dcp.sh dev stop
```

---

## Running Services

### Setup: Install Dependencies

Before running any service, install shared dependencies:

```bash
bun install
```

**Terminal 1: Start Infrastructure**

```bash
# Start all infrastructure services
./cli/dcp.sh dev start

# This will:
# - Pull required Docker images
# - Start 3 PostgreSQL containers
# - Start 1 Redis container
# - Start 1 RabbitMQ container
# - Create necessary networks and volumes
```

**Terminal 2: Gateway API**

```bash
bun dev gateway

# Gateway will be available at http://localhost:3000
```

**Terminal 3: Auth Service**

```bash
bun dev auth

# Auth Service runs on http://localhost:3001
```

**Terminal 4: Historical Event Service**

```bash
bun dev historical-event

# Event Service runs on http://localhost:3002
```

**Terminal 5: Post Service**

```bash
cd apps/post
go mod tidy
make run

# Post Service runs on configured port (check main.go)
```

**Terminal 6: Frontend Client**

```bash
cd apps/client
bun dev

# Frontend available at http://localhost:3000
# (Note: You may need to adjust port if Gateway is on 3000)
```

### Service Ports (Default Development)

| Service          | Port   | URL                    |
| ---------------- | ------ | ---------------------- |
| Gateway API      | 3000   | http://localhost:3000  |
| Auth Service     | 3001   | http://localhost:3001  |
| Historical Event | 3002   | http://localhost:3002  |
| Frontend         | 3000\* | http://localhost:3000  |
| RabbitMQ UI      | 15672  | http://localhost:15672 |

\*Frontend will use a different port if 3000 is occupied

---

## Code Generation

### Generate Protocol Buffer Files

Protocol Buffer files need to be generated before running services.

**For Go (Post Service):**

```bash
./cli/gengoproto.sh
```

**For TypeScript:**

```bash
./cli/gentsproto.sh
```

Both scripts:

- Parse `.proto` files in `api/proto/`
- Generate language-specific code
- Place generated files in appropriate service directories

**Proto files are located in:**

- `api/proto/common/` - Shared definitions
- `api/proto/post_service/` - Post service gRPC
- `api/proto/search_service/` - Search service gRPC

---

## Environment Configuration

Create `.env` files by copying from `.env.example` and editing values:

```bash
cd apps/auth && cp .env.example .env && nano .env
cd ../historical-event && cp .env.example .env && nano .env
cd ../post && cp .env.example .env && nano .env
cd ../gateway && cp .env.example .env && nano .env
```

See individual service READMEs for environment variable descriptions.

---

## Next Steps

1. ✅ Infrastructure running
2. ✅ Services started
3. 📖 Read [ARCHITECTURE.md](./ARCHITECTURE.md) for system design
4. 📖 Read individual service READMEs in [apps/](./apps/)
5. 🎨 Start developing!

---

## Additional Resources

- [NestJS Documentation](https://docs.nestjs.com/)
- [Go Language Docs](https://golang.org/doc/)
- [Prisma Documentation](https://www.prisma.io/docs/)
- [gRPC Documentation](https://grpc.io/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [Protocol Buffers Guide](https://developers.google.com/protocol-buffers)

For specific service documentation, see [Architecture.md](./ARCHITECTURE.md#service-descriptions).
