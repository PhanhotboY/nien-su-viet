import { HISTORICAL_EVENT } from 'libs/nsv-common/src/constants/historical-event.constant';
import { HistoricalEventBaseCreateDto } from '../src/modules/historical-event/dto';

export const mockEvents: HistoricalEventBaseCreateDto[] = [
  {
    name: 'Khởi nghĩa Hai Bà Trưng',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 40,
    fromMonth: 3,
    toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    toYear: 43,
    content:
      'Hai Bà Trưng phất cờ khởi nghĩa chống ách đô hộ của nhà Đông Hán, giành lại quyền tự chủ cho người Việt và lập nên nền độc lập ngắn ngủi trước khi hy sinh anh dũng.',
  },
  {
    name: 'Chiến thắng Bạch Đằng của Ngô Quyền',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 938,
    fromMonth: 12,
    toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    toYear: 938,
    toMonth: 12,
    toDay: 25,
    content:
      'Ngô Quyền đánh bại quân Nam Hán trên sông Bạch Đằng bằng trận cọc gỗ nổi tiếng, chấm dứt hơn một thiên niên kỷ Bắc thuộc và mở ra kỷ nguyên độc lập lâu dài cho dân tộc.',
  },
  {
    name: 'Chiến dịch Điện Biên Phủ',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1954,
    fromMonth: 3,
    fromDay: 13,
    toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    toYear: 1954,
    toMonth: 5,
    toDay: 7,
    content:
      'Quân đội Nhân dân Việt Nam dưới sự chỉ huy của Đại tướng Võ Nguyên Giáp bao vây và tiêu diệt tập đoàn cứ điểm Điện Biên Phủ, buộc Pháp ký Hiệp định Genève, chấm dứt chiến tranh ở Đông Dương.',
  },
  {
    name: 'Tuyên ngôn Độc lập 2/9/1945',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1945,
    fromMonth: 9,
    fromDay: 2,
    toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    toYear: 1945,
    toMonth: 9,
    toDay: 2,
    content:
      'Chủ tịch Hồ Chí Minh đọc Tuyên ngôn Độc lập tại Quảng trường Ba Đình, khai sinh nước Việt Nam Dân chủ Cộng hòa, khẳng định quyền tự do và độc lập của dân tộc Việt Nam.',
  },
  {
    name: 'Ngày Giải phóng miền Nam 30/4/1975',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1975,
    fromMonth: 4,
    fromDay: 30,
    toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    toYear: 1975,
    toMonth: 4,
    toDay: 30,
    content:
      'Chiến dịch Hồ Chí Minh toàn thắng, quân giải phóng tiến vào Sài Gòn, đánh dấu sự sụp đổ của chính quyền Sài Gòn và mở ra thời kỳ thống nhất đất nước.',
  },
  {
    name: 'Lý Công Uẩn dời đô ra Thăng Long',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1010,
    content:
      'Vua Lý Công Uẩn ban Chiếu dời đô, chuyển kinh đô từ Hoa Lư ra thành Đại La, đổi tên thành Thăng Long, mở ra thời kỳ phát triển thịnh trị của kinh đô mới.',
  },
  {
    name: 'Chiến thắng Đông Bộ Đầu',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1258,
    content:
      'Quân Trần đánh bại cuộc xâm lược lần thứ nhất của quân Nguyên Mông tại bến Đông Bộ Đầu, bảo vệ vững chắc kinh đô Thăng Long.',
  },
  {
    name: 'Trận Bạch Đằng thời Trần',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1288,
    fromMonth: 4,
    fromDay: 9,
    content:
      'Hưng Đạo Đại Vương Trần Quốc Tuấn dùng trận địa cọc và hỏa công tiêu diệt đạo thủy quân Nguyên, kết thúc cuộc kháng chiến chống Nguyên lần thứ ba.',
  },
  {
    name: 'Trận Chi Lăng',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1427,
    content:
      'Nghĩa quân Lam Sơn đánh bại viện binh nhà Minh tại ải Chi Lăng, góp phần quyết định buộc quân Minh phải rút khỏi Đại Việt.',
  },
  {
    name: 'Khởi nghĩa Yên Thế',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1884,
    content:
      'Nghĩa quân nông dân do Hoàng Hoa Thám lãnh đạo nổi dậy chống thực dân Pháp ở vùng Yên Thế, kéo dài suốt gần 30 năm.',
  },
  {
    name: 'Thành lập Đảng Cộng sản Việt Nam',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1930,
    fromMonth: 2,
    fromDay: 3,
    content:
      'Hội nghị hợp nhất các tổ chức cộng sản tại Cửu Long (Hồng Kông) tuyên bố thành lập Đảng Cộng sản Việt Nam, mở ra bước ngoặt lãnh đạo cách mạng.',
  },
  {
    name: 'Tổng tiến công và nổi dậy Tết Mậu Thân',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1968,
    fromMonth: 1,
    fromDay: 30,
    content:
      'Quân Giải phóng và lực lượng cách mạng đồng loạt tiến công vào hầu hết các đô thị miền Nam, gây chấn động mạnh và làm thay đổi cục diện chiến tranh Việt Nam.',
  },
  {
    name: 'Đại tướng Võ Nguyên Giáp chỉ huy “Điện Biên Phủ trên không”',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1972,
    fromMonth: 12,
    fromDay: 18,
    toDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    toYear: 1972,
    toMonth: 12,
    toDay: 29,
    content:
      'Quân và dân miền Bắc đánh bại cuộc tập kích chiến lược bằng B-52 của Mỹ, bảo vệ thủ đô Hà Nội và buộc Mỹ phải ký Hiệp định Paris.',
  },
  {
    name: 'Chiến tranh biên giới phía Bắc 1979',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1979,
    fromMonth: 2,
    fromDay: 17,
    content:
      'Quân dân Việt Nam chiến đấu bảo vệ chủ quyền lãnh thổ trước cuộc tiến công của quân đội Trung Quốc dọc biên giới phía Bắc.',
  },
  {
    name: 'Đại hội VI khởi xướng Đổi mới',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 1986,
    fromMonth: 12,
    content:
      'Đại hội Đại biểu toàn quốc lần thứ VI của Đảng Cộng sản Việt Nam đề ra đường lối Đổi mới, chuyển sang phát triển kinh tế thị trường định hướng xã hội chủ nghĩa.',
  },
  {
    name: 'Ký Hiệp định Thương mại song phương Việt - Mỹ (BTA)',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 2000,
    fromMonth: 7,
    fromDay: 13,
    content:
      'Việt Nam và Hoa Kỳ ký Hiệp định Thương mại song phương, mở ra giai đoạn hội nhập kinh tế sâu rộng với thị trường Mỹ.',
  },
  {
    name: 'Hiệp định Thương mại Tự do EU - Việt Nam (EVFTA) có hiệu lực',
    fromDateType: HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT,
    fromYear: 2020,
    fromMonth: 8,
    fromDay: 1,
    content:
      'EVFTA chính thức có hiệu lực, dỡ bỏ phần lớn dòng thuế giữa Việt Nam và Liên minh châu Âu, thúc đẩy mạnh mẽ xuất nhập khẩu và đầu tư.',
  },
];
