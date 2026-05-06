'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { retryFetcher } from '.';
import { IApiResponse } from '../interfaces/response.interface';
import { avoidRateLimit } from '@/helper/rate-limit.helper';
import { revalidateTag } from 'next/cache';

const historicalEventCacheTag = 'historical-events';
const nextFetchOptions = {
  tags: [historicalEventCacheTag],
};

export async function getEvents(
  query?: Record<string, string> | string,
): Promise<
  IPaginatedResponse<components['schemas']['HistoricalEventBriefResponseDto']>
> {
  const response = (await retryFetcher(
    `/search/historical-events?${new URLSearchParams(query).toString()}`,
    {
      isPublicRoute: true,
      next: nextFetchOptions,
    },
  )) as IPaginatedResponse<
    components['schemas']['HistoricalEventBriefResponseDto']
  >;

  return response;
}

export async function getEvent(
  id: string,
): Promise<
  IApiResponse<components['schemas']['HistoricalEventDetailResponseDto']>
> {
  await avoidRateLimit();
  const response = await retryFetcher(`/search/historical-events/${id}`, {
    isPublicRoute: true,
    next: nextFetchOptions,
  });

  return response;
}

export async function getEventPreview(
  id: string,
): Promise<
  IApiResponse<components['schemas']['HistoricalEventPreviewResponseDto']>
> {
  const response = await retryFetcher(
    `/search/historical-events/${id}/preview`,
    {
      isPublicRoute: true,
      next: nextFetchOptions,
    },
  );

  return response;
}

export async function createEvent(
  eventData: components['schemas']['HistoricalEventBaseCreateDto'],
): Promise<IApiResponse<components['schemas']['OperationMetadataDto']>> {
  const response = await retryFetcher<
    components['schemas']['OperationMetadataDto']
  >(`/historical-events`, {
    method: 'POST',
    body: JSON.stringify(eventData),
  });
  revalidateTag(historicalEventCacheTag, 'max');

  return response;
}

export async function updateEvent(
  id: string,
  eventData: Partial<components['schemas']['HistoricalEventBaseCreateDto']>,
): Promise<IApiResponse<components['schemas']['OperationMetadataDto']>> {
  const response = await retryFetcher(`/historical-events/${id}`, {
    method: 'PUT',
    body: JSON.stringify(eventData),
  });
  revalidateTag(historicalEventCacheTag, 'max');
  return response;
}

export async function deleteEvent(
  id: string,
): Promise<IApiResponse<components['schemas']['OperationMetadataDto']>> {
  const response = await retryFetcher(`/historical-events/${id}`, {
    method: 'DELETE',
  });
  revalidateTag(historicalEventCacheTag, 'max');
  return response;
}

export async function getCategories(): Promise<
  IPaginatedResponse<components['schemas']['EventCategoryBriefResponseDto'][]>
> {
  const response = (await retryFetcher(
    `/event-categories`,
  )) as IPaginatedResponse<
    components['schemas']['EventCategoryBriefResponseDto'][]
  >;

  return response;
}
