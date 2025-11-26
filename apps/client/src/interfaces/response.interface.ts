export interface IPaginatedResponse<T> {
  items: T[];
  total: number;
  totalPages: number;
  page: number;
  limit: number;
}
