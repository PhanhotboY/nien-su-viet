// Helper function to create date from parts
const createDate = (
  year: number,
  month?: number | null,
  day?: number | null,
) => {
  // Handle BCE dates (negative years)
  const isBCE = year < 0;
  const absYear = Math.abs(year);

  if (isBCE) {
    // For BCE, create date and adjust
    const date = new Date(0, (month || 1) - 1, day || 1);
    date.setFullYear(-absYear);
    return date;
  }

  return new Date(absYear, (month || 1) - 1, day || 1);
};

const formatHistoricalEventDate = (
  year?: number | null,
  month?: number | null,
  day?: number | null,
) => {
  const hasDay = !!(year && month && day);
  const hasMonth = !!(year && month);
  const dateStr = year
    ? `${hasDay ? `${day}/` : ''}${hasMonth ? `${month}/` : ''}${Math.abs(year)}`
    : '';

  return year && year < 0 ? `${dateStr} TCN` : dateStr;
};

export { formatHistoricalEventDate, createDate };
