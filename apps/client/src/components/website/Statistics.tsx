import { getTranslations } from 'next-intl/server';

export const Statistics = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({
    namespace: 'AboutPage.Statistics',
    locale,
  });

  const stats = [
    {
      quantity: '4000+',
      description: t('items.years'),
    },
    {
      quantity: '1000+',
      description: t('items.events'),
    },
    {
      quantity: '500+',
      description: t('items.figures'),
    },
    {
      quantity: '20+',
      description: t('items.dynasties'),
    },
  ];

  return (
    <section id="statistics">
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-8">
        {stats.map(({ quantity, description }) => (
          <div key={description} className="space-y-2 text-center">
            <h2 className="text-3xl sm:text-4xl font-bold ">{quantity}</h2>
            <p className="text-xl text-muted-foreground">{description}</p>
          </div>
        ))}
      </div>
    </section>
  );
};
