import { redirect } from 'next/navigation';
import { authClient } from '@/lib/auth-client';
import { cookies } from 'next/headers';

export default async function LoginLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookiesStore = await cookies();
  const { data, error } = await authClient
    .getSession({
      fetchOptions: { headers: { cookie: cookiesStore.toString() } },
    })
    .catch((e) => ({ data: null, error: e }));

  if (data && !error) {
    redirect('/admin');
  }

  return children;
}
