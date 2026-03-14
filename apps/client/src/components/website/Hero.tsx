import { Button } from '../ui/button';
import { buttonVariants } from '../ui/button';
import { HeroCards } from './HeroCards';
import { GitHubLogoIcon } from '@radix-ui/react-icons';
import Link from '@/i18n/navigation';
import { getTranslations } from 'next-intl/server';

export const Hero = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({ namespace: 'AboutPage.Hero', locale });

  return (
    <section className="container grid lg:grid-cols-2 place-items-center py-20 md:py-32 gap-10">
      <div className="text-center lg:text-start space-y-6">
        <main className="text-5xl md:text-6xl font-bold">
          <h1 className="inline">
            <span className="inline bg-gradient-to-r from-secondary/50 to-secondary text-transparent bg-clip-text">
              {t('brand')}
            </span>{' '}
          </h1>{' '}
          <h2 className="inline">
            <span className="inline bg-gradient-to-r from-primary/80 via-primary to-primary text-transparent bg-clip-text">
              {t('headline.emphasis')}
            </span>{' '}
            {t('headline.rest')}
          </h2>
        </main>

        <p className="text-xl text-muted-foreground md:w-10/12 mx-auto lg:mx-0">
          {t('description')}
        </p>

        <div className="space-y-4 md:space-y-0 md:space-x-4">
          <Button className="w-full md:w-1/3">{t('cta.primary')}</Button>

          <Link
            rel="noreferrer noopener"
            href="#about"
            className={`w-full md:w-1/3 ${buttonVariants({
              variant: 'outline',
            })}`}
          >
            {t('cta.secondary')}
          </Link>
        </div>
      </div>

      {/* Hero cards sections */}
      <div className="z-10">
        <HeroCards locale={locale} />
      </div>

      {/* Shadow effect */}
      <div className="shadow"></div>
    </section>
  );
};
