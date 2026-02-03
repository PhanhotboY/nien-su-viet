export const RATE_LIMIT = {
  // For public/unauthenticated endpoints
  DEFAULT: {
    short: { name: 'short', ttl: 1000, limit: 10 },
    medium: { name: 'medium', ttl: 60 * 1000, limit: 500 },
    long: { name: 'long', ttl: 60 * 60 * 1000, limit: 10000 },
  },
  // For sensitive endpoints (e.g., password reset, email change)
  STRICT: {
    short: { name: 'short', ttl: 1000, limit: 2 },
    medium: { name: 'medium', ttl: 60 * 1000, limit: 20 },
    long: { name: 'long', ttl: 60 * 60 * 1000, limit: 200 },
  },
  // For authentication endpoints (login, register)
  AUTH: {
    short: { name: 'short', ttl: 1000, limit: 1 },
    medium: { name: 'medium', ttl: 60 * 1000, limit: 5 },
    long: { name: 'long', ttl: 60 * 60 * 1000, limit: 20 },
  },
  // For premium/paid users
  PREMIUM: {
    short: { name: 'short', ttl: 1000, limit: 50 },
    medium: { name: 'medium', ttl: 60 * 1000, limit: 500 },
    long: { name: 'long', ttl: 60 * 60 * 1000, limit: 5000 },
  },
  // For internal/admin users
  INTERNAL: {
    short: { name: 'short', ttl: 1000, limit: 100 },
    medium: { name: 'medium', ttl: 60 * 1000, limit: 1000 },
    long: { name: 'long', ttl: 60 * 60 * 1000, limit: 10000 },
  },
};
