import { SharedPagination } from '@/components/shared';
import {
  BlogHeader,
  BlogSearch,
  BlogSidebar,
  BlogStats,
  EmptyState,
  FeaturedPost,
} from '@/components/website/blog';
import { PostsLayout } from '@/components/website/blog/posts-layout';
import MainPostItem from '@/components/website/post/main-post-item';
import MainPostItemLoading from '@/components/website/post/main-post-item-loading';
import { genMetadata } from '@/lib/metadata.lib';
import { getPublicPosts } from '@/services/post.service';
import { Metadata } from 'next';
import { getTranslations } from 'next-intl/server';
import { notFound } from 'next/navigation';
import { Suspense } from 'react';

interface BlogPageProps {
  params: Promise<{ locale: string }>;
}

export async function generateMetadata({
  params,
}: BlogPageProps): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'BlogPage' });

  return genMetadata({
    title: `${t('title')} - Nien Su Viet`,
    description: t('description'),
    locale,
    path: '/blog',
  });
}

export default async function HomePage({ params }: BlogPageProps) {
  const { locale } = await params;

  return (
    <div className="min-h-screen bg-gradient-to-b from-background to-muted/50">
      {/* Header Section */}
      <BlogHeader locale={locale} />

      {/* Main Content */}
      <div className="container pb-8 md:pb-12">
        {/* Search Bar */}
        <div className="mx-auto mb-8 max-w-4xl">
          <BlogSearch />
        </div>

        <PostsLayout />
      </div>
    </div>
  );
}
