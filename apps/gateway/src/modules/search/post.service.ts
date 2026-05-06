import { Inject, Injectable, OnModuleInit, Logger } from '@nestjs/common';
import type { ClientGrpc } from '@nestjs/microservices';
import { plainToInstance } from 'class-transformer';
import { firstValueFrom, timeout, catchError, throwError } from 'rxjs';

import { PostQueryDto, PostQueryGrpcDto } from './dto';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import {
  POSTS_SERVICE_NAME,
  PostsServiceClient,
} from '@phanhotboy/genproto/search_service/posts';

@Injectable()
export class SearchPostService {
  private readonly serviceName = 'Post Service';
  private postService: PostsServiceClient;
  private microserviceErrorHandler: MicroserviceErrorHandler;

  constructor(
    @Inject(GRPC_SERVICE.SEARCH.NAME)
    private readonly postClient: ClientGrpc,
    private readonly logger: Logger,
  ) {
    this.microserviceErrorHandler = new MicroserviceErrorHandler(this.logger);
    this.postService =
      this.postClient.getService<PostsServiceClient>(POSTS_SERVICE_NAME);
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
}
