import { Inject, Injectable } from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';

import { PrismaService } from '@historical-event/database';
import {
  getExcerpt,
  RedisService,
  UtilService,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import { PaginatedResponseDto } from '@phanhotboy/nsv-common/dto/response';
import { Prisma } from '@historical-event-prisma';
import {
  PostQueryDto,
  PostBriefResponseDto,
  PostDetailResponseDto,
  PostBaseCreateDto,
  PostBaseUpdateDto,
} from './dto';
import {
  OperationResponse,
  OperationResponseDto,
} from '@phanhotboy/nsv-common/dto/response/operation-response.dto';

@Injectable()
export class PostService {
  private readonly cachePrefix = 'posts';
  private readonly cacheKey: string;
  private readonly cacheTTL = 3600; // 1 hour in seconds

  constructor(
    private readonly prisma: PrismaService,
    private readonly util: UtilService,
    @Inject(RedisService)
    private readonly redisService: RedisServiceType,
  ) {
    this.cacheKey = this.util.genCacheKey(this.cachePrefix);
  }

  /**
   * Find posts with query filters and pagination
   */
  async findPosts(
    query: PostQueryDto,
  ): Promise<PaginatedResponseDto<PostBriefResponseDto>> {
    const page = query.page || 1;
    const limit = query.limit || 10;
    const skip = (page - 1) * limit;

    const where = this.buildWhereClause(query, false);
    const orderBy = this.buildOrderByClause(query);

    try {
      const [posts, total] = await Promise.all([
        this.prisma.post.findMany({
          where,
          skip,
          take: limit,
          orderBy,
          include: {
            author: {
              select: {
                id: true,
                email: true,
                name: true,
                image: true,
              },
            },
            category: true,
          },
        }),
        this.prisma.post.count({ where }),
      ]);

      return {
        data: posts as unknown as PostBriefResponseDto[],
        pagination: {
          page,
          limit,
          total,
          totalPages: Math.ceil(total / limit),
        },
      };
    } catch (error) {
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể lấy danh sách bài viết',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Get published posts for public access
   */
  async getPublishedPosts(
    query: PostQueryDto,
  ): Promise<PaginatedResponseDto<PostBriefResponseDto>> {
    const page = query.page || 1;
    const limit = query.limit || 10;
    const skip = (page - 1) * limit;

    const cacheKey = `${this.cacheKey}:published:${JSON.stringify(query)}`;

    // Try to get from cache
    const cached = await this.redisService.get(cacheKey);
    if (cached) {
      return cached;
    }

    const where = this.buildWhereClause(query, true);
    const orderBy = this.buildOrderByClause(query);

    try {
      const [posts, total] = await Promise.all([
        this.prisma.post.findMany({
          where,
          skip,
          take: limit,
          orderBy,
          include: {
            author: {
              select: {
                id: true,
                email: true,
                name: true,
                image: true,
              },
            },
            category: true,
          },
        }),
        this.prisma.post.count({ where }),
      ]);

      const result = {
        data: posts as unknown as PostBriefResponseDto[],
        pagination: {
          page,
          limit,
          total,
          totalPages: Math.ceil(total / limit),
        },
      };

      // Cache the result
      await this.redisService.setEx(
        cacheKey,
        this.cacheTTL,
        JSON.stringify(result),
      );

      return result;
    } catch (error) {
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể lấy danh sách bài viết công khai',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Find post by ID or slug
   */
  async findPostByIdOrSlug(idOrSlug: string): Promise<PostDetailResponseDto> {
    const cacheKey = `${this.cacheKey}:detail:${idOrSlug}`;

    // Try to get from cache
    const cached = await this.redisService.get(cacheKey);
    if (cached) {
      return cached;
    }

    try {
      let post;

      // Check if it's a valid UUID, then search by ID, otherwise by slug
      const uuidRegex =
        /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i;
      if (uuidRegex.test(idOrSlug)) {
        post = await this.prisma.post.findUnique({
          where: { id: idOrSlug },
          include: {
            author: {
              select: {
                id: true,
                email: true,
                name: true,
                image: true,
              },
            },
            category: true,
          },
        });
      } else {
        post = await this.prisma.post.findUnique({
          where: { slug: idOrSlug },
          include: {
            author: {
              select: {
                id: true,
                email: true,
                name: true,
                image: true,
              },
            },
            category: true,
          },
        });
      }

      if (!post) {
        throw new RpcException({
          statusCode: 404,
          message: 'Không tìm thấy bài viết',
        });
      }

      const result = post as unknown as PostDetailResponseDto;

      // Cache the result
      await this.redisService.setEx(
        cacheKey,
        this.cacheTTL,
        JSON.stringify(result),
      );

      return result;
    } catch (error) {
      if (error instanceof RpcException) {
        throw error;
      }
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể lấy thông tin bài viết',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Create a new post
   */
  async createPost(
    authorId: string,
    data: PostBaseCreateDto,
  ): Promise<OperationResponse> {
    try {
      // Generate excerpt if summary is not provided
      if (!data.summary && data.content) {
        data.summary = getExcerpt(data.content, 200);
      }
      if (data.authorId) {
        delete data.authorId;
      }

      const post = await this.prisma.post.create({
        data: {
          ...data,
          publishedAt: data.publishedAt || new Date(),
          published: data.published || false,
          author: {
            connect: { id: authorId },
          },
        } as Prisma.PostCreateInput,
      });

      // Invalidate cache
      await this.invalidateCache();

      return { success: true, id: post.id };
    } catch (error) {
      if (error instanceof Prisma.PrismaClientKnownRequestError) {
        if (error.code === 'P2002') {
          throw new RpcException({
            statusCode: 409,
            message: 'Slug đã tồn tại',
          });
        }
      }
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể tạo bài viết',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Update an existing post
   */
  async updatePost(
    id: string,
    data: PostBaseUpdateDto,
  ): Promise<OperationResponse> {
    try {
      // Check if post exists
      const existingPost = await this.prisma.post.findUnique({
        where: { id },
      });

      if (!existingPost) {
        throw new RpcException({
          statusCode: 404,
          message: 'Không tìm thấy bài viết',
        });
      }

      // Generate excerpt if summary is not provided but content is updated
      if (!data.summary && data.content) {
        data.summary = getExcerpt(data.content, 200);
      }

      // not implemented
      delete data.categoryId;

      console.log('updating post with data: ', { ...data, content: undefined });
      const post = await this.prisma.post.update({
        where: { id },
        data: data as Prisma.PostUpdateInput,
      });

      // Invalidate cache
      await this.invalidateCache(id);

      return { id: post.id, success: true };
    } catch (error) {
      if (error instanceof RpcException) {
        throw error;
      }
      if (error instanceof Prisma.PrismaClientKnownRequestError) {
        if (error.code === 'P2002') {
          throw new RpcException({
            statusCode: 409,
            message: 'Slug đã tồn tại',
          });
        }
        if (error.code === 'P2025') {
          throw new RpcException({
            statusCode: 404,
            message: 'Không tìm thấy bài viết',
          });
        }
      }
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể cập nhật bài viết',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Delete a post
   */
  async deletePost(id: string): Promise<OperationResponse> {
    try {
      await this.prisma.post.delete({
        where: { id },
      });

      // Invalidate cache
      await this.invalidateCache(id);

      return { id, success: true };
    } catch (error) {
      if (error instanceof Prisma.PrismaClientKnownRequestError) {
        if (error.code === 'P2025') {
          throw new RpcException({
            statusCode: 404,
            message: 'Không tìm thấy bài viết',
          });
        }
      }
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể xóa bài viết',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Publish a post
   */
  async publishPost(id: string): Promise<string> {
    try {
      const post = await this.prisma.post.findUnique({
        where: { id },
      });

      if (!post) {
        throw new RpcException({
          statusCode: 404,
          message: 'Không tìm thấy bài viết',
        });
      }

      await this.prisma.post.update({
        where: { id },
        data: {
          published: true,
          publishedAt: post.published ? post.publishedAt : new Date(),
        },
      });

      // Invalidate cache
      await this.invalidateCache(id);

      return id;
    } catch (error) {
      if (error instanceof RpcException) {
        throw error;
      }
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể xuất bản bài viết',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Unpublish a post
   */
  async unpublishPost(id: string): Promise<string> {
    try {
      const post = await this.prisma.post.findUnique({
        where: { id },
      });

      if (!post) {
        throw new RpcException({
          statusCode: 404,
          message: 'Không tìm thấy bài viết',
        });
      }

      await this.prisma.post.update({
        where: { id },
        data: {
          published: false,
        },
      });

      // Invalidate cache
      await this.invalidateCache(id);

      return id;
    } catch (error) {
      if (error instanceof RpcException) {
        throw error;
      }
      throw new RpcException({
        statusCode: 500,
        message: 'Không thể hủy xuất bản bài viết',
        error: error instanceof Error ? error.message : 'Unknown error',
      });
    }
  }

  /**
   * Build where clause for Prisma query
   */
  private buildWhereClause(
    query: PostQueryDto,
    publishedOnly: boolean = false,
  ): Prisma.PostWhereInput {
    const where: Prisma.PostWhereInput = {};

    // Filter by published status
    if (publishedOnly) {
      where.published = true;
    }

    // Filter by author
    if (query.authorId) {
      where.authorId = query.authorId;
    }

    // Filter by category
    if (query.categoryIds && query.categoryIds.length > 0) {
      where.categoryId = {
        in: query.categoryIds,
      };
    }

    // Search by title or content
    if (query.search) {
      where.OR = [
        {
          title: {
            contains: query.search,
            mode: 'insensitive',
          },
        },
        {
          content: {
            contains: query.search,
            mode: 'insensitive',
          },
        },
        {
          summary: {
            contains: query.search,
            mode: 'insensitive',
          },
        },
      ];
    }

    // Filter by created date range
    if (query.createdAtFrom || query.createdAtTo) {
      where.createdAt = {};
      if (query.createdAtFrom) {
        where.createdAt.gte = new Date(query.createdAtFrom);
      }
      if (query.createdAtTo) {
        where.createdAt.lte = new Date(query.createdAtTo);
      }
    }

    // Filter by updated date range
    if (query.updatedAtFrom || query.updatedAtTo) {
      where.updatedAt = {};
      if (query.updatedAtFrom) {
        where.updatedAt.gte = new Date(query.updatedAtFrom);
      }
      if (query.updatedAtTo) {
        where.updatedAt.lte = new Date(query.updatedAtTo);
      }
    }

    return where;
  }

  /**
   * Build order by clause for Prisma query
   */
  private buildOrderByClause(
    query: PostQueryDto,
  ): Prisma.PostOrderByWithRelationInput {
    const sortBy = query.sortBy || 'createdAt';
    const sortOrder = query.sortOrder || 'desc';

    return {
      [sortBy]: sortOrder,
    };
  }

  /**
   * Invalidate cache for posts
   */
  private async invalidateCache(postId?: string): Promise<void> {
    try {
      // Delete all cache keys with the prefix
      const pattern = `${this.cacheKey}*`;
      const keys = await this.redisService.keys(pattern);

      if (keys && keys.length > 0) {
        for (const key of keys) {
          await this.redisService.del(key);
        }
      }

      // Also invalidate specific post cache if ID is provided
      if (postId) {
        await this.redisService.del(`${this.cacheKey}:detail:${postId}`);
      }
    } catch (error) {
      // Log error but don't throw - cache invalidation failure shouldn't break the operation
      console.error('Failed to invalidate cache:', error);
    }
  }
}
