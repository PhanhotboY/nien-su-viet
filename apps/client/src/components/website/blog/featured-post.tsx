import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Badge } from '@/components/ui/badge';
import { Card } from '@/components/ui/card';
import { getAvatarFallback, shimmer, toBase64 } from '@/lib/utils';
import { PostBriefResponseDto } from '@/types/collection';
import { format, parseISO } from 'date-fns';
import { enUS, vi } from 'date-fns/locale';
import { Calendar, Clock, Star } from 'lucide-react';
import { getTranslations } from 'next-intl/server';
import Image from 'next/image';
import Link from 'next/link';
import readingTime from 'reading-time';

interface FeaturedPostProps {
  post: PostBriefResponseDto;
  locale: string;
}

export default async function FeaturedPost({
  post,
  locale,
}: FeaturedPostProps) {
  const tshared = await getTranslations('Shared');
  const tblog = await getTranslations('BlogPage');
  const readTime = readingTime(post.summary || '');

  return (
    <Card className="group relative overflow-hidden border-2 border-primary/20 bg-gradient-to-br from-primary/5 via-background to-secondary/5 shadow-xl transition-all duration-300 hover:shadow-2xl">
      {/* Featured Badge */}
      <div className="absolute right-4 top-4 z-10">
        <Badge className="bg-primary text-primary-foreground shadow-lg hover:bg-primary/90">
          <Star className="mr-1 h-3 w-3 fill-current" />
          {tblog('featured-posts')}
        </Badge>
      </div>

      <Link href={`/posts/${post.slug}`}>
        <div className="grid gap-6 md:grid-cols-2 px-4 md:px-8">
          {/* Image Section */}
          <div className="relative h-64 overflow-hidden md:h-full">
            <Image
              src={post.thumbnail ?? '/assets/image/not-found.webp'}
              alt={post.title ?? 'Featured post'}
              fill
              priority
              placeholder={`data:image/svg+xml;base64,${toBase64(
                shimmer(600, 400),
              )}`}
              className="object-cover transition-transform duration-500 group-hover:scale-110"
            />
            <div className="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent md:bg-gradient-to-r" />
          </div>

          {/* Content Section */}
          <div className="flex flex-col justify-center p-6 md:p-8">
            {/* Title */}
            <h2 className="mb-3 text-2xl font-bold leading-tight text-foreground transition-colors group-hover:text-primary md:text-3xl lg:text-4xl">
              {post.title}
            </h2>

            {/* Summary */}
            <p className="mb-4 line-clamp-3 text-base text-muted-foreground md:text-lg">
              {post.summary}
            </p>

            {/* Metadata */}
            <div className="mb-4 flex flex-wrap items-center gap-4 text-sm text-muted-foreground">
              <div className="flex items-center gap-1.5">
                <Calendar className="h-4 w-4" />
                <span>
                  {format(
                    parseISO(post.updatedAt as any),
                    locale === 'vi' ? 'dd MMMM, yyyy' : 'MMMM dd, yyyy',
                    { locale: locale === 'vi' ? vi : enUS },
                  )}
                </span>
              </div>
              <div className="flex items-center gap-1.5">
                <Clock className="h-4 w-4" />
                <span>
                  {Math.ceil(readTime.minutes)} {tshared('minute')}
                </span>
              </div>
            </div>

            {/* Author */}
            <div className="flex items-center gap-3 border-t border-border pt-4">
              <Avatar className="h-10 w-10">
                <AvatarImage src={post.author.image!} alt={post.author.name} />
                <AvatarFallback>
                  {getAvatarFallback(post.author.name)}
                </AvatarFallback>
              </Avatar>
              <div>
                <p className="font-semibold text-foreground">
                  {post.author.name}
                </p>
                <p className="text-xs text-muted-foreground">
                  {tshared('author')}
                </p>
              </div>
            </div>

            {/* Read More Link */}
            <div className="mt-4">
              <span className="inline-flex items-center text-sm font-semibold text-primary transition-all group-hover:translate-x-1">
                {tblog('read-more')}
                <svg
                  className="ml-1 h-4 w-4"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 5l7 7-7 7"
                  />
                </svg>
              </span>
            </div>
          </div>
        </div>
      </Link>
    </Card>
  );
}
