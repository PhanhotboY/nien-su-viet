'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { retryFetcher } from '.';
import { Post } from '@/types/collection';

export async function getPublicPosts(
  query?: Record<string, string> | string,
): Promise<IPaginatedResponse<Post>> {
  const response = (await retryFetcher(
    `/posts?${new URLSearchParams(query).toString()}`,
  )) as IPaginatedResponse<Post>;

  return response;
}

export async function findPosts(
  query?: Record<string, string> | string,
): Promise<IPaginatedResponse<Post>> {
  const response = (await retryFetcher(
    `/posts/all?${new URLSearchParams(query).toString()}`,
  )) as IPaginatedResponse<Post>;

  return response;
}

export async function getPost(id: string) {
  const response = await retryFetcher<Post>(`/posts/${id}`, {
    method: 'GET',
  });

  return response;
}

export async function createPost(post: Partial<Post>) {
  const response = await retryFetcher<IOperationResult>(`/posts`, {
    method: 'POST',
    body: JSON.stringify(post),
  });

  return response;
}

export async function updatePost(id: string, post: Partial<Post>) {
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
