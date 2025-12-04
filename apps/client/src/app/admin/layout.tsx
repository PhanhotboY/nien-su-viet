import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import { RedirectToSignIn } from '@daveyplate/better-auth-ui';

import DashboardLayout from '@/components/admin/dashboard-layout';
import { authClient } from '@/lib/auth-client';

export default async function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { data } = await authClient.getSession({
    fetchOptions: { headers: { Cookie: (await cookies()).toString() } },
  });

  if (data?.user.role !== 'admin') {
    switch (data?.user.role) {
      case 'editor':
        redirect('/cmsdesk');
      case 'user':
      default:
        redirect('/');
    }
  }

  return (
    <DashboardLayout>
      <RedirectToSignIn />
      {children}
    </DashboardLayout>
  );
}
