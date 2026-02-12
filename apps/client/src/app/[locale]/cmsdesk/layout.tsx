import { redirect } from 'next/navigation';
import { RedirectToSignIn } from '@daveyplate/better-auth-ui';

import DashboardLayout from '@/components/cmsdesk/dashboard-layout';
import { isEditor } from '@/lib/auth-client';
import { getAuthSession } from '@/helper/auth.helper';

export default async function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { user } = await getAuthSession();

  if (user && !isEditor(user.role)) {
    redirect('/');
  }

  return (
    <DashboardLayout>
      <RedirectToSignIn />
      {children}
    </DashboardLayout>
  );
}
