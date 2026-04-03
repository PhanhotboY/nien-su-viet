# Historical Event Service

Real-time historical event management and timeline service.

## Overview

The Historical Event Service manages all historical events and timelines for the Nien Su Viet platform. It provides:

- Event storage and retrieval
- Timeline management
- TCP-based real-time event updates
- Event categorization and tagging
- Full-text search across events

## Technology Stack

- **Framework:** NestJS 10+
- **Language:** TypeScript 5.0+
- **Database:** PostgreSQL with Prisma ORM
- **Real-time Communication:** TCP sockets
- **Runtime:** Node.js 18+ or Bun

## Quick Start

### Prerequisites

- Node.js 18+ or Bun installed
- PostgreSQL running and accessible
- Environment variables configured

### Installation

```bash
# From project root
bun install

# Install service-specific dependencies
cd apps/historical-event
```

### Configuration

Copy `.env.example` to `.env` and update values:

```bash
cp .env.example .env
```

Key variables:

- `DATABASE_URL` - PostgreSQL connection string (port 5434)
- `PORT` - HTTP server port (default: 3002)
- `TCP_PORT` - TCP server port for real-time events (default: 5000)

### Running

```bash
# Development mode with hot reload
bun dev historical-event

# Production build
bun run build
bun run start
```

Service will start on `http://localhost:3002` (HTTP) and `:5000` (TCP)

## Architecture

### Project Structure

```
apps/historical-event/
├── src/
│   ├── events/               # Events module
│   │   ├── controllers/      # HTTP endpoints
│   │   ├── services/         # Business logic
│   │   ├── entities/         # Database models
│   │   └── dto/              # Request/response DTOs
│   ├── timelines/            # Timeline management
│   │   ├── controllers/
│   │   ├── services/
│   │   └── entities/
│   ├── categories/           # Event categories
│   ├── tcp/                  # TCP server for real-time events
│   │   ├── tcp.gateway.ts    # TCP gateway
│   │   └── event.handler.ts  # Message handling
│   ├── config/               # Configuration
│   ├── app.module.ts         # Root module
│   └── main.ts               # Entry point
├── prisma/
│   ├── schema.prisma         # Database schema
│   └── migrations/           # Database migrations
├── test/                     # Tests
├── .env                      # Environment variables
├── package.json
└── tsconfig.json
```

---

## Development

### Hot Reload

Changes to `src/` automatically reload while service runs:

```bash
bun dev historical-event
```

### Debug Mode

Enable detailed logging:

```env
LOG_LEVEL=debug
```

## Related Documentation

- [Architecture Overview](../../ARCHITECTURE.md#4-historical-event-service-nestjstypescript)
- [Getting Started](../../GETTING_STARTED.md)
- [NestJS Documentation](https://docs.nestjs.com/)
- [Prisma Documentation](https://www.prisma.io/docs/)

---

## Support

For issues:

1. Check service logs: `docker logs historical-event-service`
2. Verify database: `psql $DATABASE_URL`
3. Test TCP: `nc -zv localhost 5000`
4. Open GitHub issue
