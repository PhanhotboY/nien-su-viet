'use client';

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { protectedPostConfig } from '@/config/cmsdesk';
import { deletePost } from '@/services/post.service';
import { useAuthenticate, useAuth } from '@/components/auth';
import {
  MoreVertical as ElipsisIcon,
  Loader2 as SpinnerIcon,
  Trash as TrashIcon,
} from 'lucide-react';
import { useLocale } from 'next-intl';
import { useRouter } from 'next/navigation';
import React, { FC, useState } from 'react';
import { toast } from 'sonner';
import Link from '@/i18n/navigation';

interface PostEditButtonProps {
  id: string;
  published?: boolean;
}

const PostEditButton: FC<PostEditButtonProps> = ({ id, published = false }) => {
  const router = useRouter();
  const [showDeleteAlert, setShowDeleteAlert] = useState<boolean>(false);
  const [isDeleteLoading, setIsDeleteLoading] = useState<boolean>(false);
  const [showLoadingAlert, setShowLoadingAlert] = useState<boolean>(false);
  const { authClient } = useAuth();
  const { data: session } = useAuthenticate(authClient);
  const locale = useLocale();

  // Delete post
  async function deleteMyPost() {
    setIsDeleteLoading(true);
    if (id && session?.user.id) {
      const myPostData = {
        id: id,
        userId: session?.user.id,
      };
      const response = await deletePost(myPostData.id);

      if (response.data?.success) {
        setIsDeleteLoading(false);
        toast.success(protectedPostConfig.successDelete);
        router.refresh();
      } else {
        setIsDeleteLoading(false);
        toast.error(protectedPostConfig.errorDelete + response.message);
      }
    } else {
      setIsDeleteLoading(false);
      toast.error(protectedPostConfig.errorDelete);
    }
  }

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger className="flex h-8 w-8 items-center justify-center rounded-md border transition-colors hover:bg-muted">
          <ElipsisIcon className="h-4 w-4" />
          <span className="sr-only">Open</span>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" className="">
          <DropdownMenuItem>
            <a
              className="flex w-full"
              onClick={() => {
                setShowLoadingAlert(true);
                router.push(`/${locale}/cmsdesk/posts/${id}`);
                setShowLoadingAlert(false);
              }}
            >
              {protectedPostConfig.edit}
            </a>
          </DropdownMenuItem>
          {published && (
            <>
              <DropdownMenuSeparator />
              <DropdownMenuItem>
                <Link
                  className="flex w-full"
                  href={`/posts/${id}`}
                  target="_blank"
                >
                  {protectedPostConfig.viewPublished}
                </Link>
              </DropdownMenuItem>
            </>
          )}
          <DropdownMenuSeparator />
          <DropdownMenuItem
            className="flex cursor-pointer items-center text-destructive focus:text-destructive"
            onSelect={() => setShowDeleteAlert(true)}
          >
            {protectedPostConfig.delete}
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      {/* Delete alert */}
      <AlertDialog open={showDeleteAlert} onOpenChange={setShowDeleteAlert}>
        <AlertDialogContent className="text-md">
          <AlertDialogHeader>
            <AlertDialogTitle>
              {protectedPostConfig.questionDelete}
            </AlertDialogTitle>
            <AlertDialogDescription>
              {protectedPostConfig.warning}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>{protectedPostConfig.cancel}</AlertDialogCancel>
            <AlertDialogAction onClick={deleteMyPost}>
              {isDeleteLoading ? (
                <SpinnerIcon className="mr-2 h-4 w-4 animate-spin" />
              ) : (
                <TrashIcon className="mr-2 h-4 w-4" />
              )}
              <span>{protectedPostConfig.confirm}</span>
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
      {/* Loading alert */}
      <AlertDialog open={showLoadingAlert} onOpenChange={setShowLoadingAlert}>
        <AlertDialogContent className="">
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

export default PostEditButton;
