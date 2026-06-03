import { HISTORICAL_EVENT } from '@/constants/historical-event.constant';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { useTranslations } from 'next-intl';

// Helper function to create date from parts
const createDate = (
  year: number,
  month?: number | null,
  day?: number | null,
) => {
  // Handle BCE dates (negative years)
  const isBCE = year < 0;
  const absYear = Math.abs(year);

  const date = new Date(year, (month || 1) - 1, day || 1); // Start with a base date
  if (year < 100) {
    // For CE years less than 100, set the full year explicitly
    date.setFullYear(absYear);
  }

  if (isBCE) {
    // For BCE, create date and adjust
    date.setFullYear(-absYear);
  }

  return date;
};

const formatHistoricalEventDate = ({
  dateType,
  sharedTranslator: t,
  year,
  month,
  day,
}: {
  dateType: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>;
  sharedTranslator: ReturnType<typeof useTranslations>;
  year?: number | null;
  month?: number | null;
  day?: number | null;
}) => {
  const hasDay = !!(year && month && day);
  const hasMonth = !!(year && month);
  const dateStr = year
    ? `${hasDay ? `${day}/` : ''}${hasMonth ? `${month}/` : ''}${Math.abs(year)}`
    : '';

  const prefix =
    dateType === HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE
      ? `${t('approximate')} `
      : '';

  return prefix + (year && year < 0 ? `${dateStr} ${t('bce')}` : dateStr);
};

function toEventPeriodString(
  event: Pick<
    components['schemas']['HistoricalEventBriefResponseDto'],
    | 'fromDateType'
    | 'fromYear'
    | 'fromMonth'
    | 'fromDay'
    | 'toDateType'
    | 'toYear'
    | 'toMonth'
    | 'toDay'
  >,
  sharedTranslator: ReturnType<typeof useTranslations>,
) {
  const startDate = formatHistoricalEventDate({
    dateType: event.fromDateType,
    sharedTranslator,
    year: event.fromYear,
    month: event.fromMonth,
    day: event.fromDay,
  });
  const endDate = formatHistoricalEventDate({
    dateType: event.toDateType,
    sharedTranslator,
    year: event.toYear,
    month: event.toMonth,
    day: event.toDay,
  });

  return `${startDate} - ${endDate}`;
}

export { toEventPeriodString, formatHistoricalEventDate, createDate };
