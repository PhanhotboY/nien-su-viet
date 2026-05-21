import { Settings } from '@/components/auth';
import { viewPaths } from '@better-auth-ui/core';

export const dynamicParams = false;

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
      <Settings path={path} hideNav />
    </main>
  );
}
