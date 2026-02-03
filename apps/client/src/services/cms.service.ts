'use server';

import { components } from '@nsv-interfaces/cms-service';
import { retryFetcher } from '.';
import { IApiResponse } from '@/interfaces/response.interface';

async function getAppInfo(
  query?: Record<string, string> | string,
): Promise<IApiResponse<components['schemas']['AppData']>> {
  const response = await retryFetcher(
    `/app?${new URLSearchParams(query).toString()}`,
  );

  return response;
}

async function updateApp(
  app: Partial<components['schemas']['AppUpdateReq']>,
): Promise<IApiResponse<components['schemas']['OperationResult']>> {
  const response = await retryFetcher(`/app`, {
    method: 'PUT',
    body: JSON.stringify(app),
  });
  return response.data;
}

async function getHeaderNavItems(): Promise<
  IApiResponse<Array<components['schemas']['HeaderNavItemData']>>
> {
  const response = await retryFetcher(`/header-nav-items`);

  return response;
}

async function updateHeaderNavItem(
  id: string,
  item: Partial<components['schemas']['HeaderNavItemUpdateReq']>,
): Promise<IApiResponse<components['schemas']['OperationResult']>> {
  const response = await retryFetcher(`/header-nav-items/${id}`, {
    method: 'PUT',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function getFooterNavItems(): Promise<
  IApiResponse<Array<components['schemas']['FooterNavItemData']>>
> {
  const response = await retryFetcher(`/footer-nav-items`);

  return response;
}

async function updateFooterNavItem(
  id: string,
  item: Partial<components['schemas']['FooterNavItemUpdateReq']>,
): Promise<IApiResponse<components['schemas']['OperationResult']>> {
  const response = await retryFetcher(`/footer-nav-items/${id}`, {
    method: 'PUT',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function createHeaderNavItem(
  item: components['schemas']['HeaderNavItemCreateReq'],
): Promise<IApiResponse<components['schemas']['HeaderNavItemCreateReq']>> {
  const response = await retryFetcher(`/header-nav-items`, {
    method: 'POST',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function deleteHeaderNavItem(
  id: string,
): Promise<IApiResponse<components['schemas']['OperationResult']>> {
  const response = await retryFetcher(`/header-nav-items/${id}`, {
    method: 'DELETE',
  });
  return response.data;
}

async function createFooterNavItem(
  item: components['schemas']['FooterNavItemCreateReq'],
): Promise<IApiResponse<components['schemas']['FooterNavItemCreateReq']>> {
  const response = await retryFetcher(`/footer-nav-items`, {
    method: 'POST',
    body: JSON.stringify(item),
  });
  return response.data;
}

async function deleteFooterNavItem(
  id: string,
): Promise<IApiResponse<components['schemas']['OperationResult']>> {
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
