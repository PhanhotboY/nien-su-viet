'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { components } from '@nsv-interfaces/historical-event';
import { retryFetcher } from '.';
import { IApiResponse } from '../interfaces/response.interface';

const HISTORICAL_EVENT_API_URL =
  process.env.HISTORICAL_EVENT_SERVICE_URL || 'http://localhost:3002/api/v1';

export async function getEvents(
  query?: URLSearchParams,
): Promise<
  IPaginatedResponse<components['schemas']['HistoricalEventBriefResponseDto']>
> {
  const response = (await retryFetcher(
    `${HISTORICAL_EVENT_API_URL}/historical-events?${query?.toString() || ''}`,
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
  const response = await retryFetcher(
    `${HISTORICAL_EVENT_API_URL}/historical-events/${id}`,
  );

  return response;
}

export async function getEventPreview(
  id: string,
): Promise<
  IApiResponse<components['schemas']['HistoricalEventPreviewResponseDto']>
> {
  const response = await retryFetcher(
    `${HISTORICAL_EVENT_API_URL}/historical-events/${id}/preview`,
  );

  return response;
}

export async function createEvent(
  eventData: components['schemas']['HistoricalEventBaseCreateDto'],
): Promise<components['schemas']['HistoricalEventDetailResponseDto']> {
  const response = await retryFetcher<
    components['schemas']['HistoricalEventDetailResponseDto']
  >(`${HISTORICAL_EVENT_API_URL}/historical-events`, {
    method: 'POST',
    body: JSON.stringify(eventData),
  });

  return response.data;
}

export async function updateEvent(
  id: string,
  eventData: Partial<components['schemas']['HistoricalEventBaseCreateDto']>,
): Promise<components['schemas']['HistoricalEventDetailResponseDto']> {
  const response = await retryFetcher(
    `${HISTORICAL_EVENT_API_URL}/historical-events/${id}`,
    {
      method: 'PUT',
      body: JSON.stringify(eventData),
    },
  );
  return response.data;
}

export async function deleteEvent(id: string): Promise<void> {
  const response = await retryFetcher(
    `${HISTORICAL_EVENT_API_URL}/historical-events/${id}`,
    {
      method: 'DELETE',
    },
  );
  return response.data;
}

export async function getCategories(): Promise<
  components['schemas']['EventCategoryBriefResponseDto'][]
> {
  const response = await retryFetcher(
    `${HISTORICAL_EVENT_API_URL}/event-categories`,
  );

  return response.data;
}
