import { LoadingStar } from '@/components/LoadingStar';
import { useTranslations } from 'next-intl';

export default function Loading() {
  const t = useTranslations('Shared');

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
