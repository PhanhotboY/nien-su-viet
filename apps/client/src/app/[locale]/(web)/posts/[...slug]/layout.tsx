import { DetailPostHeader } from '@/components/detail/post';
import { getPost, getPublicPosts } from '@/services/post.service';
import { notFound } from 'next/navigation';

export async function generateStaticParams() {
  try {
    const { data: posts } = await getPublicPosts({
      page: '1',
      limit: '1000',
    });
    return posts.map((post) => ({
      slug: post.slug.split('/'),
    }));
  } catch (error) {
    console.error('Error fetching posts for static params:', error);
    return [];
  }
}

async function getPostHandler(params: { slug: string[] }) {
  'use server';
  const slug = params?.slug?.join('/');

  if (!slug) {
    notFound();
  }

  const response = await getPost(slug);

  if (!response.data) {
    notFound();
  }

  return response.data;
}

export default async function MainLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: Promise<{
    slug: string[];
  }>;
}) {
  const resolvedParams = await params;
  const post = await getPostHandler(resolvedParams);

  if (!post) {
    notFound();
  }
  return (
    <>
      <DetailPostHeader title={post.title as string} />
      <main className="min-h-full bg-background py-3">
        <div className="mx-auto max-w-7xl px-6 lg:px-8">
          <div className="mx-auto max-w-4xl">{children}</div>
        </div>
      </main>
    </>
  );
}
