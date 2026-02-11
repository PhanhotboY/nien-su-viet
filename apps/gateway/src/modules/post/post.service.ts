import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { firstValueFrom, timeout, catchError, throwError } from 'rxjs';

import {
  PostQueryDto,
  PostBriefResponseDto,
  PostDetailResponseDto,
  PostBaseCreateDto,
  PostBaseUpdateDto,
} from './dto';
import { TCP_SERVICE } from '@phanhotboy/constants/tcp-service.constant';
import { POST_MESSAGE_PATTERN } from '@phanhotboy/constants';
import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import { PaginatedResponseDto } from '@phanhotboy/nsv-common';

@Injectable()
export class PostService {
  private readonly serviceName = 'Post Service';

  constructor(
    @Inject(TCP_SERVICE.HISTORICAL_EVENT.NAME)
    private readonly postClient: ClientProxy,
  ) {}

  async createPost(authorId: string, payload: PostBaseCreateDto) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient
            .send(POST_MESSAGE_PATTERN.CREATE_POST, {
              authorId,
              payload,
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'create post',
      this.serviceName,
    );
  }

  async findPosts(
    query: PostQueryDto,
  ): Promise<PaginatedResponseDto<PostBriefResponseDto>> {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient
            .send(POST_MESSAGE_PATTERN.GET_ALL_POSTS, { query })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get posts',
      this.serviceName,
    );
  }

  async getPublishedPosts(
    query: PostQueryDto,
  ): Promise<PaginatedResponseDto<PostBriefResponseDto>> {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient
            .send(POST_MESSAGE_PATTERN.GET_PUBLISHED_POSTS, { query })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get published posts',
      this.serviceName,
    );
  }

  async findPostByIdOrSlug(id: string): Promise<PostDetailResponseDto> {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient
            .send(POST_MESSAGE_PATTERN.GET_POST_BY_ID, { id })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get post by id or slug',
      this.serviceName,
    );
  }

  async updatePost(id: string, payload: PostBaseUpdateDto) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient
            .send(POST_MESSAGE_PATTERN.UPDATE_POST, {
              id,
              payload,
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'update post',
      this.serviceName,
    );
  }

  async deletePost(id: string, authorId: string) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient
            .send(POST_MESSAGE_PATTERN.DELETE_POST, {
              id,
              authorId,
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'delete post',
      this.serviceName,
    );
  }

  async publishPost(id: string) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient.send(POST_MESSAGE_PATTERN.PUBLISH_POST, { id }).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'publish post',
      this.serviceName,
    );
  }

  async unpublishPost(id: string) {
    return MicroserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postClient
            .send(POST_MESSAGE_PATTERN.UNPUBLISH_POST, { id })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'unpublish post',
      this.serviceName,
    );
  }
}
