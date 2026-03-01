import { AccountView } from '@daveyplate/better-auth-ui';
import { accountViewPaths } from '@daveyplate/better-auth-ui/server';

export const dynamicParams = false;

export function generateStaticParams() {
  return Object.values(accountViewPaths).map((path) => ({ path }));
}

export default async function AccountPage({
  params,
}: {
  params: Promise<{ path: string; locale: string }>;
}) {
  const { path, locale } = await params;

  return (
    <main className="container self-center p-4 md:p-6">
      <AccountView
        path={path}
        pathname={`/${locale}/cmsdesk/account`}
        classNames={{
          sidebar: {
            base: 'hidden sticky top-20',
          },
        }}
      />
    </main>
  );
}
