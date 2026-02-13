import { CLIENT_HOST } from '@/lib/config';
import { getPublicPosts } from '@/services/post.service';
import { MetadataRoute } from 'next';

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const baseUrl = CLIENT_HOST;
  const { data: posts } = await getPublicPosts();

  const routes = ['', '/gioi-thieu', '/lien-he', '/blog'];

  const defaultRoutes = routes.map((route) => ({
    url: `${baseUrl}${route}`,
    lastModified: new Date(),
    changeFrequency: 'weekly' as const,
    priority: route === '' ? 1 : 0.7,
  }));

  const postRoutes =
    posts?.map((post: any) => ({
      url: `${baseUrl}/blog/${post.slug}`,
      lastModified: new Date(post.updatedAt),
      changeFrequency: 'monthly' as const,
      priority: 0.6,
    })) || [];

  return [...defaultRoutes, ...postRoutes];
}
