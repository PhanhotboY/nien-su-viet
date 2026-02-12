'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { retryFetcher } from '.';
import {
  PostBriefResponseDto,
  PostCreateDto,
  PostResponseDto,
  PostUpdateDto,
} from '@/types/collection';

export async function getPublicPosts(
  query?: Record<string, string> | string,
): Promise<IPaginatedResponse<PostBriefResponseDto>> {
  const response = (await retryFetcher(
    `/posts?${new URLSearchParams(query).toString()}`,
  )) as IPaginatedResponse<PostResponseDto>;

  return response;
}

export async function findPosts(
  query?: Record<string, string> | string,
): Promise<IPaginatedResponse<PostBriefResponseDto>> {
  const response = (await retryFetcher(
    `/posts/all?${new URLSearchParams(query).toString()}`,
  )) as IPaginatedResponse<PostResponseDto>;

  return response;
}

export async function getPost(id: string) {
  const response = await retryFetcher<PostResponseDto>(`/posts/${id}`, {
    method: 'GET',
  });

  return response;
}

export async function createPost(post: PostCreateDto) {
  const response = await retryFetcher<IOperationResult>(`/posts`, {
    method: 'POST',
    body: JSON.stringify(post),
  });

  return response;
}

export async function updatePost(id: string, post: PostUpdateDto) {
  const response = await retryFetcher<IOperationResult>(`/posts/${id}`, {
    method: 'PUT',
    body: JSON.stringify(post),
  });

  return response;
}

export async function deletePost(id: string) {
  const response = await retryFetcher<IOperationResult>(`/posts/${id}`, {
    method: 'DELETE',
  });

  return response;
}
