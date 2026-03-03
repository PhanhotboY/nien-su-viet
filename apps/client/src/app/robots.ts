import type { MetadataRoute } from 'next';

export default function robots(): MetadataRoute.Robots {
  return {
    rules: {
      userAgent: '*',
      allow: '/',
      disallow: ['/admin', '/auth', '/api', '/cmsdesk'],
    },
    sitemap: process.env.NEXT_PUBLIC_CLIENT_HOST + '/sitemap.xml',
  };
}
