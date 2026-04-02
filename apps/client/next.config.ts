import type { NextConfig } from 'next';
import createNextIntlPlugin from 'next-intl/plugin';

const nextConfig: NextConfig = {
  /* config options here */
  typescript: {
    ignoreBuildErrors: false,
  },
  images: {
    dangerouslyAllowLocalIP: process.env.NODE_ENV === 'development',
    remotePatterns: [
      {
        hostname: 'localhost',
      },
      {
        hostname: 'res.cloudinary.com',
      },
    ],
  },
  experimental: {
    // turbopackUseSystemTlsCerts: true, // how many times Next.js will retry failed page generation attempts
    // before failing the build
    staticGenerationRetryCount: 1,
    // how many pages will be processed per worker
    staticGenerationMaxConcurrency: 4,
    // the minimum number of pages before spinning up a new export worker
    staticGenerationMinPagesPerWorker: 25,
  },
  async redirects() {
    return [
      {
        source: '/',
        destination: '/vi',
        permanent: false,
      },
    ];
  },
};

const withNextIntl = createNextIntlPlugin();

export default withNextIntl(nextConfig);
