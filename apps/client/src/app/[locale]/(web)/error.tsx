'use client';

import { SharedError } from '@/components/shared';
import { Providers } from '../provider';

export default function Error({
  error,
}: {
  error: Error & { digest?: string };
}) {
  return (
    <Providers>
      <SharedError error={error} />
    </Providers>
  );
}
