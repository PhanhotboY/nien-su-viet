# Getting JWT Token from Better Auth

Better Auth manages sessions with its own tokens. For microservices that need JWT verification, we've added a custom endpoint to generate JWT tokens.

## Getting the JWT Token

### Method 1: Call the `/auth/token` Endpoint

After logging in with Better Auth, call the token endpoint to get a JWT:

```typescript
// Client-side (Next.js/React)
import { authClient } from '@/lib/auth-client';

// 1. First, sign in
await authClient.signIn.email({
  email: 'user@example.com',
  password: 'password123',
});

// 2. Get JWT token
const response = await fetch('http://localhost:3000/api/v1/auth/token', {
  credentials: 'include', // Important: sends session cookie
});

const { token, expiresIn, type } = await response.json();
// token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."
// expiresIn: "24h"
// type: "Bearer"

// 3. Store token for use in other services
localStorage.setItem('jwt_token', token);
```

### Method 2: Use cURL

````bash
# 1. Login first
curl -c cookies.txt -X POST http://localhost:3000/api/v1/auth/sign-in/email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123"
  }'

# 2. Get JWT token using session cookie
curl -b cookies.txt http://localhost:3000/api/v1/auth/token

# Response:
# {
#   "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
#   "expiresIn": "24h",
#   "type": "Bearer"
# }
```## Using the JWT Token

### Method 1: Cookie (Recommended - Automatic)

Better Auth automatically sets the JWT in a cookie. No extra code needed:

```typescript
// Frontend - cookies are sent automatically
fetch('http://localhost:3001/api/users/me', {
  credentials: 'include', // Important!
});

// With axios
axios.get('http://localhost:3001/api/users/me', {
  withCredentials: true,
});

// With fetch in Next.js Server Actions
const response = await fetch('http://localhost:3001/api/users/me', {
  headers: await headers(), // Forwards cookies from client
});
````

### Method 2: Authorization Header (Manual)

If you want to manually handle tokens:

```typescript
// 1. Get the token from cookie
function getTokenFromCookie(name: string): string | null {
  const cookies = document.cookie.split(';');
  for (const cookie of cookies) {
    const [key, value] = cookie.trim().split('=');
    if (key === name) {
      return value;
    }
  }
  return null;
}

const token = getTokenFromCookie('better-auth.session_token');

// 2. Send in Authorization header
fetch('http://localhost:3001/api/users/me', {
  headers: {
    Authorization: `Bearer ${token}`,
  },
});
```

## Client Configuration (Next.js Example)

### Update your auth client to enable JWT

```typescript
// apps/client/src/lib/auth-client.ts
import { createAuthClient } from 'better-auth/react';

export const authClient = createAuthClient({
  baseURL: process.env.NEXT_PUBLIC_AUTH_URL || 'http://localhost:3000',
  plugins: [
    // Enable JWT in client
  ],
});
```

### Service Helper with Auto Token Handling

```typescript
// apps/client/src/lib/api-client.ts
import { headers } from 'next/headers';

export async function apiClient(url: string, options: RequestInit = {}) {
  const reqHeaders = await headers();

  return fetch(url, {
    ...options,
    headers: {
      ...options.headers,
      Cookie: reqHeaders.get('cookie') || '',
    },
    credentials: 'include',
  });
}

// Usage in server actions
export async function getUsers() {
  'use server';

  const response = await apiClient('http://localhost:3001/api/users');
  return response.json();
}
```

### Client Component Example

```typescript
'use client';

import { useEffect, useState } from 'react';

export function UserProfile() {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Cookies are automatically sent with fetch
    fetch('http://localhost:3001/api/users/me', {
      credentials: 'include',
    })
      .then(res => res.json())
      .then(setUser);
  }, []);

  if (!user) return <div>Loading...</div>;

  return <div>Hello, {user.name}</div>;
}
```

## Checking JWT Content

### Decode JWT (for debugging)

```typescript
// Client-side - decode token to see payload
function parseJwt(token: string) {
  const base64Url = token.split('.')[1];
  const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
  const jsonPayload = decodeURIComponent(
    atob(base64)
      .split('')
      .map((c) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
      .join(''),
  );

  return JSON.parse(jsonPayload);
}

const token = getTokenFromCookie('better-auth.session_token');
if (token) {
  const payload = parseJwt(token);
  console.log('User ID:', payload.sub);
  console.log('Email:', payload.email);
  console.log('Role:', payload.role);
  console.log('Expires:', new Date(payload.exp * 1000));
}
```

### Online Tool

Visit https://jwt.io and paste your token to decode it.

## Important: CORS Configuration

For cookies to work across different origins:

### Auth Service (Better Auth)

```typescript
// apps/auth/src/lib/auth.ts
export const auth = betterAuth({
  // ... other config
  trustedOrigins: [
    'http://localhost:5173', // Client app
    'http://localhost:3001', // User service
    'http://localhost:3002', // Event service
  ],
  advanced: {
    crossSubDomainCookies: {
      enabled: true,
    },
  },
});
```

### Other Services (NestJS)

```typescript
// apps/user/src/main.ts
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.enableCors({
    origin: ['http://localhost:5173', 'http://localhost:3000'],
    credentials: true, // Important for cookies!
  });

  await app.listen(3001);
}
bootstrap();
```

## Testing with cURL

### 1. Sign In and Save Cookie

```bash
curl -c cookies.txt -X POST http://localhost:3000/api/v1/auth/sign-in/email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123"
  }'
```

### 2. Use Cookie for Authenticated Request

```bash
curl -b cookies.txt http://localhost:3001/api/users/me
```

### 3. Extract Token and Use Header

```bash
# Extract token from cookie file
TOKEN=$(grep session_token cookies.txt | awk '{print $NF}')

# Use in Authorization header
curl http://localhost:3001/api/users/me \
  -H "Authorization: Bearer $TOKEN"
```

## Troubleshooting

### Token Not Being Set

**Issue**: Cookie not appearing after login

**Solutions**:

- Check `trustedOrigins` in Better Auth config
- Verify `credentials: 'include'` in fetch
- Enable `crossSubDomainCookies` if using subdomains
- Check browser console for CORS errors

### Token Not Being Sent

**Issue**: Requests fail with 401 even though logged in

**Solutions**:

```typescript
// Always include credentials
fetch(url, { credentials: 'include' });

// Check if cookie exists
console.log(document.cookie);

// Check CORS allows credentials
// Response should have: Access-Control-Allow-Credentials: true
```

### Token Expired

**Issue**: Getting "Token expired" errors

**Solution**: Implement token refresh:

```typescript
// In your API wrapper
async function fetchWithRefresh(url, options) {
  let response = await fetch(url, options);

  if (response.status === 401) {
    // Try to refresh session
    await authClient.session.refresh();
    // Retry request
    response = await fetch(url, options);
  }

  return response;
}
```

## Summary

1. ✅ Better Auth automatically generates JWT when configured
2. ✅ JWT is stored in `better-auth.session_token` cookie
3. ✅ Use `credentials: 'include'` to send cookies
4. ✅ Or manually extract and send in `Authorization: Bearer` header
5. ✅ Configure CORS to allow credentials
6. ✅ All services verify JWT locally with public key
