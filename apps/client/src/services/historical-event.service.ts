'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { components } from '@nsv-interfaces/historical-event';
import { retryFetcher } from '.';
import { IApiResponse } from '../interfaces/response.interface';

export async function getEvents(
  query?: Record<string, string> | string,
): Promise<
  IPaginatedResponse<components['schemas']['HistoricalEventBriefResponseDto']>
> {
  const response = (await retryFetcher(
    `/historical-events?${new URLSearchParams(query).toString()}`,
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
  const response = await retryFetcher(`/historical-events/${id}`);

  return response;
}

export async function getEventPreview(
  id: string,
): Promise<
  IApiResponse<components['schemas']['HistoricalEventPreviewResponseDto']>
> {
  const response = await retryFetcher(`/historical-events/${id}/preview`);

  return response;
}

export async function createEvent(
  eventData: components['schemas']['HistoricalEventBaseCreateDto'],
): Promise<components['schemas']['HistoricalEventDetailResponseDto']> {
  const response = await retryFetcher<
    components['schemas']['HistoricalEventDetailResponseDto']
  >(`/historical-events`, {
    method: 'POST',
    body: JSON.stringify(eventData),
  });

  return response.data;
}

export async function updateEvent(
  id: string,
  eventData: Partial<components['schemas']['HistoricalEventBaseCreateDto']>,
): Promise<components['schemas']['HistoricalEventDetailResponseDto']> {
  const response = await retryFetcher(`/historical-events/${id}`, {
    method: 'PUT',
    body: JSON.stringify(eventData),
  });
  return response.data;
}

export async function deleteEvent(id: string): Promise<void> {
  const response = await retryFetcher(`/historical-events/${id}`, {
    method: 'DELETE',
  });
  return response.data;
}

export async function getCategories(): Promise<
  components['schemas']['EventCategoryBriefResponseDto'][]
> {
  const response = await retryFetcher(`/event-categories`);

  return response.data;
}
