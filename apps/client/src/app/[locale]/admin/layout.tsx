import { redirect } from 'next/navigation';
import { RedirectToSignIn } from '@daveyplate/better-auth-ui';

import DashboardLayout from '@/components/admin/dashboard-layout';
import { isAdmin } from '@/lib/auth-client';
import { getAuthSession } from '@/helper/auth.helper';

export default async function AdminLayout({
  children,
  params,
}: Readonly<{
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}>) {
  const { user } = await getAuthSession();
  const { locale } = await params;

  if (user && !isAdmin(user.role)) {
    switch (user.role) {
      case 'editor':
        redirect(`/${locale}/cmsdesk`);
      default:
        redirect(`/${locale}/`);
    }
  }

  return (
    <DashboardLayout>
      <RedirectToSignIn />
      {children}
    </DashboardLayout>
  );
}
