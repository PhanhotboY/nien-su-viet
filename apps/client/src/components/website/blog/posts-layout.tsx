'use client';

import { getPublicPosts } from '@/services/post.service';
import { useLocale } from 'next-intl';
import { useSearchParams } from 'next/navigation';
import { Suspense, useEffect, useMemo, useState } from 'react';
import MainPostItemLoading from '../post/main-post-item-loading';
import FeaturedPost from './featured-post';
import BlogStats from './blog-stats';
import EmptyState from './empty-state';
import MainPostItem from '../post/main-post-item';
import { SharedPagination } from '@/components/shared';
import BlogSidebar from './blog-sidebar';

export function PostsLayout() {
  const locale = useLocale();
  const params = useSearchParams();

  const [recentPosts, setRecentPosts] = useState<
    Awaited<ReturnType<typeof getPublicPosts>>['data']
  >([]);
  useEffect(() => {
    getPublicPosts({
      page: '1',
      limit: '3',
    }).then(({ data }) => {
      setRecentPosts(data);
    });
  }, []);

  const [featuredPost, setFeaturedPost] = useState<
    Awaited<ReturnType<typeof getPublicPosts>>['data'][0] | null
  >(null);
  const [regularPosts, setRegularPosts] = useState<
    Awaited<ReturnType<typeof getPublicPosts>>['data']
  >([]);
  const [pagination, setPagination] = useState<
    Awaited<ReturnType<typeof getPublicPosts>>['pagination']
  >({
    page: 1,
    limit: 10,
    total: 0,
    totalPages: 0,
  });
  const hasSearchQuery = !!params.get('q');
  const isFirstPage = !params.get('page') || params.get('page') === '1';
  useEffect(() => {
    const urlSearchParams = new URLSearchParams(params);
    getPublicPosts({
      page: urlSearchParams.get('page') || '1',
      limit: urlSearchParams.get('limit') || '10',
      ...(urlSearchParams.get('q')
        ? { search: urlSearchParams.get('q') as string }
        : {}),
    }).then(({ data, pagination }) => {
      // Featured post (show first post on first page if no search)
      if (isFirstPage && !hasSearchQuery && data.length > 0) {
        setFeaturedPost(data[0]);
        setRegularPosts(data.slice(1));
      } else {
        setFeaturedPost(null);
        setRegularPosts(data);
      }
      setPagination(pagination);
    });
  }, [params.toString()]);

  const hasNoPosts = regularPosts.length === 0 && !featuredPost;
  const searchParamsWithoutPage = useMemo(() => {
    const urlSearchParams = new URLSearchParams(params);
    urlSearchParams.delete('page');
    return urlSearchParams;
  }, [params.toString()]);

  return (
    <>
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
              <BlogStats totalPosts={pagination.total} locale={locale} />
            </div>
          )}

          {/* Posts Grid or Empty State */}
          {hasNoPosts ? (
            <EmptyState hasFilters={hasSearchQuery} locale={locale} />
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
                    baseUrl={`/${locale}/blog${searchParamsWithoutPage?.toString() ? '?' + searchParamsWithoutPage?.toString() : ''}`}
                    pageUrl={
                      searchParamsWithoutPage?.size == 0 ? '&page=' : '?page='
                    }
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
    </>
  );
}
