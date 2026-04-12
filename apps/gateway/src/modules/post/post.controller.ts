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
import {
  ApiCreatedSerializedResponse,
  ApiOkSerializedOperationResponse,
  ApiOkSerializedPaginatedResponse,
  ApiOkSerializedResponse,
} from '@phanhotboy/nsv-common/decorators';
import { ClientRMQ } from '@nestjs/microservices';
import { RMQ } from '@phanhotboy/constants';

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
  @ApiOkSerializedPaginatedResponse(PostBriefResponseDto)
  async getPublishedPosts(@Query() query: any) {
    const res = await this.postService.getPublishedPosts(query);
    return res;
  }

  @Get('all')
  @Permissions({ post: ['read'] })
  @Serialize(PostBriefResponseDto)
  @ApiOkSerializedPaginatedResponse(PostBriefResponseDto)
  async getAllPosts(@Query() query: PostQueryDto) {
    const res = await this.postService.getAllPosts(query);
    return res;
  }

  @Get(':id')
  @Public()
  @Serialize(PostDetailResponseDto)
  @ApiOkSerializedResponse(PostDetailResponseDto)
  getPostById(@Param('id') id: string) {
    return this.postService.findPostByIdOrSlug(id);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['create'] })
  @ApiCreatedSerializedResponse()
  async createPost(
    @Body() post: PostBaseCreateDto,
    @CurrentUser('id') authorId: string,
  ) {
    await this.redis.mdel(this.routePath);
    return await this.postService.createPost(authorId, post);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['update'] })
  @ApiOkSerializedOperationResponse()
  async updatePost(@Param('id') id: string, @Body() post: PostBaseUpdateDto) {
    await this.redis.mdel(this.routePath);
    return this.postService.updatePost(id, post);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['delete'] })
  @ApiOkSerializedOperationResponse()
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
  @ApiOkSerializedOperationResponse()
  async publishPost(@Param('id') id: string) {
    await this.redis.mdel(this.routePath);
    return this.postService.publishPost(id);
  }

  @Put(':id/unpublish')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ post: ['update'] })
  @ApiOkSerializedOperationResponse()
  async unpublishPost(@Param('id') id: string) {
    await this.redis.mdel(this.routePath);
    return this.postService.unpublishPost(id);
  }
}
