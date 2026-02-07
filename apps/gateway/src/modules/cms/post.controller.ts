import type { Request } from 'express';
import {
  Body,
  Controller,
  Delete,
  Get,
  Inject,
  Post,
  Put,
  Query,
  Req,
} from '@nestjs/common';

import { Permissions, Public } from '@gateway/common/decorators';
import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { RedisService, type RedisServiceType } from '@phanhotboy/nsv-common';
import { PostBaseCreateDto } from './dto/post-base-create.dto';
import { PostQueryDto } from './dto/post-query.dto';
import { PostBaseUpdateDto } from './dto/post-base-update.dto';

@Controller('posts')
export class PostController {
  private readonly routePath = '/api/v1/posts*';
  constructor(
    private readonly cmsProxy: HttpProxyService,
    @Inject(RedisService) private readonly redis: RedisServiceType,
  ) {}

  @Get()
  @Public()
  proxyGetPosts(@Req() req: Request, @Query() query: PostQueryDto) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Get('all')
  @Public()
  // @Permissions({ post: ['create'] }) // only users have create post permission can get unpublished posts
  proxyFindAllPosts(@Req() req: Request, @Query() query: PostQueryDto) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Get(':slug')
  @Public()
  proxyGetPostBySlug(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Post()
  @Permissions({ post: ['create'] })
  async proxyCreatePost(@Req() req: Request, @Body() body: PostBaseCreateDto) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Permissions({ post: ['update'] })
  async proxyUpdatePost(@Req() req: Request, @Body() body: PostBaseUpdateDto) {
    console.log('updating post with data: ', body);
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Permissions({ post: ['delete'] })
  async proxyDeletePost(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
