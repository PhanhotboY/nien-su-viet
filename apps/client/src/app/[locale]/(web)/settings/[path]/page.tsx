import { Metadata } from 'next';
import { viewPaths } from '@better-auth-ui/core';
import { Settings } from '@/components/auth';

export const dynamicParams = false;

export const metadata: Metadata = {
  robots: {
    index: false,
  },
};

export function generateStaticParams() {
  return Object.values(viewPaths.settings).map((path) => ({ path }));
}

export default async function AccountPage({
  params,
}: {
  params: Promise<{ path: string; locale: string }>;
}) {
  const { path, locale } = await params;

  return (
    <main className="container self-center p-4 md:p-6">
      <Settings path={path} />
    </main>
  );
}
