'use client';

import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import FileUploadModal from '@/components/ui/file-upload-modal';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
import { Separator } from '@/components/ui/separator';
import { Textarea } from '@/components/ui/textarea';
import { mainCategoryConfig } from '@/config/main';
import { protectedEditorConfig, protectedPostConfig } from '@/config/cmsdesk';
import { postEditFormSchema } from '@/lib/validation/post';
import { PaperClipIcon } from '@heroicons/react/20/solid';
import { zodResolver } from '@hookform/resolvers/zod';
import { SparklesIcon, Loader2 as SpinnerIcon } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { FC, useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'sonner';
import slugify from 'react-slugify';
import { v4 } from 'uuid';
import * as z from 'zod';
import {
  EditorUploadCoverImageItem,
  EditorUploadCoverImagePlaceHolder,
} from './upload';
import { PostResponseDto } from '@/types/collection';
import { updatePost } from '@/services/post.service';
import { useAuthenticate } from '@daveyplate/better-auth-ui';
import TextEditor from '@/components/TextEditor';

export const dynamic = 'force-dynamic';

type FormData = z.infer<typeof postEditFormSchema>;

interface EditorProps {
  post: PostResponseDto;
  imageFolderName: string;
}

type EditorFormValues = z.infer<typeof postEditFormSchema>;

const Editor: FC<EditorProps> = ({ post, imageFolderName }) => {
  const router = useRouter();
  const { data } = useAuthenticate();
  const userId = data?.user.id || '';

  // These are the values that will be used to upload the image
  // States
  const [isSaving, setIsSaving] = useState(false);
  const [showLoadingAlert, setShowLoadingAlert] = useState<boolean>(false);
  const [showCoverModal, setShowCoverModal] = useState<boolean>(false);

  // Editor
  const [saveStatus, setSaveStatus] = useState('Saved');

  const [content, setContent] = useState<string | null>(post?.content || null);

  // Handle cover image upload complete
  const handleCoverUploadComplete = (urls: string[]) => {
    form.setValue('thumbnail', urls[0]);
    toast.success(protectedEditorConfig.successMessageImageUpload);
    router.refresh();
  };
  // Handle upload error
  const handleUploadError = (error: string) => {
    toast.error(error || protectedEditorConfig.errorMessageImageUpload);
  };

  // Default values for the form
  const defaultValues: Partial<EditorFormValues> = {
    title: post.title ?? 'Untitled',
    slug: post.slug ?? `post-${v4()}`,
    thumbnail: post.thumbnail ?? undefined,
    summary: post.summary ?? 'Post summary',
    content: content ?? protectedEditorConfig.placeholderContent,
    published: post.published === true,
  };

  const form = useForm<EditorFormValues>({
    resolver: zodResolver(postEditFormSchema),
    defaultValues,
    mode: 'onChange',
  });

  async function onSubmit(data: EditorFormValues) {
    setShowLoadingAlert(true);
    setIsSaving(true);

    const response = await updatePost(post.id, {
      title: data.title,
      slug: data.slug,
      thumbnail: data.thumbnail,
      summary: data.summary,
      content: content!,
      categoryId: data.categoryId,
      published: data.published,
    });

    if (response.statusCode < 400) {
      toast.success(protectedEditorConfig.successMessage);
      router.push(`/cmsdesk/posts?search=refresh`);
    } else {
      toast.error(protectedEditorConfig.errorMessage + ': ' + response.message);
    }

    setIsSaving(false);
    setShowLoadingAlert(false);
  }

  return (
    <>
      <Form {...form}>
        {/* Title */}
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="grid grid-cols-12 gap-8"
        >
          {/* General information */}
          <Card className="col-span-9">
            <CardHeader>
              <CardTitle>{protectedEditorConfig.generalTitle}</CardTitle>
              <CardDescription>
                {protectedEditorConfig.generalDescription}
              </CardDescription>
            </CardHeader>
            <Separator className="" />
            <CardContent className="space-y-4">
              <FormField
                control={form.control}
                name="title"
                render={({ field }) => (
                  <FormItem className="w-full">
                    <FormLabel>{protectedEditorConfig.formTitle}</FormLabel>
                    <FormControl>
                      <Input
                        placeholder={protectedEditorConfig.placeHolderTitle}
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Slug */}
              <FormField
                control={form.control}
                name="slug"
                render={({ field }) => (
                  <FormItem className="w-full">
                    <FormLabel>{protectedEditorConfig.slug}</FormLabel>
                    <FormControl>
                      <Input
                        placeholder={protectedEditorConfig.placeholderSlug}
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      <Button
                        type="button"
                        variant="outline"
                        size="sm"
                        className="mt-2"
                        onClick={() =>
                          field.onChange(slugify(form.getValues('title')))
                        }
                      >
                        <SparklesIcon className="mr-2 h-4 w-4" />
                        {protectedEditorConfig.generateSlug}
                      </Button>
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </CardContent>
          </Card>

          {/* Category */}
          <Card className="col-span-3">
            <CardHeader>
              <CardTitle>{protectedEditorConfig.categoryTitle}</CardTitle>
              <CardDescription>
                {protectedEditorConfig.categoryDescription}
              </CardDescription>
            </CardHeader>
            <Separator className="" />
            <CardContent className="space-y-4">
              <FormField
                control={form.control}
                name="categoryId"
                render={({ field }) => (
                  <FormItem className="space-y-3">
                    <FormControl>
                      <RadioGroup
                        onValueChange={field.onChange}
                        defaultValue={field.value}
                        className="flex flex-col space-y-1"
                      >
                        {mainCategoryConfig.map(
                          (category) =>
                            category.slug !== '/' && (
                              <FormItem
                                key={category.id}
                                className="flex items-center space-x-3 space-y-0"
                              >
                                <FormControl>
                                  <RadioGroupItem value={category.id} />
                                </FormControl>
                                <FormLabel className="font-normal">
                                  {category.title}
                                </FormLabel>
                              </FormItem>
                            ),
                        )}
                      </RadioGroup>
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </CardContent>
          </Card>

          {/* Cover Image */}
          <Card className="col-span-4">
            <CardHeader>
              <CardTitle>{protectedEditorConfig.coverImageTitle}</CardTitle>
              <CardDescription>
                {protectedEditorConfig.coverImageDescription}
              </CardDescription>
            </CardHeader>
            <Separator className="" />
            <CardContent className="space-y-4">
              {/* Image */}
              <FormField
                control={form.control}
                name="thumbnail"
                render={({ field }) => (
                  <FormItem className="m-0">
                    <FormControl>
                      <Input
                        placeholder={protectedEditorConfig.placeholderImage}
                        {...field}
                        disabled={true}
                        className="hidden bg-gray-50"
                      />
                    </FormControl>
                  </FormItem>
                )}
              />

              <div className="flex w-full flex-col">
                <FileUploadModal
                  open={showCoverModal}
                  onOpenChange={setShowCoverModal}
                  folderName={imageFolderName}
                  maxFiles={1}
                  maxSize={6 * 1024 * 1024}
                  onUploadComplete={handleCoverUploadComplete}
                  onUploadError={handleUploadError}
                  title="Upload Cover Image"
                  description={protectedEditorConfig.formImageNote}
                />
                {/* {coverImageFileName === '' && ( */}
                <div className="col-span-full">
                  <div className="mb-1 flex items-center gap-x-3">
                    <button
                      onClick={() => setShowCoverModal(!showCoverModal)}
                      type="button"
                      className="inline-flex items-center rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
                    >
                      <PaperClipIcon className="mr-1 h-4 w-4" />
                      <span className="">
                        {protectedEditorConfig.formCoverImageUploadFile}
                      </span>
                    </button>
                  </div>
                </div>
                {/* )} */}

                {form.getValues('thumbnail') !== '' ? (
                  <EditorUploadCoverImageItem
                    userId={userId}
                    postId={post.id}
                    imageUrl={form.getValues('thumbnail')!}
                    deleteCoverHandler={() => form.setValue('thumbnail', '')}
                  />
                ) : (
                  <EditorUploadCoverImagePlaceHolder />
                )}
              </div>
            </CardContent>
          </Card>

          {/* Short Description */}
          <Card className="col-span-8">
            <CardHeader>
              <CardTitle>
                {protectedEditorConfig.shortDescriptionTitle}
              </CardTitle>
              <CardDescription>
                {protectedEditorConfig.shortDescriptionDescription}
              </CardDescription>
            </CardHeader>
            <Separator className="" />
            <CardContent className="space-y-4">
              {/* Description */}
              <FormField
                control={form.control}
                name="summary"
                render={({ field }) => (
                  <FormItem>
                    <FormControl>
                      <Textarea
                        placeholder={
                          protectedEditorConfig.placeholderDescription
                        }
                        className="resize-none"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </CardContent>
          </Card>

          <div className="col-span-12">
            <TextEditor
              value={content}
              onChange={setContent}
              imageFolderName={imageFolderName}
            />
          </div>

          <div className="col-span-12 infline-flex flex items-center justify-end space-x-3 mb-8">
            {post.published ? (
              <>
                <Button
                  type="submit"
                  className="flex px-10 text-red-500! border-red-500! hover:bg-red-100!"
                  disabled={isSaving}
                  variant={'outline'}
                  onClick={() => form.setValue('published', false)}
                >
                  {protectedEditorConfig.unpublish}
                </Button>

                <Button
                  type="submit"
                  className="flex bg-gray-900! px-10 text-white! hover:bg-gray-800!"
                  disabled={isSaving}
                >
                  {protectedEditorConfig.submit}
                </Button>
              </>
            ) : (
              <>
                <Button
                  type="submit"
                  className="flex px-10 text-white! bg-green-500! hover:bg-green-600!"
                  disabled={isSaving}
                  onClick={() => form.setValue('published', true)}
                >
                  {protectedEditorConfig.publish}
                </Button>

                <Button
                  type="submit"
                  className="flex bg-gray-900! px-10 text-white! hover:bg-gray-800!"
                  disabled={isSaving}
                >
                  {protectedEditorConfig.saveDraft}
                </Button>
              </>
            )}

            <Button
              type="button"
              onClick={() => router.back()}
              className="flex !bg-gray-100 px-10 !text-gray-900 hover:!bg-gray-200"
              disabled={isSaving}
            >
              {protectedEditorConfig.cancel}
            </Button>
          </div>
        </form>
      </Form>
      <AlertDialog open={showLoadingAlert} onOpenChange={setShowLoadingAlert}>
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

export default Editor;
