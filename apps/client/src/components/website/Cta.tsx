import { Button } from '../ui/button';
import { getTranslations } from 'next-intl/server';

export const Cta = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({ namespace: 'AboutPage.Cta', locale });

  return (
    <section id="cta" className="bg-muted/50 py-16 my-24 sm:my-32">
      <div className="container lg:grid lg:grid-cols-2 place-items-center">
        <div className="lg:col-start-1">
          <h2 className="text-3xl md:text-4xl font-bold ">
            {t('title.prefix')}
            <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
              {' '}
              {t('title.highlight')}{' '}
            </span>
            {t('title.suffix')}
          </h2>
          <p className="text-muted-foreground text-xl mt-4 mb-8 lg:mb-0">
            {t('description')}
          </p>
        </div>

        <div className="space-y-4 lg:col-start-2">
          <Button className="w-full md:mr-4 md:w-auto">
            {t('actions.primary')}
          </Button>
          <Button variant="outline" className="w-full md:w-auto">
            {t('actions.secondary')}
          </Button>
        </div>
      </div>
    </section>
  );
};
