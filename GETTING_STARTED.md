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

### Generate Prisma

**For Auth Service:**

```bash
make generate-auth
```

**For Histroical Event Service:**

```bash
make generate-historical-event
```

---

## Environment Configuration

Create `.env` files by copying from `.env.example`, keep as is if want to use values from [infrastructure services](#terminal-1-start-infrastructure).

For more detailed, see individual service READMEs.

---

## Running Services

### Setup: Install Dependencies

Before running any service, install shared dependencies:

```bash
bun install
```

#### Terminal 1: Start Infrastructure

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

_Infrastructure Credentials (Development)_

| Service            | Port  | Username      | Password        |
| ------------------ | ----- | ------------- | --------------- |
| PostgreSQL (Auth)  | 5432  | nsv_auth_user | secure_password |
| PostgreSQL (Event) | 5433  | nsv_auth_user | secure_password |
| PostgreSQL (Post)  | 5434  | nsv_auth_user | secure_password |
| Redis              | 6379  | -             | -               |
| RabbitMQ           | 5672  | nsv_rmq_user  | secure_password |
| RabbitMQ UI        | 15672 | nsv_rmq_user  | secure_password |

#### Terminal 2: Gateway API

```bash
bun dev gateway
```

#### Terminal 3: Auth Service

```bash
bun dev auth
```

#### Terminal 4: Historical Event Service

```bash
bun dev historical-event
```

#### Terminal 5: Post Service

```bash
cd apps/post
go mod tidy
make run
```

#### Terminal 6: Frontend Client

```bash
cd apps/client
bun dev
```

### Service Ports (Default Development)

| Service          | Port  | Method |
| ---------------- | ----- | ------ |
| Gateway API      | 8080  | HTTP   |
| Auth Service     | 8081  | HTTP   |
| Historical Event | 8082  | TCP    |
| Post Service     | 8082  | gRPC   |
| Frontend         | 3000  | HTTP   |
| RabbitMQ UI      | 15672 | HTTP   |

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
