'use client';

import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { Button } from '@/components/ui/button';
import { protectedPostConfig } from '@/config/cmsdesk';
import { createPost } from '@/services/post.service';
import { useAuthenticate } from '@daveyplate/better-auth-ui';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { Loader2 as SpinnerIcon } from 'lucide-react';
import { useLocale } from 'next-intl';
import { useRouter } from 'next/navigation';
import React, { useState } from 'react';
import { toast } from 'sonner';

const PostCreateButton = () => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const { data: session } = useAuthenticate();
  const router = useRouter();
  const locale = useLocale();

  async function createPostHandler() {
    setIsLoading(true);

    if (session?.user.id) {
      const post: components['schemas']['PostBaseCreateDto'] = {
        title: protectedPostConfig.untitled,
        slug: `untitled-${Math.random().toString(36).substring(2, 8)}`,
        content: '{"blocks": []}',
        published: false,
      };

      const response = await createPost(post);

      if (response.data?.success) {
        toast.success(protectedPostConfig.successCreate);
        // This forces a cache invalidation.
        router.refresh();
        // Redirect to the new post
        router.push(`/${locale}/cmsdesk/posts/` + response.data.id);
        setIsLoading(false);
      } else {
        setIsLoading(false);
        toast.error(response.message || protectedPostConfig.errorCreate);
      }
    } else {
      setIsLoading(false);
      toast.error(protectedPostConfig.errorCreate);
    }
  }

  return (
    <>
      <Button
        type="button"
        onClick={createPostHandler}
        className="flex items-center rounded-md px-3.5 py-2.5 shadow-sm"
      >
        {isLoading && <SpinnerIcon className="mr-2 h-4 w-4 animate-spin" />}
        {protectedPostConfig.newPost}
      </Button>
      <AlertDialog open={isLoading} onOpenChange={setIsLoading}>
        <AlertDialogContent className="font-sans">
          <AlertDialogHeader>
            <AlertDialogTitle className="text-center">
              {protectedPostConfig.pleaseWait}
            </AlertDialogTitle>
            <AlertDialogDescription className="mx-auto text-center">
              <SpinnerIcon className="h-6 w-6 animate-spin" />
            </AlertDialogDescription>
          </AlertDialogHeader>
        </AlertDialogContent>
      </AlertDialog>
    </>
  );
};

export default PostCreateButton;
