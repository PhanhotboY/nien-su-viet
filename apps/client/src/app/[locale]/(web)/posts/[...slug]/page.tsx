'@/components/detail/post';
import TextRenderer from '@/components/TextRenderer';
import { seoData } from '@/config/root/seo';
import { getPost } from '@/services/post.service';
import { format, parseISO } from 'date-fns';
import { vi, enUS } from 'date-fns/locale';
import { Metadata } from 'next';
import { notFound } from 'next/navigation';
import readingTime, { ReadTimeResults } from 'reading-time';
import DetailPostHeading from '@/components/detail/post/detail-post-heading';
import { Card, CardContent } from '@/components/ui/card';
import { CLIENT_HOST } from '@/lib/config';

export const revalidate = 0;

interface PostPageProps {
  params: Promise<{
    slug: string[];
    locale: string;
  }>;
}

// async function getBookmark(postId: string, userId: string) {
//   if (postId && userId) {
//     const bookmark = {
//       id: postId,
//       user_id: userId,
//     };
//     const response = await GetBookmark(bookmark);

//     return response;
//   }
// }

async function getPostHandler(params: { slug: string[] }) {
  'use server';
  const slug = params?.slug?.join('/');
  const response = await getPost(slug);

  if (!response.data) {
    notFound();
  }

  return response.data;
}

export async function generateMetadata({
  params,
}: PostPageProps): Promise<Metadata> {
  const resolvedParams = await params;
  const post = await getPostHandler(resolvedParams);
  const description = post?.summary?.slice(0, 100) + ('...' as string);
  const slug = '/posts/' + post?.slug;

  if (!post) {
    return {};
  }

  return {
    title: post.title,
    description,
    authors: {
      name: seoData.author.name,
      url: seoData.author.twitterUrl,
    },
    openGraph: {
      title: post.title as string,
      description,
      type: 'article',
      url: CLIENT_HOST + slug,
      images: [
        {
          url: post.thumbnail!,
          width: 1200,
          height: 630,
          alt: post.title as string,
        },
      ],
    },
    twitter: {
      card: 'summary_large_image',
      title: post.title as string,
      description,
      images: [post.thumbnail!],
    },
  };
}

// async function getComments(postId: string) {
//   const { data: comments, error } = await supabase
//     .from('comments')
//     .select('*, profiles(*)')
//     .eq('post_id', postId)
//     .order('created_at', { ascending: true })
//     .returns<CommentWithProfile[]>();

//   if (error) {
//     console.error(error.message);
//   }
//   return comments;
// }

export default async function PostPage({ params }: PostPageProps) {
  const { locale, ...resolvedParams } = await params;
  // Get post data
  const post = await getPostHandler(resolvedParams);
  if (!post) {
    notFound();
  }
  // Set post views

  // Get bookmark status
  // const isBookmarked = await getBookmark(post.id as string, user?.id as string);

  // // Get comments
  // const comments = await getComments(post.id as string);
  const readTime = readingTime(post.content ? post.content : '');

  return (
    <Card>
      <div className="mx-auto max-w-7xl px-0 sm:px-8">
        <CardContent>
          {/*<div className="mx-auto max-w-4xl rounded-lg bg-white px-6 py-4 shadow-sm shadow-gray-300 ring-1 ring-black/5 sm:px-14 sm:py-10">*/}
          <div className="relative mx-auto max-w-4xl py-2">
            {/* Heading */}
            <DetailPostHeading
              id={post.id}
              title={post.title as string}
              thumbnail={post.thumbnail || '/assets/image/not-found.webp'}
              authorName={post.author.name}
              authorImage={post.author.image!}
              date={format(
                parseISO(post.createdAt as any),
                locale === 'vi' ? 'dd MMMM, yyyy' : 'MMMM dd, yyyy',
                {
                  locale: locale === 'vi' ? vi : enUS,
                },
              )}
              // category={post.categories?.title as string}
              readTime={readTime as ReadTimeResults}
            />
            {/* Top Floatingbar */}
            {/*<div className="mx-auto">
                <DetailPostFloatingBar
                    id={post.id as string}
                    title={post.title as string}
                    text={post.summary as string}
                    url={`${getUrl()}${encodeURIComponent(
                      `/posts/${post.slug}`,
                    )}`}
                    // totalComments={comments?.length}
                    // isBookmarked={isBookmarked}
                  />
              </div>*/}
            {/*</div>*/}
            {/* Content */}
            <main className="relative mx-auto max-w-3xl border-slate-500/50 py-5">
              <TextRenderer content={post.content} />
            </main>
            {/*<div className="mx-auto mt-10">*/}
            {/* Bottom Floatingbar */}
            {/*<DetailPostFloatingBar
                  id={post.id as string}
                  title={post.title as string}
                  text={post.summary as string}
                  url={`${getUrl()}${encodeURIComponent(
                    `/posts/${post.slug}`,
                  )}`}
                  // totalComments={comments?.length}
                  // isBookmarked={isBookmarked}
                />*/}
            {/*</div>*/}
          </div>
        </CardContent>
        {/*<DetailPostComment
              postId={post.id as string}
              comments={comments as CommentWithProfile[]}
            />*/}
      </div>
      {/*<DetailPostScrollUpButton />*/}
    </Card>
  );
}
