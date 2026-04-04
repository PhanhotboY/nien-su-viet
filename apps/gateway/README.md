# Gateway API

Central API gateway for routing, authentication, and cross-cutting concerns.

## Overview

The Gateway API serves as the single entry point for all client requests. It handles:

- Request routing to microservices
- Authentication proxy to Auth Service
- Rate limiting and throttling
- Response compression
- Input validation
- CORS policy enforcement
- Request/response serialization
- Graceful shutdown

## Technology Stack

- **Framework:** NestJS 10+
- **Language:** TypeScript 5.0+
- **Runtime:** Node.js 18+ or Bun

## Quick Start

### Prerequisites

- Node.js 18+ or Bun installed
- All other services running or configured
- Environment variables set

### Installation

```bash
# From project root
bun install
```

### Configuration

Copy `.env.example` to `.env` and update values:

```bash
cp .env.example .env
```

Key variables:

- `PORT` - Gateway server port (default: 3000)
- `AUTH_SERVICE_URL` - Auth service URL
- `POST_SERVICE_URL` - Post service gRPC URL
- `EVENT_SERVICE_URL` - Event service URL

### Running

```bash
# Development mode with hot reload
bun dev gateway

# Production build
bun run build
bun run start
```

---

## Architecture

### Request Flow

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │ HTTP Request
       ▼
┌──────────────────────┐
│  Gateway API         │
├──────────────────────┤
│ 1. CORS Handler      │
│ 2. Auth Guard        │  ← Validates JWT token
│ 3. Validation Pipe   │  ← Validates input
│ 4. Rate Limiting     │  ← Throttle requests
│ 5. Router            │  ← Route to service
│ 6. Compression       │  ← Compress response
└────────────┬─────────┘
             │
     ┌───────┴───────┬──────────┐
     │               │          │
     ▼               ▼          ▼
 Auth Service   Post Service   Event Service
```

### Project Structure

```
apps/gateway/
├── src
│   ├── common                     # Shared decorators, filters, guards, interceptors,...
│   ├── config                     # Load config from .env
│   ├── gateway.middleware.ts      # Inject midlewares
│   ├── gateway.module.ts          # Root module
│   ├── modules                    # Service modules
│   ├── utils
│   └── main.ts                    # Entry point
├── README.md
└── tsconfig.app.json              # Extended TS config
```

---

## Related Documentation

- [Architecture Overview](../../ARCHITECTURE.md#1-gateway-api-nestjstypescript)
- [Getting Started](../../GETTING_STARTED.md)
- [Auth Service](../auth/README.md)
- [Post Service](../post/README.md)
- [Event Service](../historical-event/README.md)
- [NestJS Documentation](https://docs.nestjs.com/)

---

## Support

For issues:

1. Check logs
2. Review configuration
3. Open GitHub issue
