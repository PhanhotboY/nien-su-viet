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
// import '@/styles/landing-page.css';
// import '@/styles/landing-page.css';

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
