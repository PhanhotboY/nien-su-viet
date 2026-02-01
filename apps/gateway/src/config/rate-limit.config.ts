export const RATE_LIMIT = {
  // For public/unauthenticated endpoints
  DEFAULT: [
    { name: 'short', ttl: 1000, limit: 10 },
    { name: 'medium', ttl: 60 * 1000, limit: 500 },
    { name: 'long', ttl: 60 * 60 * 1000, limit: 10000 },
  ],
  // For sensitive endpoints (e.g., password reset, email change)
  STRICT: [
    { name: 'short', ttl: 1000, limit: 2 },
    { name: 'medium', ttl: 60 * 1000, limit: 20 },
    { name: 'long', ttl: 60 * 60 * 1000, limit: 200 },
  ],
  // For authentication endpoints (login, register)
  AUTH: [
    { name: 'short', ttl: 1000, limit: 1 },
    { name: 'medium', ttl: 60 * 1000, limit: 5 },
    { name: 'long', ttl: 60 * 60 * 1000, limit: 20 },
  ],
  // For premium/paid users
  PREMIUM: [
    { name: 'short', ttl: 1000, limit: 50 },
    { name: 'medium', ttl: 60 * 1000, limit: 500 },
    { name: 'long', ttl: 60 * 60 * 1000, limit: 5000 },
  ],
  // For internal/admin users
  INTERNAL: [
    { name: 'short', ttl: 1000, limit: 100 },
    { name: 'medium', ttl: 60 * 1000, limit: 1000 },
    { name: 'long', ttl: 60 * 60 * 1000, limit: 10000 },
  ],
};
