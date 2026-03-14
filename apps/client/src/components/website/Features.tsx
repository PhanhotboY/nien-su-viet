import { Badge } from '../ui/badge';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { getTranslations } from 'next-intl/server';

interface FeatureItem {
  title: string;
  description: string;
}

export const Features = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({ namespace: 'AboutPage.Features', locale });

  const rawItems = t.raw('items');
  const rawBadges = t.raw('badges');

  // Ensure we're getting arrays, handle cases where t.raw returns objects
  const features: FeatureItem[] = Array.isArray(rawItems)
    ? (rawItems as FeatureItem[])
    : typeof rawItems === 'object' && rawItems !== null
      ? Object.values(rawItems)
      : [];

  const featureList: string[] = Array.isArray(rawBadges)
    ? (rawBadges as string[])
    : typeof rawBadges === 'object' && rawBadges !== null
      ? Object.values(rawBadges)
      : [];

  return (
    <section id="features" className="container py-24 sm:py-32 space-y-8">
      <h2 className="text-3xl lg:text-4xl font-bold md:text-center">
        {t('title.prefix')}{' '}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {t('title.highlight')}
        </span>
      </h2>

      <div className="flex flex-wrap md:justify-center gap-4">
        {featureList.map((feature: string) => (
          <div key={feature}>
            <Badge variant="secondary" className="text-sm">
              {feature}
            </Badge>
          </div>
        ))}
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {features.map(({ title, description }: FeatureItem) => (
          <Card key={title}>
            <CardHeader>
              <CardTitle>{title}</CardTitle>
            </CardHeader>

            <CardContent>{description}</CardContent>

            <CardFooter>{/* reserved */}</CardFooter>
          </Card>
        ))}
      </div>
    </section>
  );
};
