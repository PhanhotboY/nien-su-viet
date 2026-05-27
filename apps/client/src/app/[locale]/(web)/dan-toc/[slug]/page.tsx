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

import { getMetadata } from '@/content/landing/metadata';
import { genMetadata } from '@/lib/metadata.lib';
import { ethnicGroupsMap } from '@/data/vi/vietnamese-ethnic-groups';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';

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
                  <p className="text-sm font-medium text-foreground">
                    {detail.self_name}
                  </p>
                </CardContent>
              </Card>
            )}
          </div>

          {/* Right Column - Detailed Tabs */}
          <div className="lg:col-span-2">
            <Tabs defaultValue="overview" className="w-full">
              <TabsList className="grid w-full grid-cols-3 mb-6">
                <TabsTrigger value="overview">{t('overview')}</TabsTrigger>
                <TabsTrigger value="culture">{t('culture')}</TabsTrigger>
                <TabsTrigger value="lifestyle">{t('lifestyle')}</TabsTrigger>
              </TabsList>

              {/* Overview Tab */}
              <TabsContent value="overview" className="space-y-6">
                {detail.other_names && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">
                        {t('otherNames')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.other_names}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.local_groups && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">
                        {t('localGroups')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.local_groups}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.history && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg flex items-center gap-2">
                        <Calendar className="h-5 w-5 text-primary" />
                        {t('history')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.history}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.social_relations && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">
                        {t('socialRelations')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.social_relations}
                      </p>
                    </CardContent>
                  </Card>
                )}
              </TabsContent>

              {/* Culture Tab */}
              <TabsContent value="culture" className="space-y-6">
                {detail.arts && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">{t('arts')}</CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.arts}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.worship && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">{t('worship')}</CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.worship}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.entertainment && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">
                        {t('entertainment')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.entertainment}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.marriage && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">{t('marriage')}</CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.marriage}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.funeral && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">{t('funeral')}</CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.funeral}
                      </p>
                    </CardContent>
                  </Card>
                )}
              </TabsContent>

              {/* Lifestyle Tab */}
              <TabsContent value="lifestyle" className="space-y-6">
                {detail.food && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg flex items-center gap-2">
                        <UtensilsCrossed className="h-5 w-5 text-primary" />
                        {t('food')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.food}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.clothing && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg flex items-center gap-2">
                        <Shirt className="h-5 w-5 text-primary" />
                        {t('clothing')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.clothing}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.housing && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg flex items-center gap-2">
                        <Home className="h-5 w-5 text-primary" />
                        {t('housing')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.housing}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.economic_activities && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg flex items-center gap-2">
                        <Briefcase className="h-5 w-5 text-primary" />
                        {t('economicActivities')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.economic_activities}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.transportation && detail.transportation.trim() && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">
                        {t('transportation')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.transportation}
                      </p>
                    </CardContent>
                  </Card>
                )}

                {detail.childbirth && detail.childbirth.trim() && (
                  <Card>
                    <CardHeader>
                      <CardTitle className="text-lg">
                        {t('childbirth')}
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <p className="text-sm text-muted-foreground leading-relaxed">
                        {detail.childbirth}
                      </p>
                    </CardContent>
                  </Card>
                )}
              </TabsContent>
            </Tabs>
          </div>
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
