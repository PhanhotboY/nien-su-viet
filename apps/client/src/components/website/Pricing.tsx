import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Check } from 'lucide-react';
import { getTranslations } from 'next-intl/server';

enum PopularPlanType {
  NO = 0,
  YES = 1,
}

interface PricingProps {
  title: string;
  popular: PopularPlanType;
  price: number;
  description: string;
  buttonText: string;
  benefitList: string[];
}

export const Pricing = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({ namespace: 'AboutPage.Pricing', locale });

  const pricingList: PricingProps[] = [
    {
      title: t('plans.free.title'),
      popular: PopularPlanType.NO,
      price: 0,
      description: t('plans.free.description'),
      buttonText: t('plans.free.buttonText'),
      benefitList: [
        t('plans.free.benefits.timeline'),
        t('plans.free.benefits.search100'),
        t('plans.free.benefits.biographies'),
        t('plans.free.benefits.communitySupport'),
        t('plans.free.benefits.multiDevice'),
      ],
    },
    {
      title: t('plans.premium.title'),
      popular: PopularPlanType.YES,
      price: 0,
      description: t('plans.premium.description'),
      buttonText: t('plans.premium.buttonText'),
      benefitList: [
        t('plans.premium.benefits.fullAccess'),
        t('plans.premium.benefits.downloads'),
        t('plans.premium.benefits.interactiveMap'),
        t('plans.premium.benefits.noAds'),
        t('plans.premium.benefits.prioritySupport'),
      ],
    },
    {
      title: t('plans.organization.title'),
      popular: PopularPlanType.NO,
      price: 0,
      description: t('plans.organization.description'),
      buttonText: t('plans.organization.buttonText'),
      benefitList: [
        t('plans.organization.benefits.unlimitedAccounts'),
        t('plans.organization.benefits.centralManagement'),
        t('plans.organization.benefits.customContent'),
        t('plans.organization.benefits.apiIntegration'),
        t('plans.organization.benefits.trainingSupport'),
      ],
    },
  ];

  return (
    <section id="pricing" className="container py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold text-center">
        {t('title.prefix')}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {' '}
          {t('title.highlight')}{' '}
        </span>
        {t('title.suffix')}
      </h2>
      <h3 className="text-xl text-center text-muted-foreground pt-4 pb-8">
        {t('subtitle')}
      </h3>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {pricingList.map((pricing: PricingProps) => (
          <Card
            key={pricing.title}
            className={
              pricing.popular === PopularPlanType.YES
                ? 'drop-shadow-primary/25 drop-shadow-xl shadow-primary/10 dark:shadow-white/10'
                : ''
            }
          >
            <CardHeader>
              <CardTitle className="flex item-center justify-between">
                {pricing.title}
                {pricing.popular === PopularPlanType.YES ? (
                  <Badge variant="secondary">{t('popularBadge')}</Badge>
                ) : null}
              </CardTitle>

              <div>
                <span className="text-3xl font-bold">
                  {pricing.price === 0
                    ? t('price.free')
                    : t('price.amount', { amount: `${pricing.price}đ` })}
                </span>
                {pricing.price > 0 && (
                  <span className="text-muted-foreground">
                    {' '}
                    {t('price.perMonth')}
                  </span>
                )}
              </div>

              <CardDescription>{pricing.description}</CardDescription>
            </CardHeader>

            <CardContent>
              <Button className="w-full">{pricing.buttonText}</Button>
            </CardContent>

            <hr className="w-4/5 m-auto mb-4" />

            <CardFooter className="flex">
              <div className="space-y-4">
                {pricing.benefitList.map((benefit: string) => (
                  <span key={benefit} className="flex">
                    <Check className="text-green-500" />{' '}
                    <h3 className="ml-2">{benefit}</h3>
                  </span>
                ))}
              </div>
            </CardFooter>
          </Card>
        ))}
      </div>
    </section>
  );
};
