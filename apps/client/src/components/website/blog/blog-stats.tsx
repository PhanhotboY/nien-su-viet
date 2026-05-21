import { getTranslations } from 'next-intl/server';

interface BlogStatsProps {
  totalPosts: number;
  locale: string;
}

export default async function BlogStats({
  totalPosts,
  locale,
}: BlogStatsProps) {
  const t = await getTranslations({ locale, namespace: 'BlogPage' });

  return (
    <div className="mb-6 flex items-center justify-between border-b border-border pb-4">
      <div className="text-sm text-muted-foreground">
        {t('showing-results', { count: totalPosts })}
      </div>
    </div>
  );
}
