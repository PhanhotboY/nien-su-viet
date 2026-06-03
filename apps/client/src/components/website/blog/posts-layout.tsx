import { getPublicPosts } from '@/services/post.service';
import FeaturedPost from './featured-post';
import BlogStats from './blog-stats';
import EmptyState from './empty-state';
import MainPostItem from '../post/main-post-item';
import { SharedPagination } from '@/components/shared';
import BlogSidebar from './blog-sidebar';

interface PostsLayoutProps {
  locale: string;
  searchParams?: Record<string, string | string[] | undefined>;
}

const getSearchParam = (
  searchParams: PostsLayoutProps['searchParams'],
  key: string,
) => {
  const value = searchParams?.[key];
  return Array.isArray(value) ? value[0] : value;
};

export async function PostsLayout({
  locale,
  searchParams,
}: PostsLayoutProps) {
  const pageParam = getSearchParam(searchParams, 'page') || '1';
  const limitParam = getSearchParam(searchParams, 'limit') || '10';
  const searchQuery = getSearchParam(searchParams, 'q')?.trim();
  const hasSearchQuery = !!searchQuery;
  const isFirstPage = pageParam === '1';

  const [mainResponse, recentResponse] = await Promise.all([
    getPublicPosts({
      page: pageParam,
      limit: limitParam,
      ...(hasSearchQuery ? { search: searchQuery } : {}),
    }),
    getPublicPosts({
      page: '1',
      limit: '3',
    }),
  ]);

  const mainPosts = mainResponse?.data ?? [];
  const featuredPost =
    isFirstPage && !hasSearchQuery && mainPosts.length > 0
      ? mainPosts[0]
      : null;
  const regularPosts = featuredPost ? mainPosts.slice(1) : mainPosts;
  const pagination = mainResponse?.pagination ?? {
    page: Number(pageParam) || 1,
    limit: Number(limitParam) || 10,
    total: mainPosts.length,
    totalPages: 1,
  };
  const recentPosts = recentResponse?.data ?? [];
  const hasNoPosts = regularPosts.length === 0 && !featuredPost;

  const searchParamsWithoutPage = new URLSearchParams();
  if (searchParams) {
    Object.entries(searchParams).forEach(([key, value]) => {
      if (key === 'page' || value === undefined) return;
      if (Array.isArray(value)) {
        value.forEach((item) => searchParamsWithoutPage.append(key, item));
      } else {
        searchParamsWithoutPage.set(key, value);
      }
    });
  }

  return (
    <>
      {/* Featured Post - Only on first page without search */}
      {featuredPost && (
        <div className="mb-12">
          <FeaturedPost post={featuredPost} locale={locale} />
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
                {regularPosts?.map((post, index) => (
                  <div
                    key={post.id}
                    className="transition-all duration-200 hover:scale-[1.01]"
                  >
                    <MainPostItem
                      post={post}
                      locale={locale}
                      priority={!featuredPost && index === 0}
                    />
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
                      searchParamsWithoutPage?.size == 0 ? '?page=' : '&page='
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
