const { writeFileSync, readFileSync } = require('fs');
const { parse } = require('node-html-parser');

/**
 * name: 'Khởi nghĩa Hai Bà Trưng',
     fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
     fromYear: 40,
     fromMonth: 3,
     toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
     toYear: 43,
     content:
       'Hai Bà Trưng phất cờ khởi nghĩa chống ách đô hộ của nhà Đông Hán, giành lại quyền tự chủ cho người Việt và lập nên nền độc lập ngắn ngủi trước khi hy sinh anh dũng.',
 */
const dateType = {
  fromDateType: 'EXACT',
  toDateType: 'EXACT',
};
function main() {
  const htmlBuffer = readFileSync('./niensuviet.html');
  const html = Buffer.from(htmlBuffer).toString();
  const root = parse(html);

  const events = root.querySelectorAll(
    '.mw-content-ltr.mw-parser-output > p:not(:first-of-type), .mw-content-ltr.mw-parser-output > dl',
  );
  const fileName = './events.json';
  // clear previous data

  const eventsArr = [];
  for (const idx in events) {
    const event = events[idx];
    if (event.tagName === 'DL') continue;

    if (event.tagName === 'P') {
      const anchors = event.querySelectorAll('a');
      if (anchors.length === 0) {
        // year row
        const year = event.innerText.replace('\n', '').trim();
        const eventList = events[+idx + 1].querySelectorAll('dd');
        eventList.entries().forEach(([_, event], i) => {
          const links = event.querySelectorAll('a');
          const { year: date, name } = extractContent(event);
          const parsedDate = extractDate(`${date}, ${year}`);
          const content = serializeContent(links);
          eventsArr.push({
            ...parsedDate,
            ...dateType,
            name,
            content,
          });
        });
      } else {
        // event row
        const { year, name } = extractContent(event);
        const parsedDate = extractDate(year);
        if (!parsedDate) continue;
        const content = serializeContent(anchors);
        eventsArr.push({
          ...parsedDate,
          ...dateType,
          name,
          content,
        });
      }
    }
  }
  writeFileSync(fileName, JSON.stringify(eventsArr, null, '\t'));
}

const extractContent = (node) => {
  const year = node.querySelector('b');
  const name = node.innerText
    .replace(year.innerHTML, '')
    .replaceAll('\n', '')
    .replaceAll('"', "'")
    .trim();
  return {
    year: year.innerHTML.replaceAll('\n', '').replace(':', '').trim(),
    name,
  };
};

const serializeContent = (nodes = []) => {
  const wikiEndpoint = 'https://vi.wikipedia.org';
  const addWikiEndpoint = (a) =>
    a.toString().replace('href="', `href="${wikiEndpoint}`);

  return nodes.reduce(
    (p, c, i) => (i == 0 ? addWikiEndpoint(c) : `${p}, ${addWikiEndpoint(c)}`),
    '',
  );
};

const parseYear = (yearStr = '') => {
  const yearNumber = +yearStr.replace('TCN', '').replace('.', '').trim();
  return yearStr.includes('TCN') ? -yearNumber : yearNumber;
};

const parseDate = (dateStr = '') => {
  /**
   * 1. 23 tháng 6
   * 2. Tháng 5
   * 3. Tháng 1 - Tháng 4
   * 4. 18 tháng 12 - 30 tháng 12
   */
  if (dateStr.includes('-')) {
    const [fromDate, toDate] = dateStr.split('-');
    const { fromMonth, fromDay } = parseDate(fromDate);
    const { fromMonth: toMonth, fromDay: toDay } = parseDate(toDate);
    return { fromMonth, fromDay, toMonth, toDay };
  }

  if (dateStr.includes('tháng')) {
    const [day, month] = dateStr.split('tháng');
    return {
      fromMonth: parseInt(month),
      fromDay: parseInt(day),
    };
  }

  if (dateStr.includes('Tháng')) {
    const month = dateStr.replace('Tháng', '');
    return {
      fromMonth: parseInt(month),
    };
  }
};

const extractDate = (dateStr = '') => {
  /**
   * 1. 23 tháng 6, 229
   * 2. Tháng 5, 618
   * 3. Tháng 1 - Tháng 4, 981
   * 4. 18 tháng 12 - 30 tháng 12, 1972
   * 5. 1–630
   * 6. 25.000 TCN–7.000 TCN
   * 7. 23.000 TCN
   * 8. 40
   */

  // case 7,8
  const parsedYear = parseYear(dateStr);
  if (!isNaN(parsedYear)) {
    return {
      fromYear: parsedYear,
    };
  }

  const [splitedDate, splitedYear] = dateStr.split(', ');
  // case 1,2,3,4
  if (splitedYear) {
    const year = parseYear(splitedYear);
    const { fromMonth, fromDay, toMonth, toDay } = parseDate(splitedDate);
    return {
      fromYear: year,
      fromMonth,
      fromDay,
      toYear: dateStr.includes('-') ? year : undefined,
      toMonth,
      toDay,
    };
  }

  // case 5,6
  let [fromYear, toYear] = dateStr.includes('–')
    ? dateStr.split('–')
    : dateStr.split('-');
  fromYear = parseYear(fromYear);
  toYear = parseYear(toYear);
  if (!isNaN(fromYear) && !isNaN(toYear)) {
    return {
      fromYear,
      toYear,
    };
  }

  // no matching cases
  return null;
};

main();
