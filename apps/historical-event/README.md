# Historical Event Service

Real-time historical event management and timeline service.

## Overview

The Historical Event Service manages all historical events and timelines for the Nien Su Viet platform. It provides:

- Event storage and retrieval
- Timeline management
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
```

### Configuration

Copy `.env.example` to `.env` and update values:

```bash
cp .env.example .env
```

Key variables:

- `DATABASE_URL` - PostgreSQL connection string (port 5433)
- `NODE_PORT` - TCP server port (default: 8082)

### Running

```bash
# Development mode with hot reload
bun dev historical-event

# Production build
bun run build
bun run start
```

## Architecture

### Project Structure

```
apps/historical-event/
├── generated
│   └── prisma          # Generated Prisma assets
├── prisma
│   ├── migrations      # Database migrations
│   └── schema.prisma   # Database schema
├── setup               # Set up scripts
├── src
│   ├── config          # Load config from .env
│   ├── database        # Prisma client module
│   ├── modules         # App logic modules
│   ├── app.module.ts   # Root module
│   └── main.ts         # Entry point
├── README.md
├── prisma.config.ts    # Prisma config
└── tsconfig.app.json   # Extended TS config
```

---

## Development

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

1. Check service logs
2. Verify database: `psql $DATABASE_URL`
3. Test TCP: `nc -zv localhost 5000`
4. Open GitHub issue
