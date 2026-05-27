import { Metadata } from 'next';
import { getTranslations } from 'next-intl/server';
import Link from '@/i18n/navigation';
import {
  ArrowLeft,
  Users,
  Calendar,
  BookOpen,
  Briefcase,
  UtensilsCrossed,
  Shirt,
  Home,
} from 'lucide-react';
import Image from 'next/image';

import { getMetadata } from '@/content/landing/metadata';
import { genMetadata } from '@/lib/metadata.lib';
import { ethnicGroupsMap } from '@/data/vi/vietnamese-ethnic-groups';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

type Props = {
  params: Promise<{ locale: string; slug: string }>;
};

export async function generateStaticParams() {
  try {
    return Array.from(ethnicGroupsMap.keys()).map((slug) => ({
      slug,
    }));
  } catch (error) {
    console.error('Error fetching posts for static params:', error);
    return [];
  }
}

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  const { locale, slug } = await params;
  const path = `/dan-toc/${slug}`;
  const group = ethnicGroupsMap.get(slug);

  try {
    const metadata = await getMetadata({ locale });
    const t = await getTranslations({ locale, namespace: 'EthnicGroupPage' });

    if (!group) {
      return genMetadata({
        title: `Not found - ${metadata.title}`,
        description: 'Ethnic group not found',
        locale,
        path,
        logo: metadata.logo,
      });
    }

    return genMetadata({
      title: `${group.detail.title || group.name} - ${metadata.title}`,
      description: group.detail.history || t('description'),
      locale,
      path,
      logo: metadata.logo,
    });
  } catch (error) {
    const title = 'Nien Su Viet';
    const description = 'Vietnam history timeline website';

    return genMetadata({ title, description, locale: 'vi', path });
  }
}

