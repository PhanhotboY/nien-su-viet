import { SharedPagination } from '@/components/shared';
import MainPostItem from '@/components/website/post/main-post-item';
import MainPostItemLoading from '@/components/website/post/main-post-item-loading';
import { getPublicPosts } from '@/services/post.service';
import { notFound } from 'next/navigation';
import { Suspense } from 'react';

export const revalidate = 0;

interface HomePageProps {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
  params: Promise<{ locale: string }>;
}

export default async function HomePage({
  searchParams,
  params,
}: HomePageProps) {
  const resolvedSearchParams = await searchParams;
  const { locale } = await params;
  // Fetch posts
  const { data, pagination } = await getPublicPosts({
    page: (resolvedSearchParams.page as string) || '1',
    limit: (resolvedSearchParams.limit as string) || '10',
    ...(resolvedSearchParams.q ? { q: resolvedSearchParams.q as string } : {}),
  });

  if (!data) {
    notFound();
  }

  return (
    <div className="container py-4 md:py-8">
      <h1 className="text-3xl md:text-5xl text-center font-bold leading-8 md:leading-20 mb-4 md:mb-8">
        Blog
      </h1>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        {data?.map((post) => (
          <div key={post.id}>
            <Suspense fallback={<MainPostItemLoading />}>
              <MainPostItem post={post} locale={locale} />
            </Suspense>
          </div>
        ))}
      </div>
      {/* Pagination */}
      {pagination.totalPages > 1 && (
        <SharedPagination
          page={pagination.page}
          totalPages={pagination.totalPages}
          baseUrl="/"
          pageUrl="?page="
        />
      )}
    </div>
  );
}
