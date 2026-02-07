import {
  Body,
  Controller,
  Delete,
  Get,
  Inject,
  Param,
  Post,
  Put,
  Query,
} from '@nestjs/common';
import { Throttle } from '@nestjs/throttler';

import { PostService } from './post.service';
import {
  PostBaseCreateDto,
  PostBaseDto,
  PostBaseUpdateDto,
  PostBriefResponseDto,
  PostDetailResponseDto,
  PostQueryDto,
} from './dto';
import {
  RedisService,
  type RedisServiceType,
  ConfigService,
  Serialize,
} from '@phanhotboy/nsv-common';
import { RATE_LIMIT } from '@gateway/config';
import { Public, Permissions, CurrentUser } from '@gateway/common/decorators';

@Controller('posts')
export class PostController {
  private readonly routePath = '/api/v1/posts*';

  constructor(
    private readonly postService: PostService,
    @Inject(RedisService) private readonly redis: RedisServiceType,
    private readonly config: ConfigService,
  ) {}

  @Get()
  @Public()
  @Serialize(PostBriefResponseDto)
  getAllPosts(@Query() query: PostQueryDto) {
    return this.postService.getPublishedPosts(query);
  }

  @Get('all')
  @Permissions({ post: ['read'] })
  @Serialize(PostBriefResponseDto)
  getPosts(@Query() query: PostQueryDto) {
    return this.postService.findPosts(query);
  }

  @Get(':id')
  @Public()
  @Serialize(PostDetailResponseDto)
  getPostById(@Param('id') id: string) {
    return this.postService.findPostByIdOrSlug(id);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['create'] })
  async createPost(
    @Body() post: PostBaseCreateDto,
    @CurrentUser('id') authorId: string,
  ) {
    await this.redis.mdel(this.routePath);
    return this.postService.createPost(authorId, post);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['update'] })
  @Serialize(PostBaseDto)
  async updatePost(@Param('id') id: string, @Body() post: PostBaseUpdateDto) {
    await this.redis.mdel(this.routePath);
    return this.postService.updatePost(id, post);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['delete'] })
  async deletePost(
    @Param('id') id: string,
    @CurrentUser('id') authorId: string,
  ) {
    await this.redis.mdel(this.routePath);
    return this.postService.deletePost(id, authorId);
  }

  @Put(':id/publish')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['update'] })
  async publishPost(@Param('id') id: string) {
    await this.redis.mdel(this.routePath);
    return this.postService.publishPost(id);
  }

  @Put(':id/unpublish')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['update'] })
  async unpublishPost(@Param('id') id: string) {
    await this.redis.mdel(this.routePath);
    return this.postService.unpublishPost(id);
  }
}
