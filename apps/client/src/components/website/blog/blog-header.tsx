import { BookOpen } from 'lucide-react';
import { getTranslations } from 'next-intl/server';

export default async function BlogHeader() {
  const t = await getTranslations('BlogPage');

  return (
    <div className="relative overflow-hidden bg-gradient-to-b from-muted/50 to-background">
      <div className="absolute inset-0 bg-[radial-gradient(circle_at_50%_50%,hsl(var(--muted)/0.1),transparent_50%)]" />
      <div className="container relative py-12 md:py-16">
        <div className="mx-auto max-w-3xl text-center">
          <h1 className="mb-4 text-4xl font-bold tracking-tight text-foreground md:text-5xl lg:text-6xl">
            {t('title')}
          </h1>
          <p className="text-lg text-muted-foreground md:text-xl">
            {t('description')}
          </p>
        </div>
      </div>
    </div>
  );
}
