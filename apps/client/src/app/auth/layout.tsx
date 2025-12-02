'use client';

import { authClient } from '@/lib/auth-client';
import { useRouter, useSearchParams } from 'next/navigation';

export default function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const queries = useSearchParams();
  const router = useRouter();
  const { data, error } = authClient.useSession();

  let redirectTo = queries.get('redirectTo');
  if (!redirectTo) {
    switch (data?.user.role) {
      case 'admin':
        redirectTo = '/admin';
        break;
      case 'editor':
        redirectTo = '/cmsdesk';
        break;
      default:
        redirectTo = '/';
    }
  }

  if (data && !error) {
    router.push(redirectTo);
  }

  return <>{children}</>;
}
