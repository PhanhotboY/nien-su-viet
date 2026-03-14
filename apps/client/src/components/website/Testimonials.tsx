import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { getTranslations } from 'next-intl/server';

interface TestimonialProps {
  image: string;
  name: string;
  userName: string;
  comment: string;
}

export const Testimonials = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({
    namespace: 'AboutPage.Testimonials',
    locale,
  });

  let testimonials: TestimonialProps[] = [];

  try {
    const rawItems = t.raw('items');
    if (Array.isArray(rawItems)) {
      testimonials = rawItems as TestimonialProps[];
    } else if (typeof rawItems === 'object' && rawItems !== null) {
      // If it's an object but not an array, it might be keyed testimonials
      const itemsArray = Object.values(rawItems);
      if (Array.isArray(itemsArray)) {
        testimonials = itemsArray as TestimonialProps[];
      }
    }
  } catch (error) {
    console.error('Error loading testimonials:', error);
  }

  return (
    <section id="testimonials" className="container py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold">
        {t('title.prefix')}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {' '}
          {t('title.highlight')}{' '}
        </span>
        {t('title.suffix')}
      </h2>

      <p className="text-xl text-muted-foreground pt-4 pb-8">{t('subtitle')}</p>

      <div className="grid md:grid-cols-2 lg:grid-cols-4 sm:block columns-2 lg:columns-3 lg:gap-6 mx-auto space-y-4 lg:space-y-6">
        {testimonials.map(
          ({ image, name, userName, comment }: TestimonialProps) => (
            <Card
              key={userName}
              className="md:break-inside-avoid overflow-hidden gap-2"
            >
              <CardHeader className="flex flex-row items-center gap-4">
                <Avatar className="size-10">
                  <AvatarImage alt="" src={image} />
                  <AvatarFallback>OM</AvatarFallback>
                </Avatar>

                <div className="flex flex-col">
                  <CardTitle className="text-lg">{name}</CardTitle>
                  <CardDescription>{userName}</CardDescription>
                </div>
              </CardHeader>

              <CardContent>{comment}</CardContent>
            </Card>
          ),
        )}
      </div>
    </section>
  );
};
