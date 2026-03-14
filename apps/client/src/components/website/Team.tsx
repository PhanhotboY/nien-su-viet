import { buttonVariants } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Facebook, Instagram, Linkedin } from 'lucide-react';
import Link from '@/i18n/navigation';
import { getTranslations } from 'next-intl/server';

interface TeamProps {
  imageUrl: string;
  name: string;
  position: string;
  socialNetworks: SociaNetworkslProps[];
}

interface SociaNetworkslProps {
  name: string;
  url: string;
}

const teamList: TeamProps[] = [
  {
    imageUrl: 'https://i.pravatar.cc/150?img=35',
    name: 'TS. Nguyễn Văn Minh',
    position: 'Chuyên gia Lịch sử',
    socialNetworks: [
      {
        name: 'Linkedin',
        url: 'https://www.linkedin.com/',
      },
      {
        name: 'Facebook',
        url: 'https://www.facebook.com/',
      },
      {
        name: 'Instagram',
        url: 'https://www.instagram.com/',
      },
    ],
  },
  {
    imageUrl: 'https://i.pravatar.cc/150?img=60',
    name: 'PGS. Trần Thị Hương',
    position: 'Giám đốc Nội dung',
    socialNetworks: [
      {
        name: 'Linkedin',
        url: 'https://www.linkedin.com/',
      },
      {
        name: 'Facebook',
        url: 'https://www.facebook.com/',
      },
      {
        name: 'Instagram',
        url: 'https://www.instagram.com/',
      },
    ],
  },
  {
    imageUrl: 'https://i.pravatar.cc/150?img=36',
    name: 'Lê Quang Đức',
    position: 'Trưởng phòng Số hóa',
    socialNetworks: [
      {
        name: 'Linkedin',
        url: 'https://www.linkedin.com/',
      },

      {
        name: 'Instagram',
        url: 'https://www.instagram.com/',
      },
    ],
  },
  {
    imageUrl: 'https://i.pravatar.cc/150?img=17',
    name: 'Phạm Minh Tuấn',
    position: 'Kỹ sư Phát triển',
    socialNetworks: [
      {
        name: 'Linkedin',
        url: 'https://www.linkedin.com/',
      },
      {
        name: 'Facebook',
        url: 'https://www.facebook.com/',
      },
    ],
  },
];

export const Team = async ({ locale }: { locale: string }) => {
  const t = await getTranslations({ namespace: 'AboutPage.Team', locale });

  const socialIcon = (iconName: string) => {
    switch (iconName) {
      case 'Linkedin':
        return <Linkedin size="20" />;

      case 'Facebook':
        return <Facebook size="20" />;

      case 'Instagram':
        return <Instagram size="20" />;
    }
  };

  return (
    <section id="team" className="container py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold">
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {t('title.prefix')}{' '}
        </span>
        {t('title.suffix')}
      </h2>

      <p className="mt-4 mb-10 text-xl text-muted-foreground">
        {t('description')}
      </p>

      <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8 gap-y-10">
        {teamList.map(
          ({ imageUrl, name, position, socialNetworks }: TeamProps) => (
            <Card
              key={name}
              className="bg-muted/50 relative mt-8 flex flex-col justify-center items-center gap-0"
            >
              <CardHeader className="w-full mt-8 flex justify-center flex-col items-center px-0 pb-2">
                <img
                  src={imageUrl}
                  alt={`${name} ${position}`}
                  className="absolute -top-12 rounded-full w-24 h-24 aspect-square object-cover"
                />
                <CardTitle className="text-center">{name}</CardTitle>
                <CardDescription className="text-secondary">
                  {position}
                </CardDescription>
              </CardHeader>

              <CardContent className="text-center pb-2">
                <p>{t('memberTagline')}</p>
              </CardContent>

              <CardFooter>
                {socialNetworks.map(({ name, url }: SociaNetworkslProps) => (
                  <div key={name}>
                    <Link
                      rel="noreferrer noopener"
                      href={url}
                      target="_blank"
                      className={buttonVariants({
                        variant: 'ghost',
                        size: 'sm',
                      })}
                    >
                      <span className="sr-only">
                        {t('social.srOnly', { name })}
                      </span>
                      {socialIcon(name)}
                    </Link>
                  </div>
                ))}
              </CardFooter>
            </Card>
          ),
        )}
      </div>
    </section>
  );
};
