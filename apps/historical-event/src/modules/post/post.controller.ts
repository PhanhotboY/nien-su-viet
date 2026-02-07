import { Controller, Logger } from '@nestjs/common';
import { MessagePattern, Payload, Transport } from '@nestjs/microservices';

import { PostService } from './post.service';
import { PostBaseCreateDto, PostBaseUpdateDto, PostQueryDto } from './dto';
import { POST_MESSAGE_PATTERN } from '@phanhotboy/constants';

@Controller('posts')
export class PostController {
  private readonly logger = new Logger(PostController.name);

  constructor(private readonly postService: PostService) {}

  @MessagePattern(POST_MESSAGE_PATTERN.GET_POST_BY_ID)
  getPostById(@Payload('id') id: string) {
    this.logger.log(`Getting post by id or slug: ${id}`);
    return this.postService.findPostByIdOrSlug(id);
  }

  @MessagePattern(POST_MESSAGE_PATTERN.GET_ALL_POSTS)
  getAllPosts(@Payload('query') query: PostQueryDto) {
    this.logger.log(`Getting all posts with query`);
    return this.postService.findPosts(query);
  }

  @MessagePattern(POST_MESSAGE_PATTERN.GET_PUBLISHED_POSTS)
  getPublishedPosts(@Payload('query') query: PostQueryDto) {
    this.logger.log(`Getting published posts with query`);
    return this.postService.getPublishedPosts(query);
  }

  @MessagePattern(POST_MESSAGE_PATTERN.CREATE_POST)
  async createPost(
    @Payload('authorId') authorId: string,
    @Payload('payload') data: PostBaseCreateDto,
  ) {
    this.logger.log(`Creating new post for author: ${authorId}`);
    return this.postService.createPost(authorId, data);
  }

  @MessagePattern(POST_MESSAGE_PATTERN.UPDATE_POST, {
    transport: Transport.TCP,
  })
  async updatePost(
    @Payload('id') id: string,
    @Payload('payload') data: PostBaseUpdateDto,
  ) {
    this.logger.log(`Updating post with id: ${id}`);
    return this.postService.updatePost(id, data);
  }

  @MessagePattern(POST_MESSAGE_PATTERN.DELETE_POST, {
    transport: Transport.TCP,
  })
  async deletePost(@Payload('id') id: string) {
    this.logger.log(`Deleting post with id: ${id}`);
    return this.postService.deletePost(id);
  }

  @MessagePattern(POST_MESSAGE_PATTERN.PUBLISH_POST, {
    transport: Transport.TCP,
  })
  async publishPost(@Payload('id') id: string) {
    this.logger.log(`Publishing post with id: ${id}`);
    return this.postService.publishPost(id);
  }

  @MessagePattern(POST_MESSAGE_PATTERN.UNPUBLISH_POST, {
    transport: Transport.TCP,
  })
  async unpublishPost(@Payload('id') id: string) {
    this.logger.log(`Unpublishing post with id: ${id}`);
    return this.postService.unpublishPost(id);
  }
}
