import { Button } from '@/components/ui/button';
import { FileX } from 'lucide-react';
import { getTranslations } from 'next-intl/server';
import Link from '@/i18n/navigation';

interface EmptyStateProps {
  hasFilters?: boolean;
  locale: string;
}

export default async function EmptyState({
  hasFilters = false,
  locale,
}: EmptyStateProps) {
  const t = await getTranslations({ locale, namespace: 'BlogPage' });

  return (
    <div className="flex min-h-[400px] flex-col items-center justify-center rounded-lg border border-dashed border-border bg-muted/50 p-12 text-center">
      <div className="mb-4 inline-flex items-center justify-center rounded-full bg-muted p-4">
        <FileX className="h-10 w-10 text-muted-foreground" />
      </div>
      <h3 className="mb-2 text-xl font-semibold text-foreground">
        {t('no-posts-found')}
      </h3>
      <p className="mb-6 max-w-md text-muted-foreground">
        {t('no-posts-description')}
      </p>
      {hasFilters && (
        <Button asChild variant="outline">
          <Link href="/blog">{t('clear-filters')}</Link>
        </Button>
      )}
    </div>
  );
}
