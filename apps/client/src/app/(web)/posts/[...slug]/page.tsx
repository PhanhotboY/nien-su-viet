'@/components/detail/post';
import TextRenderer from '@/components/TextRenderer';
import { seoData } from '@/config/root/seo';
import { getOgImageUrl, getUrl, sanitizeHtml } from '@/lib/utils';
import { getPost } from '@/services/post.service';
import { format, parseISO } from 'date-fns';
import { Metadata } from 'next';
import { notFound } from 'next/navigation';
import readingTime, { ReadTimeResults } from 'reading-time';
import DetailPostHeading from '@/components/detail/post/detail-post-heading';
import { Card, CardContent } from '@/components/ui/card';

export const revalidate = 0;

interface PostPageProps {
  params: Promise<{
    slug: string[];
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
  const truncateDescription = post?.summary?.slice(0, 100) + ('...' as string);
  const slug = '/posts/' + post?.slug;

  if (!post) {
    return {};
  }

  return {
    title: post.title,
    description: truncateDescription,
    authors: {
      name: seoData.author.name,
      url: seoData.author.twitterUrl,
    },
    // openGraph: {
    //   title: post.title as string,
    //   description,
    //   type: 'article',
    //   url: getUrl() + slug,
    //   images: [
    //     {
    //       url: getOgImageUrl(
    //         post.title as string,
    //         truncateDescription as string,
    //         [post.categories?.title as string] as string[],
    //         slug as string,
    //       ),
    //       width: 1200,
    //       height: 630,
    //       alt: post.title as string,
    //     },
    //   ],
    // },
    // twitter: {
    //   card: 'summary_large_image',
    //   title: post.title as string,
    //   description: post.description as string,
    //   images: [
    //     getOgImageUrl(
    //       post.title as string,
    //       truncateDescription as string,
    //       [post.categories?.title as string] as string[],
    //       slug as string,
    //     ),
    //   ],
    // },
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
  const resolvedParams = await params;
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
              thumbnail={post.thumbnail as string}
              // authorName={post.profiles.full_name as string}
              // authorImage={post.profiles.avatar_url as string}
              date={format(parseISO(post.updatedAt!), 'MMMM dd, yyyy')}
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
