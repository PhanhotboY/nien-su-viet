import { Controller, Get, Inject, Param, Query } from '@nestjs/common';
import { Throttle } from '@nestjs/throttler';

import { SearchPostService } from './post.service';
import {
  PostBriefResponseDto,
  PostDetailResponseDto,
  PostQueryDto,
} from './dto';
import { Serialize } from '@phanhotboy/nsv-common';
import { RATE_LIMIT } from '@gateway/config';
import { Public, Permissions, CurrentUser } from '@gateway/common/decorators';
import {
  ApiOkSerializedPaginatedResponse,
  ApiOkSerializedResponse,
} from '@phanhotboy/nsv-common/decorators';

@Controller('search/posts')
export class SearchPostController {
  constructor(private readonly postService: SearchPostService) {}

  @Get()
  @Public()
  @Serialize(PostBriefResponseDto)
  @ApiOkSerializedPaginatedResponse(PostBriefResponseDto)
  async getPublishedPosts(@Query() query: PostQueryDto) {
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
}
