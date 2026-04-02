import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { cn, getAvatarFallback } from '@/lib/utils';
import { getUserById } from '@/services/user.service';
import { getTranslations } from 'next-intl/server';

interface AuthorInfoProps {
  authorId: string;
  className?: string;
  showAuthorLabel?: boolean;
}

export default async function AuthorInfo({
  authorId,
  className,
  showAuthorLabel,
}: AuthorInfoProps) {
  const tshared = await getTranslations('Shared');
  let { data: author } = await getUserById(authorId).catch((e) => {
    console.error('Error fetching author:', e);
    return { data: null };
  });
  if (!author) {
    author = {
      id: '',
      name: tshared('unknown-author'),
    };
  }

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
