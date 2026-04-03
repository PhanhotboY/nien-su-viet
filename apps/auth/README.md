# Auth Service

Authentication and Authorization service for Nien Su Viet platform.

## Overview

The Auth Service handles all authentication and authorization logic for the Nien Su Viet platform, including:

- User registration and login
- JWT token generation and validation
- Role-based access control (RBAC)
- Permission management
- Session management

## Technology Stack

- **Framework:** NestJS 10+
- **Language:** TypeScript 5.0+
- **Database:** PostgreSQL with Prisma ORM
- **Authentication:** JWT (RS256)
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

# Or from service directory
cd apps/auth
bun install
```

### Configuration

Copy `.env.example` to `.env` and update values:

```bash
cp .env.example .env
```

Key variables:

- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Secret for signing tokens (min 32 chars)
- `JWT_EXPIRATION` - Access token TTL in seconds
- `NODE_ENV` - Environment (development/production)

### Running

```bash
# Development mode with hot reload
bun dev auth

# Production build
bun run build
bun run start
```

Service will start on `http://localhost:3001`

## Architecture

### Project Structure

```
apps/auth/
├── src/
│   ├── auth/              # Auth module
│   │   ├── controllers/   # HTTP endpoints
│   │   ├── services/      # Business logic
│   │   ├── strategies/    # Auth strategies (JWT, etc)
│   │   └── guards/        # Auth guards
│   ├── users/             # Users module
│   │   ├── model/         # Entity definitions
│   │   └── services/      # User business logic
│   ├── roles/             # Role management
│   ├── permissions/       # Permission management
│   ├── config/            # Configuration
│   ├── app.module.ts      # Root module
│   └── main.ts            # Entry point
├── prisma/
│   ├── schema.prisma      # Database schema
│   └── migrations/        # Database migrations
├── test/                  # Tests
├── .env                   # Environment variables
├── package.json
└── tsconfig.json
```

### Token Structure

**Access Token (JWT):**

- Short-lived (1 hour default)
- Used for API requests
- Contains user ID, roles, permissions
- Signed with RS256

**Refresh Token:**

- Long-lived (7 days default)
- Used to obtain new access tokens
- Stored securely server-side
- Rotated on use

---

## Related Documentation

- [Architecture Overview](../../ARCHITECTURE.md#architecture-documentation)
- [Getting Started](../../GETTING_STARTED.md)
- [Gateway API](../gateway/README.md)
- [Prisma Documentation](https://www.prisma.io/docs/)
- [NestJS Documentation](https://docs.nestjs.com/)

---

## Support

For issues or questions:

1. Review service logs: `docker logs auth-service`
2. Open an issue on GitHub
3. Check existing issues and discussions
