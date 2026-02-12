'use server';

import { components } from '@nsv-interfaces/nsv-api-documentation';
import { retryFetcher } from '.';
import { IApiResponse } from '@/interfaces/response.interface';

async function getAppInfo(
  query?: Record<string, string> | string,
): Promise<IApiResponse<components['schemas']['AppDto']>> {
  const response = await retryFetcher(
    `/app?${new URLSearchParams(query).toString()}`,
  );

  return response;
}

async function updateApp(
  app: Partial<components['schemas']['AppUpdateDto']>,
): Promise<IApiResponse<components['schemas']['OperationResponse']>> {
  const response = await retryFetcher(`/app`, {
    method: 'PUT',
    body: JSON.stringify(app),
  });
  return response.data;
}

async function getHeaderNavItems(): Promise<
  IApiResponse<Array<components['schemas']['HeaderNavItemDto']>>
> {
  const response = await retryFetcher(`/header-nav-items`);

  return response;
}

async function updateHeaderNavItem(
  id: string,
  item: Partial<components['schemas']['HeaderNavItemUpdateDto']>,
): Promise<IApiResponse<components['schemas']['OperationResponse']>> {
  const response = await retryFetcher(`/header-nav-items/${id}`, {
    method: 'PUT',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function getFooterNavItems(): Promise<
  IApiResponse<Array<components['schemas']['FooterNavItemDto']>>
> {
  const response = await retryFetcher(`/footer-nav-items`);

  return response;
}

async function updateFooterNavItem(
  id: string,
  item: Partial<components['schemas']['FooterNavItemCreateDto']>,
): Promise<IApiResponse<components['schemas']['OperationResponse']>> {
  const response = await retryFetcher(`/footer-nav-items/${id}`, {
    method: 'PUT',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function createHeaderNavItem(
  item: components['schemas']['HeaderNavItemCreateDto'],
): Promise<IApiResponse<components['schemas']['OperationResponse']>> {
  const response = await retryFetcher(`/header-nav-items`, {
    method: 'POST',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function deleteHeaderNavItem(
  id: string,
): Promise<IApiResponse<components['schemas']['OperationResponse']>> {
  const response = await retryFetcher(`/header-nav-items/${id}`, {
    method: 'DELETE',
  });
  return response.data;
}

async function createFooterNavItem(
  item: components['schemas']['FooterNavItemCreateDto'],
): Promise<IApiResponse<components['schemas']['OperationResponse']>> {
  const response = await retryFetcher(`/footer-nav-items`, {
    method: 'POST',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function deleteFooterNavItem(
  id: string,
): Promise<IApiResponse<components['schemas']['OperationResponse']>> {
  const response = await retryFetcher(`/footer-nav-items/${id}`, {
    method: 'DELETE',
  });
  return response.data;
}

export {
  getAppInfo,
  updateApp,
  getHeaderNavItems,
  updateHeaderNavItem,
  createHeaderNavItem,
  deleteHeaderNavItem,
  getFooterNavItems,
  updateFooterNavItem,
  createFooterNavItem,
  deleteFooterNavItem,
};
