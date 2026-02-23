import { HISTORICAL_EVENT } from '@/constants/historical-event.constant';
import { components } from '@nsv-interfaces/nsv-api-documentation';

export async function getImportantEvents(): Promise<
  (Omit<
    components['schemas']['HistoricalEventBriefResponseDto'],
    'author' | 'id'
  > & {
    description: string;
    highlights: string[];
  })[]
> {
  return [
    {
      name: 'Thời kỳ Hồng Bàng',
      fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE,
      fromYear: -2879,
      toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE,
      toYear: -258,
      description: 'Khởi nguồn dựng nước và giữ nước của tổ tiên',
      highlights: ['18 đời Vua Hùng', 'Văn Lang quốc', 'Nền văn hóa Đồng Đào'],
    },
    {
      name: 'Các triều đại phong kiến',
      fromYear: -968,
      fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
      toYear: 1945,
      toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
      description: 'Thời kỳ độc lập, tự chủ và phát triển văn hóa dân tộc',
      highlights: [
        'Nhà Lý - Trần - Lê',
        'Đánh bại quân Mông Cổ',
        'Hội nhập Đông Nam Á',
      ],
    },
    {
      name: 'Kháng chiến chống Pháp',
      fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
      fromYear: 1945,
      fromMonth: 8,
      fromDay: 19,
      toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
      toYear: 1954,
      toMonth: 7,
      toDay: 21,
      description: 'Giành độc lập, tự do cho dân tộc',
      highlights: [
        'Cách mạng tháng Tám',
        'Chiến thắng Điện Biên Phủ',
        'Độc lập dân tộc',
      ],
    },
    {
      name: 'Thống nhất đất nước',
      fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
      fromYear: 1954,
      fromMonth: 7,
      fromDay: 21,
      toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
      toYear: 1975,
      toMonth: 4,
      toDay: 30,
      description: 'Hoàn thành thống nhất non sông',
      highlights: [
        'Kháng chiến chống Mỹ',
        'Giải phóng miền Nam',
        'Thống nhất Tổ quốc',
      ],
    },
  ];
}
