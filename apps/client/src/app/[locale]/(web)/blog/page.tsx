import { SharedPagination } from '@/components/shared';
import {
  BlogHeader,
  BlogSearch,
  BlogSidebar,
  BlogStats,
  EmptyState,
  FeaturedPost,
} from '@/components/website/blog';
import MainPostItem from '@/components/website/post/main-post-item';
import MainPostItemLoading from '@/components/website/post/main-post-item-loading';
import { CLIENT_HOST } from '@/lib/config';
import { genMetadata } from '@/lib/metadata.lib';
import { getAppInfo } from '@/services/cms.service';
import { getPublicPosts } from '@/services/post.service';
import { Metadata } from 'next';
import { getTranslations } from 'next-intl/server';
import { notFound } from 'next/navigation';
import { Suspense } from 'react';

export const revalidate = 0;

interface BlogPageProps {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
  params: Promise<{ locale: string }>;
}

export async function generateMetadata({
  params,
}: BlogPageProps): Promise<Metadata> {
  const { locale } = await params;
  const { data } = await getAppInfo();
  const t = await getTranslations('BlogPage');

  return genMetadata({
    title: `${t('title')} - ${data?.title || 'Nien Su Viet'}`,
    description: t('description'),
    locale,
    path: '/blog',
  });
}

export default async function HomePage({
  searchParams,
  params,
}: BlogPageProps) {
  const resolvedSearchParams = await searchParams;
  const { locale } = await params;

  // Fetch posts
  const { data, pagination } = await getPublicPosts({
    page: (resolvedSearchParams.page as string) || '1',
    limit: (resolvedSearchParams.limit as string) || '10',
    ...(resolvedSearchParams.q
      ? { search: resolvedSearchParams.q as string }
      : {}),
  });
  const { data: recentPosts } = await getPublicPosts({
    page: '1',
    limit: '3',
  });

  if (!data) {
    notFound();
  }

  const hasSearchQuery = !!resolvedSearchParams.q;
  const hasNoPosts = data.length === 0;
  const isFirstPage =
    !resolvedSearchParams.page || resolvedSearchParams.page === '1';

  // Featured post (show first post on first page if no search)
  const featuredPost =
    isFirstPage && !hasSearchQuery && data.length > 0 ? data[0] : null;
  const regularPosts = featuredPost ? data.slice(1) : data;

  return (
    <div className="min-h-screen bg-gradient-to-b from-background to-muted/50">
      {/* Header Section */}
      <BlogHeader />

      {/* Main Content */}
      <div className="container pb-8 md:pb-12">
        {/* Search Bar */}
        <div className="mx-auto mb-8 max-w-4xl">
          <BlogSearch />
        </div>

        {/* Featured Post - Only on first page without search */}
        {featuredPost && (
          <div className="mb-12">
            <Suspense fallback={<MainPostItemLoading />}>
              <FeaturedPost post={featuredPost} locale={locale} />
            </Suspense>
          </div>
        )}

        {/* Two Column Layout */}
        <div className="grid grid-cols-1 gap-8 lg:grid-cols-12">
          {/* Main Content Area */}
          <div className="lg:col-span-8">
            {/* Stats and Filters */}
            {!hasNoPosts && (
              <div className="mb-8">
                <BlogStats totalPosts={pagination.total} />
              </div>
            )}

            {/* Posts Grid or Empty State */}
            {hasNoPosts ? (
              <EmptyState hasFilters={hasSearchQuery} />
            ) : (
              <>
                <div className="space-y-8">
                  {regularPosts?.map((post) => (
                    <div
                      key={post.id}
                      className="transition-all duration-200 hover:scale-[1.01]"
                    >
                      <Suspense fallback={<MainPostItemLoading />}>
                        <MainPostItem post={post} locale={locale} />
                      </Suspense>
                    </div>
                  ))}
                </div>

                {/* Pagination */}
                {pagination.totalPages > 1 && (
                  <div className="mt-12">
                    <SharedPagination
                      page={pagination.page}
                      totalPages={pagination.totalPages}
                      baseUrl="/blog"
                      pageUrl="?page="
                    />
                  </div>
                )}
              </>
            )}
          </div>

          {/* Sidebar */}
          <div className="lg:col-span-4">
            <div className="sticky top-4">
              <BlogSidebar recentPosts={recentPosts} locale={locale} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
