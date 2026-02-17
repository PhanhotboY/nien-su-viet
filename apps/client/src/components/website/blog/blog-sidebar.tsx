import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { PostBriefResponseDto } from '@/types/collection';
import { Calendar } from 'lucide-react';
import { getTranslations } from 'next-intl/server';
import Link from 'next/link';

interface BlogSidebarProps {
  recentPosts: PostBriefResponseDto[];
  locale: string;
  className?: string;
}

export default async function BlogSidebar({
  recentPosts,
  locale,
  className,
}: BlogSidebarProps) {
  const t = await getTranslations('BlogPage');

  return (
    <aside className={className}>
      <div className="space-y-6">
        {/* Recent Posts */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-lg">
              <Calendar className="h-5 w-5 text-primary" />
              {t('recent-posts')}
            </CardTitle>
          </CardHeader>
          <CardContent>
            <ul className="space-y-4">
              {recentPosts.map((post) => (
                <li key={post.id}>
                  <Link
                    href={`/posts/${post.slug}`}
                    className="group block transition-colors hover:text-primary"
                  >
                    <h4 className="mb-1 line-clamp-2 text-sm font-medium leading-tight group-hover:underline">
                      {post.title}
                    </h4>
                    <time className="text-xs text-muted-foreground">
                      {new Date(post.createdAt as any).toLocaleDateString(
                        locale === 'vi' ? 'vi-VN' : 'en-US',
                        {
                          year: 'numeric',
                          month: 'long',
                          day: 'numeric',
                        },
                      )}
                    </time>
                  </Link>
                </li>
              ))}
            </ul>
          </CardContent>
        </Card>

        {/* Categories */}
        {/*<Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-lg">
              <TrendingUp className="h-5 w-5 text-secondary" />
              Danh mục
            </CardTitle>
          </CardHeader>
          <CardContent>
            <ul className="space-y-2">
              {categories.map((category) => (
                <li key={category.slug}>
                  <Link
                    href={`/blog?category=${category.slug}`}
                    className="group flex items-center justify-between rounded-md px-3 py-2 text-sm transition-colors hover:bg-accent"
                  >
                    <span className="group-hover:text-primary">
                      {category.name}
                    </span>
                    <Badge variant="secondary" className="text-xs">
                      {category.count}
                    </Badge>
                  </Link>
                </li>
              ))}
            </ul>
          </CardContent>
        </Card>*/}

        {/* Popular Tags */}
        {/*<Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-lg">
              <Tag className="h-5 w-5 text-accent-foreground" />
              {t('popular-tags')}
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2">
              {popularTags.map((tag) => (
                <Link key={tag.slug} href={`/blog?tag=${tag.slug}`}>
                  <Badge
                    variant="outline"
                    className="cursor-pointer transition-colors hover:bg-accent hover:text-accent-foreground"
                  >
                    {tag.name} ({tag.count})
                  </Badge>
                </Link>
              ))}
            </div>
          </CardContent>
        </Card>*/}

        {/* Newsletter CTA */}
        {/*<Card className="border-primary/20 bg-gradient-to-br from-primary/5 to-accent/5">
          <CardHeader>
            <CardTitle className="text-lg">
              {t('subscribe-newsletter')}
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="mb-4 text-sm text-muted-foreground">
              {t('subscription-cta')}
            </p>
            <div className="space-y-2">
              <input
                type="email"
                placeholder={t('enter-email')}
                className="h-10 w-full rounded-md border border-input bg-background px-3 text-sm focus:border-primary focus:outline-none focus:ring-1 focus:ring-ring"
              />
              <button className="h-10 w-full rounded-md bg-primary text-sm font-medium text-primary-foreground transition-colors hover:bg-primary/90">
                {t('subscribe')}
              </button>
            </div>
          </CardContent>
        </Card>*/}
      </div>
    </aside>
  );
}
