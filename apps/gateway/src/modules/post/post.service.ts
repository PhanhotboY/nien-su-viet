import { Inject, Injectable, OnModuleInit, Logger } from '@nestjs/common';
import type { ClientGrpc } from '@nestjs/microservices';
import { plainToInstance } from 'class-transformer';
import { firstValueFrom, timeout, catchError, throwError } from 'rxjs';

import {
  PostQueryDto,
  PostBaseCreateDto,
  PostBaseUpdateDto,
  PostBaseCreateGrpcDto,
  PostBaseUpdateGrpcDto,
  PostQueryGrpcDto,
} from './dto';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import {
  POSTS_SERVICE_NAME,
  PostsServiceClient,
} from '@phanhotboy/genproto/post_service/posts';

@Injectable()
export class PostService {
  private readonly serviceName = 'Post Service';
  private postService: PostsServiceClient;
  private microserviceErrorHandler: MicroserviceErrorHandler;

  constructor(
    @Inject(GRPC_SERVICE.POST.NAME)
    private readonly postClient: ClientGrpc,
    private readonly logger: Logger,
  ) {
    this.microserviceErrorHandler = new MicroserviceErrorHandler(this.logger);
    this.postService =
      this.postClient.getService<PostsServiceClient>(POSTS_SERVICE_NAME);
  }

  async createPost(authorId: string, payload: PostBaseCreateDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService
            .createPost({
              ...plainToInstance(PostBaseCreateGrpcDto, payload),
              authorId,
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

  async getAllPosts(query: PostQueryDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService
            .getAllPosts(plainToInstance(PostQueryGrpcDto, query))
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get posts',
      this.serviceName,
    );
  }

  async getPublishedPosts(query: PostQueryDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService
            .getPublishedPosts(plainToInstance(PostQueryGrpcDto, query))
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'get published posts',
      this.serviceName,
    );
  }

  async findPostByIdOrSlug(id: string) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService.getPost({ id }).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'get post by id or slug',
      this.serviceName,
    );
  }

  async updatePost(id: string, payload: PostBaseUpdateDto) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService
            .updatePost({
              id,
              ...plainToInstance(PostBaseUpdateGrpcDto, payload),
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
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService
            .deletePost({
              id,
              // authorId,
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
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService.publishPost({ id }).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'publish post',
      this.serviceName,
    );
  }

  async unpublishPost(id: string) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.postService.unpublishPost({ id }).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'unpublish post',
      this.serviceName,
    );
  }
}
