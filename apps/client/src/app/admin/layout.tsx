import DashboardLayout from '@/components/admin/dashboard-layout';
import { authClient } from '@/lib/auth-client';
import { RedirectToSignIn } from '@daveyplate/better-auth-ui';
import { headers } from 'next/headers';
import { redirect } from 'next/navigation';

export default async function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const reqHeaders = await headers();
  const { data } = await authClient.getSession({
    fetchOptions: { headers: reqHeaders },
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
