import { PrismaClient } from '@historical-event-prisma';
import { GetAllHistoricalEventsRequest } from '@phanhotboy/genproto/historical_event_service/historical_events';

function requestToListQuery(
  query: GetAllHistoricalEventsRequest,
): Parameters<PrismaClient['historicalEvent']['findMany']>[0] {
  const {
    categoryIds,
    fromDay,
    fromMonth,
    fromYear,
    toDay,
    toMonth,
    toYear,
    sortOrder = 'desc',
    sortBy = 'fromYear',
    page = 1,
    limit = 10,
  } = query;

  const options = {
    where: {} as any,
    skip: (page - 1) * limit,
    take: limit,
    orderBy: {
      [sortBy]: sortOrder,
    },
  };

  if (categoryIds && categoryIds.length > 0) {
    options.where.categories = {
      some: { categoryId: { in: categoryIds } },
    };
  }

  const hasFromYear = fromYear !== undefined;
  const hasFromMonth = fromMonth !== undefined;
  const hasFromDay = fromDay !== undefined;
  const hasToYear = toYear !== undefined;
  const hasToMonth = toMonth !== undefined;
  const hasToDay = toDay !== undefined;

  if (hasFromYear) {
    options.where.fromYear = fromYear;
  }

  if (hasFromMonth && hasFromYear) {
    options.where.fromMonth = fromMonth;
  }

  if (hasFromDay && hasFromMonth && hasFromYear) {
    options.where.fromDay = fromDay;
  }
  if (hasToYear) {
    options.where.toYear = toYear;
  }

  if (hasToMonth && hasToYear) {
    options.where.toMonth = toMonth;
  }

  if (hasToDay && hasToMonth && hasToYear) {
    options.where.toDay = toDay;
  }

  return options;
}

function uniqueQueryBuilder(
  id: string,
  authorId?: string,
): Parameters<PrismaClient['historicalEvent']['findUnique']>[0] {
  return {
    where: { id, authorId },
    include: {
      author: true,
      categories: { include: { category: true, event: false } },
    },
  };
}

export { requestToListQuery, uniqueQueryBuilder };
