import { About } from '@/components/website/About';
import { Cta } from '@/components/website/Cta';
import { FAQ } from '@/components/website/FAQ';
import { Features } from '@/components/website/Features';
import { Hero } from '@/components/website/Hero';
import { HowItWorks } from '@/components/website/HowItWorks';
import { Newsletter } from '@/components/website/Newsletter';
import { Pricing } from '@/components/website/Pricing';
import { Services } from '@/components/website/Services';
import { Sponsors } from '@/components/website/Sponsors';
import { Team } from '@/components/website/Team';
import { Testimonials } from '@/components/website/Testimonials';
import { getMetadata } from '@/content/landing/metadata';
import { genMetadata } from '@/lib/metadata.lib';
import { Metadata } from 'next';
import { getTranslations } from 'next-intl/server';

type Props = {
  params: Promise<{ locale: string }>;
};
export async function generateMetadata({ params }: Props): Promise<Metadata> {
  try {
    const { locale } = await params;
    const metadata = await getMetadata({ locale });

    return genMetadata({
      title: metadata.title,
      description: metadata.description,
      locale,
      path: '/gioi-thieu',
      logo: metadata.logo,
    });
  } catch (error) {
    const title = 'Nien Su Viet';
    const description = 'Vietnam history timeline website';

    return genMetadata({ title, description, locale: 'vi', path: '/' });
  }
}

export default async function App({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;

  return (
    <main>
      <Hero locale={locale} />
      <Sponsors locale={locale} />
      <About locale={locale} />
      <HowItWorks locale={locale} />
      <Features locale={locale} />
      <Services locale={locale} />
      <Cta locale={locale} />
      <Testimonials locale={locale} />
      <Team locale={locale} />
      <Pricing locale={locale} />
      <Newsletter locale={locale} />
      <FAQ locale={locale} />
    </main>
  );
}
