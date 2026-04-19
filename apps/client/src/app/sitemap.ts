import { CLIENT_HOST } from '@/lib/config';
import { getPublicPosts } from '@/services/post.service';
import { getEvents } from '@/services/historical-event.service';
import { MetadataRoute } from 'next';

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const locales = ['vi', 'en'] as const;

  const paths = ['', '/gioi-thieu', '/lien-he', '/blog'];

  const defaultRoutes = locales.map((locale) =>
    paths.map((path) =>
      genSitemap({
        path: `${path}`,
        locale,
        lastModified: new Date(),
        changeFrequency: 'weekly' as const,
        priority: path === '' ? 1 : 0.7,
      }),
    ),
  );

  try {
    // const [{ data: posts }, { data: events }] = await Promise.all([
    //   getPublicPosts(),
    //   getEvents({ page: '1', limit: '100' }),
    // ]);
    const { data: posts } = await getPublicPosts();
    const postRoutes = locales.map(
      (locale) =>
        posts?.map((post) =>
          genSitemap({
            path: `/posts/${post.slug}`,
            locale,
            lastModified: new Date(post.updatedAt as any),
            changeFrequency: 'monthly' as const,
            priority: 0.6,
          }),
        ) || [],
    );

    // const eventRoutes = locales.map(
    //   (locale) =>
    //     events?.map((event) =>
    //       genSitemap({
    //         path: `/su-kien/${event.id}`,
    //         locale,
    //         lastModified: new Date(),
    //         changeFrequency: 'monthly' as const,
    //         priority: 0.6,
    //       }),
    //     ) || [],
    // );

    return [...defaultRoutes, ...postRoutes].flat();
  } catch (e) {
    return defaultRoutes.flat();
  }
}

const genSitemap = ({
  locale,
  path,
  lastModified,
  changeFrequency,
  priority,
}: Omit<MetadataRoute.Sitemap[number], 'url'> & {
  locale: string;
  path: string;
}) => ({
  url: `${CLIENT_HOST}/${locale}${path}`,
  lastModified,
  changeFrequency,
  priority,
});
