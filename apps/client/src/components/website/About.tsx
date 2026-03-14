import { Statistics } from './Statistics';
import Image from 'next/image';
import { getTranslations } from 'next-intl/server';

export const About = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({ namespace: 'AboutPage', locale });

  return (
    <section id="about" className="container py-24 sm:py-32">
      <div className="bg-muted/50 border rounded-lg py-12">
        <div className="px-6 flex flex-col-reverse md:flex-row gap-8 md:gap-12">
          {/*<Image
            src={pilot}
            alt=""
            className="w-[300px] object-contain rounded-lg"
          />*/}
          <div className="bg-green-0 flex flex-col justify-between">
            <div className="pb-6">
              <h2 className="text-3xl md:text-4xl font-bold">
                <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
                  {t('about')}{' '}
                </span>
                {t('title')}
              </h2>
              <p className="text-xl text-muted-foreground mt-4">
                {t('description')}
              </p>
            </div>

            <Statistics locale={locale} />
          </div>
        </div>
      </div>
    </section>
  );
};
