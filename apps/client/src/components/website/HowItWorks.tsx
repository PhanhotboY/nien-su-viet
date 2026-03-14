import { Card, CardContent, CardHeader, CardTitle } from '../ui/card';
import { MedalIcon, MapIcon, PlaneIcon, GiftIcon } from '../Icons';
import { JSX } from 'react';
import { getTranslations } from 'next-intl/server';

interface FeatureProps {
  icon: JSX.Element;
  title: string;
  description: string;
}

export const HowItWorks = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({
    namespace: 'AboutPage.HowItWorks',
    locale,
  });

  const features: FeatureProps[] = [
    {
      icon: <MedalIcon />,
      title: t('features.timeline.title'),
      description: t('features.timeline.description'),
    },
    {
      icon: <MapIcon />,
      title: t('features.map.title'),
      description: t('features.map.description'),
    },
    {
      icon: <PlaneIcon />,
      title: t('features.library.title'),
      description: t('features.library.description'),
    },
    {
      icon: <GiftIcon />,
      title: t('features.learning.title'),
      description: t('features.learning.description'),
    },
  ];

  return (
    <section id="howItWorks" className="container text-center py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold ">
        {t('title.prefix')}{' '}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {t('title.highlight')}{' '}
        </span>
        {t('title.suffix')}
      </h2>

      <p className="md:w-3/4 mx-auto mt-4 mb-8 text-xl text-muted-foreground">
        {t('subtitle')}
      </p>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        {features.map(({ icon, title, description }: FeatureProps) => (
          <Card key={title} className="bg-muted/50">
            <CardHeader>
              <CardTitle className="grid gap-4 place-items-center">
                {icon}
                {title}
              </CardTitle>
            </CardHeader>
            <CardContent>{description}</CardContent>
          </Card>
        ))}
      </div>
    </section>
  );
};
