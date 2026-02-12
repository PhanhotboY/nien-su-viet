'use client';

import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { protectedPostConfig } from '@/config/cmsdesk';
import { createPost } from '@/services/post.service';
import { useAuthenticate } from '@daveyplate/better-auth-ui';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { Loader2 as SpinnerIcon } from 'lucide-react';
import { useRouter } from 'next/navigation';
import React, { useState } from 'react';
import { toast } from 'sonner';

const PostCreateButton = () => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const { data: session } = useAuthenticate();
  const router = useRouter();

  async function createPostHandler() {
    setIsLoading(true);

    if (session?.user.id) {
      const post: components['schemas']['PostBaseCreateDto'] = {
        title: protectedPostConfig.untitled,
        authorId: session?.user.id,
        slug: `untitled-${Math.random().toString(36).substring(2, 8)}`,
        content: '{blocks: []}',
        published: false,
      };

      const response = await createPost(post);

      if (response.data?.success) {
        toast.success(protectedPostConfig.successCreate);
        // This forces a cache invalidation.
        router.refresh();
        // Redirect to the new post
        router.push('/cmsdesk/posts/' + response.data.id);
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
      <button
        type="button"
        onClick={createPostHandler}
        className="flex items-center rounded-md bg-gray-900 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-gray-600"
      >
        {isLoading && <SpinnerIcon className="mr-2 h-4 w-4 animate-spin" />}
        {protectedPostConfig.newPost}
      </button>
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
