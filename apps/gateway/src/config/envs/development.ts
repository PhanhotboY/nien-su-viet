export const config = {
  port: process.env.NODE_PORT || 3000,
  env: process.env.NODE_ENV || 'development',
  trustedOrigins: process.env.ALLOWED_ORIGINS?.split(',') || [],
  authServiceEndpoint: process.env.AUTH_SERVICE_ENDPOINT,
  cmsServiceEndpoint: process.env.CMS_SERVICE_ENDPOINT,
  historicalEventServiceHost: process.env.HISTORICAL_EVENT_SERVICE_HOST,
  betterAuth: {
    secret: process.env.BETTER_AUTH_SECRET || '',
    cookieDomain: process.env.COOKIE_DOMAIN || 'localhost',
    cookiePrefix: process.env.AUTH_COOKIE_PREFIX || 'nsv-auth',
  },
} as const;
