import { redirect } from 'next/navigation';
import { authClient } from '@/lib/auth-client'; // import the auth client

export function generateMetadata() {
  return {
    title: 'Cài đặt trang',
  };
}

export default async function PageManager() {
  const { data: session, error } = await authClient.getSession();

  if (!session) {
    redirect('/cmsdesk/login');
  }

  // const pages = await getPages({ user: auth });
  const pages: any[] = []; // Uncomment above and remove this line when getPages is available

  return <div />;
}
