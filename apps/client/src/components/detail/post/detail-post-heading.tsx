import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { getMinutes, shimmer, toBase64 } from '@/lib/utils';
import { ArchiveIcon, CalendarIcon, ClockIcon } from 'lucide-react';
import { getTranslations } from 'next-intl/server';
import Image from 'next/image';
import { FC } from 'react';
import { ReadTimeResults } from 'reading-time';

interface DetailPostHeadingProps {
  id: string;
  title: string;
  thumbnail: string;
  authorImage?: string;
  authorName: string;
  date: string;
  // category: string;
  readTime: ReadTimeResults;
}

const DetailPostHeading: FC<DetailPostHeadingProps> = async ({
  id,
  title,
  thumbnail,
  authorName,
  authorImage,
  date,
  // category,
  readTime,
}) => {
  const tshared = await getTranslations('Shared');
  return (
    <section className="flex flex-col items-start justify-between">
      <div className="relative w-full">
        <Image
          src={thumbnail}
          alt={title}
          width={512}
          height={288}
          className="h-[288px] w-full rounded-2xl bg-gray-100 object-cover"
          placeholder={`data:image/svg+xml;base64,${toBase64(
            shimmer(512, 288),
          )}`}
        />
        <div className="absolute inset-0 rounded-2xl ring-1 ring-inset ring-gray-900/10" />
      </div>
      <div className="w-full">
        <p className="my-5 overflow-hidden text-xl md:text-2xl font-semibold leading-6 text-card-foreground">
          {title}
        </p>

        {/* Mobile view */}
        <div className="mb-5 grid grid-cols-2 gap-2 rounded-md border border-gray-100 px-3 py-2.5 sm:hidden">
          {/* Author */}
          <div className="inline-flex items-start justify-start">
            <Avatar>
              <AvatarImage src={authorImage} />
              <AvatarFallback>P</AvatarFallback>
            </Avatar>
            <div className="ml-2 flex flex-col">
              <span className="text-md flex font-semibold">{authorName}</span>
            </div>
          </div>

          {/* Date */}
          <div className="inline-flex space-x-2 border-gray-400 border-opacity-50">
            <p className="mt-0.5 opacity-80">
              <span className="sr-only">{tshared('date')}</span>
              <CalendarIcon className="h-4 w-4" aria-hidden="true" />
            </p>
            <span className="text-sm">{date}</span>
          </div>
          {/* Category */}
          <div className="inline-flex space-x-2 border-gray-400 border-opacity-50">
            <p className="mt-0.5 opacity-80">
              <span className="sr-only">{tshared('category')}</span>
              <ArchiveIcon className="h-4 w-4" aria-hidden="true" />
            </p>
            {/*<span className="text-sm">{category}</span>*/}
          </div>

          {/* Reading time */}
          <div className="inline-flex space-x-2 border-gray-400 border-opacity-50">
            <p className="mt-0.5 opacity-80">
              <span className="sr-only">{tshared('minute-to-read')}</span>
              <ClockIcon className="h-4 w-4" aria-hidden="true" />
            </p>
            <span className="text-sm">{getMinutes(readTime.minutes)}</span>
          </div>
        </div>

        {/* Desktop view */}
        <div className="mb-7 hidden justify-start text-card-foreground sm:flex sm:flex-row">
          <Avatar>
            <AvatarImage src={authorImage} />
            <AvatarFallback>P</AvatarFallback>
          </Avatar>
          {/* Author */}
          <div className="mb-5 flex flex-row items-center justify-start pr-3.5 md:mb-0">
            <div className="ml-2 flex flex-col">
              <span className="text-md flex font-semibold">{authorName}</span>
            </div>
          </div>
          <div className="flex flex-row items-center">
            {/* Date */}
            <div className="flex space-x-2 border-gray-400 border-opacity-50 pl-0 pr-3.5 md:border-l md:pl-3.5">
              <p className="mt-0.5 opacity-80">
                <span className="sr-only">{tshared('date')}</span>
                <CalendarIcon className="h-4 w-4" aria-hidden="true" />
              </p>
              <span className="text-sm">{date}</span>
            </div>
            {/* Category */}
            <div className="flex space-x-2 border-l border-gray-400 border-opacity-50 pl-3.5 pr-3.5">
              <p className="mt-0.5 opacity-80">
                <span className="sr-only">{tshared('category')}</span>
                <ArchiveIcon className="h-4 w-4" aria-hidden="true" />
              </p>
              {/*<span className="text-sm">{category}</span>*/}
            </div>
            {/* Reading time */}
            <div className="flex space-x-2 border-l border-gray-400 border-opacity-50 pl-3.5">
              <p className="mt-0.5 opacity-80">
                <span className="sr-only">{tshared('minute-to-read')}</span>
                <ClockIcon className="h-4 w-4" aria-hidden="true" />
              </p>
              <span className="text-sm">
                {getMinutes(readTime.minutes, tshared('minute'))}
              </span>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default DetailPostHeading;
