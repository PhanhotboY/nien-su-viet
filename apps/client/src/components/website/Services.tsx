import { Card, CardDescription, CardHeader, CardTitle } from '../ui/card';
import { MagnifierIcon, WalletIcon, ChartIcon } from '../Icons';
import { JSX } from 'react';
import Image from 'next/image';
import { getTranslations } from 'next-intl/server';

interface ServiceProps {
  title: string;
  description: string;
  icon: JSX.Element;
}

export const Services = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({ namespace: 'Services', locale });

  const serviceList: ServiceProps[] = [
    {
      title: t('items.eventLookup.title'),
      description: t('items.eventLookup.description'),
      icon: <ChartIcon />,
    },
    {
      title: t('items.figureLibrary.title'),
      description: t('items.figureLibrary.description'),
      icon: <WalletIcon />,
    },
    {
      title: t('items.historicalDocuments.title'),
      description: t('items.historicalDocuments.description'),
      icon: <MagnifierIcon />,
    },
  ];

  return (
    <section className="container py-24 sm:py-32">
      <div className="grid lg:grid-cols-2 gap-8 place-items-center">
        <div>
          <h2 className="text-3xl md:text-4xl font-bold">
            <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
              {t('heading.emphasis')}{' '}
            </span>
            {t('heading.rest')}
          </h2>

          <p className="text-muted-foreground text-xl mt-4 mb-8 ">
            {t('description')}
          </p>

          <div className="flex flex-col gap-8">
            {serviceList.map(({ icon, title, description }: ServiceProps) => (
              <Card key={title}>
                <CardHeader className="space-y-1 flex md:flex-row justify-start items-start gap-4">
                  <div className="mt-1 bg-primary/20 p-1 rounded-2xl">
                    {icon}
                  </div>
                  <div>
                    <CardTitle>{title}</CardTitle>
                    <CardDescription className="text-md mt-2">
                      {description}
                    </CardDescription>
                  </div>
                </CardHeader>
              </Card>
            ))}
          </div>
        </div>

        {/*<Image
          src={cubeLeg}
          className="w-[300px] md:w-[500px] lg:w-[600px] object-contain"
          alt="About services"
        />*/}
      </div>
    </section>
  );
};
