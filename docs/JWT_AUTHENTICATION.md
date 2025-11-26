# JWT Authentication Setup

This guide explains how services verify requests using JWT tokens without calling back to the auth service.

## Overview

The system uses **RS256 (RSA with SHA-256)** for JWT signing:

- **Auth Service**: Signs JWTs with a **private key**
- **Other Services**: Verify JWTs with a **public key**

This eliminates the need to call the auth service for every request verification.

## How It Works

1. **User logs in** → Auth service creates session → Issues JWT token
2. **Client stores JWT** → In cookie or Authorization header
3. **Client makes request** → Sends JWT to any service
4. **Service verifies JWT locally** → Using public key (no auth service call needed)
5. **Request proceeds** → With user data extracted from JWT

## Setup

### 1. Generate RSA Key Pair

Generate a new RSA key pair for JWT signing:

```bash
# Generate private key
openssl genrsa -out private.pem 2048

# Extract public key
openssl rsa -in private.pem -pubout -out public.pem

# For use in .env (convert to single line)
# Private key:
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' private.pem

# Public key:
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' public.pem
```

### 2. Environment Variables

Add to your `.env` files:

**Auth Service** (needs both keys):

```bash
JWT_PRIVATE_KEY="-----BEGIN RSA PRIVATE KEY-----\nMIIE...your key here...\n-----END RSA PRIVATE KEY-----"
JWT_PUBLIC_KEY="-----BEGIN PUBLIC KEY-----\nMIIB...your key here...\n-----END PUBLIC KEY-----"
JWT_EXPIRES_IN=1h
```

**Other Services** (only need public key):

```bash
JWT_PUBLIC_KEY="-----BEGIN PUBLIC KEY-----\nMIIB...your key here...\n-----END PUBLIC KEY-----"
```

### 3. How JWT is Issued

The auth service automatically issues JWT tokens when using Better Auth's JWT plugin:

```typescript
// In auth.ts - already configured
jwt({
  jwt: {
    sign: async (payload) => {
      const jwtLib = await import('jsonwebtoken');
      return jwtLib.sign(payload, config.jwt.privateKey, {
        algorithm: 'RS256',
        expiresIn: config.jwt.signOptions.expiresIn,
      });
    },
  },
});
```

### 4. How JWT is Verified

The `BetterAuthGuard` in other services verifies tokens locally:

```typescript
// Automatically extracts JWT from:
// 1. Authorization: Bearer <token>
// 2. Cookie: better-auth.session_token

const decoded = jwt.verify(token, publicKey, {
  algorithms: ['RS256'],
});
```

## Client Integration

### Using Authorization Header

**Frontend (Next.js/React):**

```typescript
// After login, Better Auth provides the JWT
const token = await authClient.getSession();

// Send with requests
fetch('http://localhost:3001/api/users/me', {
  headers: {
    Authorization: `Bearer ${token}`,
  },
});
```

### Using Cookies (Recommended)

Better Auth automatically sets JWT in cookies. No additional code needed:

```typescript
// Just make requests - cookies are sent automatically
fetch('http://localhost:3001/api/users/me', {
  credentials: 'include', // Important: include cookies
});
```

## Usage in Controllers

### Protected Route (Default)

```typescript
@Controller('users')
export class UserController {
  @Get('me')
  getProfile(@CurrentUser() user: any) {
    // user is automatically extracted from JWT
    return {
      id: user.id,
      email: user.email,
      role: user.role,
    };
  }
}
```

### Public Route

```typescript
import { Public } from '@phanhotboy/nsv-common';

@Controller('health')
export class HealthController {
  @Public()
  @Get()
  check() {
    return { status: 'ok' };
  }
}
```

### Role-Based Access

```typescript
import { Roles, CurrentUser } from '@phanhotboy/nsv-common';

@Controller('admin')
export class AdminController {
  @Roles('admin')
  @Get('users')
  getAllUsers() {
    // Only users with role 'admin' can access
    return [];
  }

  @Roles('admin', 'moderator')
  @Delete('posts/:id')
  deletePost(@Param('id') id: string) {
    // Admins or moderators can access
  }
}
```

## JWT Payload Structure

The JWT token contains:

```json
{
  "sub": "user-id-here", // User ID (standard JWT claim)
  "userId": "user-id-here", // Better Auth user ID
  "email": "user@example.com",
  "role": "admin",
  "sessionId": "session-id",
  "iat": 1234567890, // Issued at timestamp
  "exp": 1234571490 // Expiration timestamp
}
```

Access in your code:

```typescript
@Get('profile')
getProfile(@CurrentUser() user: any) {
  console.log(user.id);        // User ID
  console.log(user.email);     // User email
  console.log(user.role);      // User role
}

// Or access specific field
@Post('posts')
createPost(@CurrentUser('id') userId: string) {
  return { authorId: userId };
}
```

## Testing

### 1. Get JWT Token

**Login via Auth Service:**

```bash
curl -c cookies.txt -X POST http://localhost:3000/api/v1/auth/sign-in/email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123"
  }'
```

**Extract token from response or cookie.**

### 2. Test with Authorization Header