export default async function DanTocDetailPage({ params }: Props) {
  const { locale, slug } = await params;
  const ethnicGroup = ethnicGroupsMap.get(slug);
  const t = await getTranslations({ locale, namespace: 'EthnicGroupPage' });

  if (!ethnicGroup) {
    return (
      <div className="min-h-screen bg-background">
        <div className="container mx-auto px-4 py-24">
          <div className="text-center">
            <h1 className="text-3xl font-bold text-foreground mb-2">
              {t('notFound')}
            </h1>
            <p className="text-muted-foreground mb-6">
              {t('notFoundDescription')}
            </p>
            <Link href="/dan-toc" className="text-primary hover:underline">
              <ArrowLeft className="h-4 w-4 inline mr-1" />
              {t('backToList')}
            </Link>
          </div>
        </div>
      </div>
    );
  }

  const { detail } = ethnicGroup;

  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <div className="border-b border-border/40 bg-card/30 backdrop-blur-sm">
        <div className="container mx-auto px-4 py-8">
          <Link
            href="/dan-toc"
            className="inline-flex items-center text-sm text-muted-foreground hover:text-foreground transition-colors mb-6"
          >
            <ArrowLeft className="h-4 w-4 mr-2" />
            {t('backToList')}
          </Link>

          <div className="flex items-start justify-between mb-4">
            <div>
              <h1 className="text-4xl md:text-5xl font-bold text-foreground">
                {detail.title || ethnicGroup.name}
              </h1>
            </div>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="container mx-auto px-4 py-12">
        <div className="grid gap-8 lg:grid-cols-3">
          {/* Left Column - Quick Info */}
          <div className="lg:col-span-1 space-y-6">
            {/* Info Cards */}
            {detail.population && (
              <Card>
                <CardHeader>
                  <CardTitle className="text-base flex items-center gap-2">
                    <Users className="h-5 w-5 text-primary" />
                    {t('population')}
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <p className="text-sm text-muted-foreground leading-relaxed">
                    {detail.population}
                  </p>
                </CardContent>
              </Card>
            )}

            {detail.language && (
              <Card>
                <CardHeader>
                  <CardTitle className="text-base flex items-center gap-2">
                    <BookOpen className="h-5 w-5 text-primary" />
                    {t('language')}
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <p className="text-sm text-muted-foreground leading-relaxed">
                    {detail.language}
                  </p>
                </CardContent>
              </Card>
            )}

            {detail.self_name && (
              <Card>
                <CardHeader>
                  <CardTitle className="text-base">{t('selfName')}</CardTitle>
                </CardHeader>
                <CardContent>
                  <p className="text-sm font-medium">{detail.self_name}</p>
                </CardContent>
              </Card>
            )}
          </div>

          {/* Right Column - Detailed Tabs */}
          <Card className="lg:col-span-2">
            <CardContent>
              <article className="w-full [&_h2]:mb-2 [&>section:not(&>section:first-child)_h2]:mt-6">
                {/* Overview Tab */}
                {detail.other_names && (
                  <section>
                    <h2 className="text-lg font-bold">{t('otherNames')}</h2>
                    <p className="text-sm text-muted-foreground leading-relaxed">
                      {detail.other_names}
                    </p>
                  </section>
                )}

                {detail.local_groups && (
                  <section>
                    <h2 className="text-lg font-bold">{t('localGroups')}</h2>
                    <p className="text-sm text-muted-foreground leading-relaxed">
                      {detail.local_groups}
                    </p>
                  </section>
                )}

                {detail.history && (
                  <section>
                    <h2 className="text-lg flex items-center gap-2 font-bold">
                      <Calendar className="h-5 w-5 text-primary font-bold" />
                      {t('history')}
                    </h2>
                    <p className="text-sm text-muted-foreground leading-relaxed">
                      {detail.history}
                    </p>
                  </section>
                )}

                {detail.social_relations && (
                  <section>
                    <h2 className="text-lg font-bold">
                      {t('socialRelations')}
                    </h2>
                    <p className="text-sm text-muted-foreground leading-relaxed">
                      {detail.social_relations}
                    </p>
                  </section>
                )}

                {/* Culture Tab */}
                <section>
                  {detail.arts && (
                    <section>
                      <h2 className="text-lg font-bold">{t('arts')}</h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.arts}
                      </p>
                    </section>
                  )}

                  {detail.worship && (
                    <section>
                      <h2 className="text-lg font-bold">{t('worship')}</h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.worship}
                      </p>
                    </section>
                  )}

                  {detail.entertainment && (
                    <section>
                      <h2 className="text-lg font-bold">
                        {t('entertainment')}
                      </h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.entertainment}
                      </p>
                    </section>
                  )}

                  {detail.marriage && (
                    <section>
                      <h2 className="text-lg font-bold">{t('marriage')}</h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.marriage}
                      </p>
                    </section>
                  )}

                  {detail.funeral && (
                    <section>
                      <h2 className="text-lg font-bold">{t('funeral')}</h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.funeral}
                      </p>
                    </section>
                  )}
                </section>

                {/* Lifestyle Tab */}
                <section>
                  {detail.food && (
                    <section>
                      <h2 className="text-lg flex items-center gap-2 font-bold">
                        <UtensilsCrossed className="h-5 w-5 text-primary" />
                        {t('food')}
                      </h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.food}
                      </p>
                    </section>
                  )}

                  {detail.clothing && (
                    <section>
                      <h2 className="text-lg flex items-center gap-2 font-bold">
                        <Shirt className="h-5 w-5 text-primary" />
                        {t('clothing')}
                      </h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.clothing}
                      </p>
                    </section>
                  )}

                  {detail.housing && (
                    <section>
                      <h2 className="text-lg flex items-center gap-2 font-bold">
                        <Home className="h-5 w-5 text-primary" />
                        {t('housing')}
                      </h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.housing}
                      </p>
                    </section>
                  )}

                  {detail.economic_activities && (
                    <section>
                      <h2 className="text-lg flex items-center gap-2 font-bold">
                        <Briefcase className="h-5 w-5 text-primary" />
                        {t('economicActivities')}
                      </h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.economic_activities}
                      </p>
                    </section>
                  )}

                  {detail.transportation && detail.transportation.trim() && (
                    <section>
                      <h2 className="text-lg font-bold">
                        {t('transportation')}
                      </h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.transportation}
                      </p>
                    </section>
                  )}

                  {detail.childbirth && detail.childbirth.trim() && (
                    <section>
                      <h2 className="text-lg font-bold">{t('childbirth')}</h2>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.childbirth}
                      </p>
                    </section>
                  )}
                </section>

                {/* Images */}
                {detail.images.length > 0 && (
                  <section>
                    <h2 className="text-lg font-bold">{t('images')}</h2>
                    <div className="grid grid-cols-3 gap-4">
                      {detail.images.map((img) => (
                        <Image
                          alt={img.alt}
                          src={img.url}
                          width={200}
                          height={200}
                          className="object-contain h-full w-full aspect-square"
                        />
                      ))}
                    </div>
                  </section>
                )}
              </article>
            </CardContent>
          </Card>
        </div>
      </div>

      {/* Bottom Navigation */}
      <div className="border-t border-border/40 bg-card/30 backdrop-blur-sm">
        <div className="container mx-auto px-4 py-8">
          <Link
            href="/dan-toc"
            className="inline-flex items-center text-sm font-medium text-primary hover:underline"
          >
            <ArrowLeft className="h-4 w-4 mr-2" />
            {t('backToList')}
          </Link>
        </div>
      </div>
    </div>
  );
}
