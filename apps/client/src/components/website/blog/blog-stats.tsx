import { getTranslations } from 'next-intl/server';

interface BlogStatsProps {
  totalPosts: number;
}

export default async function BlogStats({ totalPosts }: BlogStatsProps) {
  const t = await getTranslations('BlogPage');

  return (
    <div className="mb-6 flex items-center justify-between border-b border-border pb-4">
      <div className="text-sm text-muted-foreground">
        {t('showing-results', { count: totalPosts })}
      </div>
    </div>
  );
}
