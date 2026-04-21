'use client';

import { useTranslations } from 'next-intl';

interface BlogStatsProps {
  totalPosts: number;
  locale: string;
}

export default function BlogStats({ totalPosts, locale }: BlogStatsProps) {
  const t = useTranslations('BlogPage');

  return (
    <div className="mb-6 flex items-center justify-between border-b border-border pb-4">
      <div className="text-sm text-muted-foreground">
        {t('showing-results', { count: totalPosts })}
      </div>
    </div>
  );
}
