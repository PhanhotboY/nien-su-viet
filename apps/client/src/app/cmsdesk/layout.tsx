'use client';

import DashboardLayout from '@/components/cmsdesk/dashboard-layout';
import { authClient } from '@/lib/auth-client';
import { RedirectToSignIn } from '@daveyplate/better-auth-ui';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export default function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter();
  const { data: session, isPending } = authClient.useSession();

  useEffect(() => {
    if (!isPending && !['admin', 'editor'].includes(session?.user.role!)) {
      router.push('/');
    }
  }, [session, isPending]);

  if (isPending) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  return (
    <DashboardLayout>
      <RedirectToSignIn />
      {children}
    </DashboardLayout>
  );
}
