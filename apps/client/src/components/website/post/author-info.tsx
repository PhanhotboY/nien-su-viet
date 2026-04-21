'use client';

import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { cn, getAvatarFallback } from '@/lib/utils';
import { getUserById } from '@/services/user.service';
import { useTranslations } from 'next-intl';
import { useEffect, useState } from 'react';

interface AuthorInfoProps {
  authorId: string;
  className?: string;
  showAuthorLabel?: boolean;
}

export default function AuthorInfo({
  authorId,
  className,
  showAuthorLabel,
}: AuthorInfoProps) {
  const tshared = useTranslations('Shared');
  const [author, setAuthor] = useState<
    Awaited<ReturnType<typeof getUserById>>['data']
  >({
    id: '',
    name: tshared('unknown-author'),
  });

  useEffect(() => {
    let isMounted = true;

    getUserById(authorId)
      .then(({ data }) => {
        if (isMounted && data) {
          setAuthor(data);
        }
      })
      .catch((e) => {
        console.error('Error fetching author:', e);
      });

    return () => {
      isMounted = false;
    };
  }, [authorId, tshared]);

  return (
    <>
      <Avatar className={cn('shadow-sm/30', className)}>
        <AvatarImage src={author.image || '/assets/logo/nsv-logo-96.webp'} />
        <AvatarFallback>{getAvatarFallback(author.name)}</AvatarFallback>
      </Avatar>
      <div className="ml-4 flex flex-col">
        <span className="text-md flex font-semibold">{author.name}</span>
        {showAuthorLabel && <p className="text-xs">{tshared('author')}</p>}
      </div>
    </>
  );
}
