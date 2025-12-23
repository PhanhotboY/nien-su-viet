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

export default function App() {
  return (
    <main>
      <Hero />
      <Sponsors />
      <About />
      <HowItWorks />
      <Features />
      <Services />
      <Cta />
      <Testimonials />
      <Team />
      <Pricing />
      <Newsletter />
      <FAQ />
    </main>
  );
}
