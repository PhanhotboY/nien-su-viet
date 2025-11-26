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
  return (
    <DashboardLayout>
      <RedirectToSignIn />
      {children}
    </DashboardLayout>
  );
}
