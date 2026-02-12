import PostTableEmpty from '@/components/cmsdesk/post/post-emtpy-table';
import PostRefreshOnce from '@/components/cmsdesk/post/post-refresh-once';
import PostTableTitle from '@/components/cmsdesk/post/post-table-title';
import { columns } from '@/components/cmsdesk/post/table/columns';
import { DataTable } from '@/components/cmsdesk/post/table/data-table';
import { protectedPostConfig } from '@/config/cmsdesk';
import { findPosts } from '@/services/post.service';
import { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { FC } from 'react';

export const revalidate = 0;

export const metadata: Metadata = {
  title: protectedPostConfig.title,
  description: protectedPostConfig.description,
};

interface PostsPageProps {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}

async function findPostsHandler(query: any) {
  try {
    ('use server');
    const { data } = await findPosts({
      page: (query.page as string) || '1',
      limit: (query.limit as string) || '10',
      ...(query.q ? { q: query.q as string } : {}),
    });
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

const PostsPage: FC<PostsPageProps> = async ({ searchParams }) => {
  const resolvedSearchParams = await searchParams;
  const posts = await findPostsHandler(resolvedSearchParams);

  if (!posts) {
    notFound();
  }
  return (
    <>
      <div className="container p-4 sm:p-6 lg:p-8">
        {posts?.length && posts?.length > 0 ? (
          <>
            <PostTableTitle />
            <DataTable data={posts ? posts : []} columns={columns} />
          </>
        ) : (
          <PostTableEmpty />
        )}
        <PostRefreshOnce />
      </div>
    </>
  );
};

export default PostsPage;