```bash
TOKEN="your-jwt-token-here"

curl http://localhost:3001/api/users/me \
  -H "Authorization: Bearer $TOKEN"
```

### 3. Test with Cookie

```bash
curl -b cookies.txt http://localhost:3001/api/users/me
```

### 4. Test Public Route

```bash
curl http://localhost:3001/api/health
# Should work without token
```

### 5. Test Protected Route Without Token

```bash
curl http://localhost:3001/api/users/me
# Should return 401 Unauthorized
```

## Security Considerations

### ✅ Advantages

1. **No Network Calls**: Services verify tokens locally (faster)
2. **Scalable**: No bottleneck on auth service
3. **Resilient**: Other services work even if auth service is down
4. **Stateless**: No session storage needed in other services

### ⚠️ Important Notes

1. **Private Key Security**:
   - Never commit private key to git
   - Store in secure environment variables
   - Rotate keys periodically

2. **Token Expiration**:
   - Keep expiration time reasonable (1h recommended)
   - Implement token refresh mechanism

3. **Token Revocation**:
   - JWT tokens cannot be revoked before expiration
   - For immediate revocation, maintain a token blacklist in Redis

4. **HTTPS Required**:
   - Always use HTTPS in production
   - Tokens in headers/cookies must be encrypted in transit

## Troubleshooting

### "Invalid token" Error

**Causes:**

- Wrong public key in service config
- Token expired
- Token malformed

**Solutions:**

```bash
# Verify public key matches
echo $JWT_PUBLIC_KEY

# Check token expiration
# Decode JWT at https://jwt.io

# Ensure keys are properly formatted with \n for line breaks
```

### "No authentication token found"

**Causes:**

- Token not sent in request
- Wrong cookie name
- CORS blocking cookies

**Solutions:**

```typescript
// Ensure cookies are sent
fetch(url, { credentials: 'include' });

// Or use Authorization header
fetch(url, {
  headers: { Authorization: `Bearer ${token}` },
});
```

### User Object is Undefined

**Causes:**

- JWT payload doesn't match expected structure
- Token verification failed silently

**Solutions:**

```typescript
// Check what's in the JWT
@Get('debug')
debug(@CurrentUser() user: any) {
  console.log('User object:', user);
  return user;
}
```

### Keys Mismatch Between Services

**Problem**: Different services using different public keys

**Solution**: Use environment variables or config service:

```bash
# Share same public key across all services
# Use a config management system in production
# Or mount as Kubernetes secret
```

## Token Refresh Strategy

When tokens expire, implement a refresh mechanism:

```typescript
// In your auth service
@Post('refresh')
async refreshToken(@Req() req: Request) {
  // Verify refresh token (stored in DB)
  // Issue new access token
  const newToken = await auth.api.refreshJwtToken({
    refreshToken: req.cookies['refresh_token']
  });

  return { token: newToken };
}
```

**Client-side:**

```typescript
// Intercept 401 errors and refresh
axios.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      // Refresh token
      const { token } = await refreshToken();
      // Retry original request
      error.config.headers.Authorization = `Bearer ${token}`;
      return axios(error.config);
    }
    return Promise.reject(error);
  },
);
```

## Production Deployment

### Environment Setup

```bash
# Generate strong keys
openssl genrsa -out private.pem 4096
openssl rsa -in private.pem -pubout -out public.pem

# Add to your deployment secrets
# - Kubernetes: as Secret
# - Docker: as environment variables
# - Cloud providers: use secrets manager
```

### Key Rotation

Implement key rotation for enhanced security:

1. Generate new key pair
2. Update auth service with new private key
3. Keep old public key for verification (grace period)
4. Update all services with new public key
5. Remove old public key after grace period

### Monitoring

Monitor JWT-related metrics:

- Token verification failures
- Expired token attempts
- Invalid signature errors
- Token usage by service

```typescript
// Add logging to guard
if (error instanceof jwt.TokenExpiredError) {
  logger.warn('Expired token attempt', { userId: decoded.sub });
}
```

## Complete Example

**Controller:**

```typescript
import { Controller, Get, Post, Body } from '@nestjs/common';
import { CurrentUser, Public, Roles } from '@phanhotboy/nsv-common';

@Controller('posts')
export class PostsController {
  // Public - no auth needed
  @Public()
  @Get()
  getAllPosts() {
    return { posts: [] };
  }

  // Protected - requires valid JWT
  @Get('my-posts')
  getMyPosts(@CurrentUser() user: any) {
    return {
      userId: user.id,
      posts: [],
    };
  }

  // Admin only
  @Roles('admin')
  @Delete(':id')
  deletePost(@Param('id') id: string) {
    return { deleted: id };
  }

  // Create post with current user
  @Post()
  createPost(@Body() data: any, @CurrentUser('id') userId: string) {
    return {
      ...data,
      authorId: userId,
    };
  }
}
```

## Summary

- ✅ JWT tokens are signed by auth service with private key
- ✅ Other services verify tokens locally with public key
- ✅ No need to call auth service for verification
- ✅ Fast, scalable, and resilient architecture
- ✅ Use `@Public()` for routes that don't need auth
- ✅ Use `@Roles()` for role-based access control
- ✅ Use `@CurrentUser()` to get user info from JWT
