export const config = {
  port: process.env.NODE_PORT || 3000,
  env: process.env.NODE_ENV || 'development',
  trustedOrigins: process.env.ALLOWED_ORIGINS?.split(',') || [],
  authServiceUrl: process.env.AUTH_SERVICE_URL,
  historicalEventServiceHost: process.env.HISTORICAL_EVENT_SERVICE_HOST,
  cmsServiceHost: process.env.CMS_SERVICE_HOST,
} as const;
