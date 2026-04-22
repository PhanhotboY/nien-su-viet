import { Metadata } from 'next';
import { Mail } from 'lucide-react';
import { getTranslations } from 'next-intl/server';

import { Separator } from '@/components/ui/separator';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import Link from '@/i18n/navigation';
import { getMetadata } from '@/content/landing/metadata';
import ContactForm from '@/components/website/contact/ContactForm';
import {
  FacebookIcon,
  XIcon,
  InstagramIcon,
  ThreadsIcon,
} from '@/icons/socials';
import { genMetadata } from '@/lib/metadata.lib';

type Props = {
  params: Promise<{ locale: string }>;
};
export async function generateMetadata({ params }: Props): Promise<Metadata> {
  try {
    const { locale } = await params;
    const metadata = await getMetadata({ locale });

    const t = await getTranslations({ locale, namespace: 'ContactPage' });

    return genMetadata({
      title: `${t('header.title')} - ${metadata.title}`,
      description: t('header.description'),
      locale,
      path: '/lien-he',
      logo: metadata.logo,
    });
  } catch (error) {
    const title = 'Nien Su Viet';
    const description = 'Vietnam history timeline website';

    return genMetadata({ title, description, locale: 'vi', path: '/' });
  }
}

export default async function ContactPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  const appData = await getMetadata({ locale });
  const t = await getTranslations({ locale, namespace: 'ContactPage' });

  return (
    <div className="container mx-auto py-12">
      <div className="space-y-8">
        {/* Header */}
        <div className="text-center space-y-4">
          <h1 className="text-4xl font-bold tracking-tight">
            {t('header.title')}
          </h1>
          <p className="text-muted-foreground text-lg mx-auto">
            {t('header.description')}
          </p>
        </div>

        <Separator />

        <div className="grid md:grid-cols-3 gap-8">
          {/* Contact Information */}
          <div className="md:col-span-1 space-y-6">
            {/* Contact Cards */}
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">{t('info.title')}</CardTitle>
              </CardHeader>

              <CardContent className="space-y-6">
                {/* Address */}
                {/*<div className="flex gap-4">
                  <div className="flex-shrink-0">
                    <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                      <MapPin className="h-5 w-5 text-primary" />
                    </div>
                  </div>

                  <div className="space-y-1">
                    <h3 className="font-semibold">{t('info.address.label')}</h3>
                    <p className="text-sm text-muted-foreground">
                      {addressLines.length
                        ? addressLines.map((line, idx) => (
                            <span key={idx}>
                              {line}
                              <br />
                            </span>
                          ))
                        : t('common.empty')}
                    </p>
                  </div>
                </div>*/}

                {/* Phone */}
                {/*<div className="flex gap-4">
                  <div className="flex-shrink-0">
                    <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                      <Phone className="h-5 w-5 text-primary" />
                    </div>
                  </div>

                  <div className="space-y-1">
                    <h3 className="font-semibold">{t('info.phone.label')}</h3>
                    <p className="text-sm text-muted-foreground">
                      {appData?.msisdn ?? t('common.empty')}
                    </p>
                  </div>
                </div>*/}

                {/* Email */}
                <div className="flex gap-4">
                  <div className="flex-shrink-0">
                    <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                      <Mail className="h-5 w-5 text-primary" />
                    </div>
                  </div>

                  <div className="space-y-1">
                    <h3 className="font-semibold">{t('info.email.label')}</h3>
                    <p className="text-sm text-muted-foreground">
                      {appData?.email ?? t('common.empty')}
                    </p>
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* Social Media */}
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">{t('social.title')}</CardTitle>
                <CardDescription>{t('social.description')}</CardDescription>
              </CardHeader>

              <CardContent>
                <div className="flex gap-3">
                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href={appData?.social?.facebook ?? 'https://facebook.com'}
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label={t('social.aria.facebook')}
                    >
                      <FacebookIcon className="h-5 w-5" />
                    </a>
                  </Button>

                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href={appData?.social?.twitter ?? 'https://twitter.com'}
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label={t('social.aria.twitter')}
                    >
                      <XIcon className="h-5 w-5" />
                    </a>
                  </Button>

                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href={
                        appData?.social?.instagram ?? 'https://instagram.com'
                      }
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label={t('social.aria.instagram')}
                    >
                      <InstagramIcon className="h-5 w-5" />
                    </a>
                  </Button>

                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href={appData?.social?.threads ?? 'https://linkedin.com'}
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label={t('social.aria.linkedin')}
                    >
                      <ThreadsIcon className="h-5 w-5" />
                    </a>
                  </Button>
                </div>
              </CardContent>
            </Card>

            {/* FAQ Link */}
            <Card className="bg-primary/5 border-primary/20">
              <CardContent className="pt-6">
                <h3 className="font-semibold mb-2">{t('faq.title')}</h3>
                <p className="text-sm text-muted-foreground mb-4">
                  {t('faq.description')}
                </p>
                <Button variant="outline" className="w-full" asChild>
                  <Link href="/gioi-thieu#faq">{t('faq.cta')}</Link>
                </Button>
              </CardContent>
            </Card>
          </div>

          {/* Contact Form (Client Component) */}
          <div className="md:col-span-2">
            <ContactForm
              contactShortcuts={{
                email: appData?.email,
              }}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
