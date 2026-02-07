import Editor from '@/components/cmsdesk/editor/editor';
import { Separator } from '@/components/ui/separator';
import { protectedEditorConfig } from '@/config/cmsdesk';
import { getPost } from '@/services/post.service';
import { notFound } from 'next/navigation';

export const revalidate = 0;

interface PostEditorPageProps {
  params: Promise<{ postId: string }>;
}

async function getPostHandler(postId: string) {
  'use server';
  const res = await getPost(postId);

  if (!res.data) {
    console.log('Error has occured while getting post data');
    console.log('Error message : ', res.message);
    return null;
  }

  return res.data;
}

export default async function PostEditorPage({ params }: PostEditorPageProps) {
  const resolvedParams = await params;

  const post = await getPostHandler(resolvedParams.postId);

  if (!post) {
    notFound();
  }

  return (
    <div className="container px-10">
      <div>
        <h3 className="text-lg font-medium">{protectedEditorConfig.title}</h3>
        <p className="py-2 text-sm text-muted-foreground">
          {protectedEditorConfig.description}
        </p>
      </div>
      <Separator className="mb-5" />
      <Editor post={post} imageFolderName={`posts/${resolvedParams.postId}`} />
    </div>
  );
}
