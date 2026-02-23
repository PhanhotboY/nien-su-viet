import { LoadingStar } from '@/components/LoadingStar';
import { getTranslations } from 'next-intl/server';

export default async function Loading() {
  const t = await getTranslations('Shared');

  return (
    <div className="flex h-screen w-full items-center justify-center bg-gradient-to-br from-background via-background to-muted/20">
      <LoadingStar />

      {/* Optional loading text */}
      <p className="absolute mt-32 text-sm text-muted-foreground animate-pulse">
        {t('loading')}
      </p>
    </div>
  );
}
