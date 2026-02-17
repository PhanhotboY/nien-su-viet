import { Metadata } from 'next';
import { CLIENT_HOST } from './config';

interface GenMetadataParams {
  title: string;
  description: string;
  locale: string;
  logo?: string | null;
  path?: string;
}
function genMetadata({
  title,
  description,
  locale,
  logo,
  path = '/',
}: GenMetadataParams): Metadata {
  return {
    title,
    description,
    openGraph: {
      title,
      description,
      type: 'website',
      locale,
      siteName: title,
      images: logo ? [logo] : undefined,
    },
    twitter: {
      card: 'summary_large_image',
      title,
      description,
      images: logo ? [logo] : undefined,
    },
    alternates: {
      canonical: `${CLIENT_HOST}/${locale}${path === '/' ? '' : path}`,
    },
    robots: {
      index: true,
      follow: true,
      nocache: false,
      googleBot: {
        index: true,
        follow: true,
        noimageindex: false,
        'max-video-preview': -1,
        'max-image-preview': 'large',
        'max-snippet': -1,
      },
    },
  };
}

export { genMetadata };
