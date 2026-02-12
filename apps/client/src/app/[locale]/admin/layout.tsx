import { redirect } from 'next/navigation';
import { RedirectToSignIn } from '@daveyplate/better-auth-ui';

import DashboardLayout from '@/components/admin/dashboard-layout';
import { isAdmin } from '@/lib/auth-client';
import { getAuthSession } from '@/helper/auth.helper';

export default async function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { user } = await getAuthSession();

  if (user && !isAdmin(user.role)) {
    switch (user.role) {
      case 'editor':
        redirect('/cmsdesk');
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
