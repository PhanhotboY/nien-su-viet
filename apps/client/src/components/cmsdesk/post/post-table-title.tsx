import PostCreateButton from '@/components/cmsdesk/post/buttons/post-create-button';
import { protectedPostConfig } from '@/config/cmsdesk';

const PostTableTitle = () => {
  return (
    <>
      <div className="mb-5 flex flex-row border-b border-gray-200 pb-5">
        <div className="flex-none items-center justify-start">
          <h1 className="text-3xl font-bold tracking-tight">
            {protectedPostConfig.title}
          </h1>
          <p className="text-muted-foreground">
            {protectedPostConfig.description}
          </p>
        </div>
        <div className="flex-grow"></div>
        <div className="flex-none items-center justify-end">
          <PostCreateButton />
        </div>
      </div>
    </>
  );
};

export default PostTableTitle;
