import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Card } from '@/components/ui/card';
import { getAvatarFallback, getMinutes, shimmer, toBase64 } from '@/lib/utils';
import { PostBriefResponseDto } from '@/types/collection';
import { format, parseISO } from 'date-fns';
import { enUS, vi } from 'date-fns/locale';
import { CalendarIcon, Clock10Icon, MessageCircleIcon } from 'lucide-react';
import { getTranslations } from 'next-intl/server';
import Image from 'next/image';
import Link from 'next/link';
import React from 'react';
import readingTime from 'reading-time';

export const dynamic = 'force-dynamic';

// async function getPublicImageUrl(postId: string, fileName: string) {

//   if (data && data.publicUrl) return data.publicUrl;

//   return "/images/not-found.jpg";
// }

// async function getComments(postId: string) {
//   const cookieStore = await cookies();
//   const supabase = createClient(cookieStore);
//   const { data: comments, error } = await supabase
//     .from("comments")
//     .select()
//     .eq("postId", postId)
//     .order("created_at", { ascending: true })
//     .returns<Comment[]>();

//   if (error) {
//     console.error(error.message);
//   }
//   return comments;
// }

interface MainPostItemProps {
  post: PostBriefResponseDto;
  locale: string;
}

const MainPostItem: React.FC<MainPostItemProps> = async ({ post, locale }) => {
  const tshared = await getTranslations('Shared');
  const readTime = readingTime('');
  // const comments = await getComments(post.id ? post.id : "");

  return (
    <Card>
      <Link href={`/posts/${post.slug}`}>
        <article className="relative isolate flex flex-col gap-2 rounded-lg sm:gap-8 px-4 sm:px-6 lg:flex-row">
          <div className="relative aspect-video sm:aspect-2/1 lg:aspect-square lg:w-64 lg:shrink-0">
            <Image
              src={post.thumbnail ?? '/images/not-found.jpg'}
              alt={post.title ?? 'Cover'}
              height={256}
              width={256}
              priority
              placeholder={`data:image/svg+xml;base64,${toBase64(
                shimmer(256, 256),
              )}`}
              className="absolute inset-0 h-full w-full rounded-2xl bg-gray-500 object-cover"
            />
            <div className="absolute inset-0 rounded-2xl ring-1 ring-inset ring-gray-900/10" />
          </div>
          <div>
            {/* Desktop category view */}
            {/* <div className="hidden items-center gap-x-3 text-sm sm:flex">
                  <span className="relative z-10 rounded-full bg-gray-100 px-3 py-1.5 font-medium hover:bg-gray-200">
                    {post.categories?.title}
                  </span>
                </div> */}

            <div className="group relative">
              <h3 className="mt-3 text-lg font-semibold leading-6 group-hover:opacity-50">
                <span className="absolute inset-0" />
                {post.title}
              </h3>
              {/* Mobile category and toolbar view*/}
              <div className="mt-2 flex items-center gap-x-3 text-sm sm:hidden">
                {/* <div className="inline-flex items-center">
                      <span className="relative z-10 rounded-full bg-gray-100 px-3 py-1.5 font-medium hover:bg-gray-200">
                        {post.categories?.title}
                      </span>
                    </div> */}
                <div className="inline-flex items-center">
                  <CalendarIcon className="h-4 w-4" />
                  <span className="ml-1">
                    {format(parseISO(post.updatedAt as any), 'dd/MM/yyyy')}
                  </span>
                </div>
                <div className="inline-flex items-center ">
                  <Clock10Icon className="h-4 w-4" />
                  <span className="ml-1">
                    {getMinutes(
                      readTime.minutes ? readTime.minutes : 0,
                      tshared('minute'),
                    )}
                  </span>
                </div>
              </div>
              <p className="mt-3 text-sm leading-6">{post.summary}</p>
              {/* Desktop toolbar view */}
              <div className="mt-3 hidden items-center gap-x-3 text-sm sm:flex">
                <div className="inline-flex items-center">
                  <CalendarIcon className="h-4 w-4" />
                  <span className="ml-1">
                    {format(
                      parseISO(post.updatedAt as any),
                      locale === 'vi' ? 'dd MMMM, yyyy' : 'MMMM dd, yyyy',
                      { locale: locale === 'vi' ? vi : enUS },
                    )}
                  </span>
                </div>
                <div className="inline-flex items-center ">
                  <Clock10Icon className="h-4 w-4" />
                  <span className="ml-1">
                    {getMinutes(readTime.minutes, tshared('minute'))}
                  </span>
                </div>
                {/* <div className="inline-flex items-center">
                      <MessageCircleIcon className="h-4 w-4" />
                      <span className="ml-1">{comments?.length}</span>
                    </div> */}
              </div>
            </div>

            <div className="mt-3 flex border-t border-gray-900/5 pt-2">
              <div className="relative flex items-center gap-x-2">
                <Avatar>
                  <AvatarImage
                    src={post.author.image!}
                    alt={post.author.name}
                  />

                  <AvatarFallback>
                    {getAvatarFallback(post.author.name)}
                  </AvatarFallback>
                </Avatar>
                <div className="text-sm">
                  <p className="font-semibold">{post.author.name}</p>
                  <p className="">{tshared('author')}</p>
                </div>
              </div>
            </div>
          </div>
        </article>
      </Link>
    </Card>
  );
};

export default MainPostItem;
