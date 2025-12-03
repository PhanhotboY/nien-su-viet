'use client';

import DashboardLayout from '@/components/admin/dashboard-layout';
import { authClient } from '@/lib/auth-client';
import { RedirectToSignIn } from '@daveyplate/better-auth-ui';
import { redirect } from 'next/navigation';
import { useEffect } from 'react';

export default function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { data: session, isPending } = authClient.useSession();

  useEffect(() => {
    if (!isPending && session?.user.role !== 'admin') {
      switch (session?.user.role) {
        case 'editor':
          redirect('/cmsdesk');
        case 'user':
        default:
          redirect('/');
      }
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
