import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';

interface TestimonialProps {
  image: string;
  name: string;
  userName: string;
  comment: string;
}

const testimonials: TestimonialProps[] = [
  {
    image: 'https://github.com/shadcn.png',
    name: 'Nguyễn Văn An',
    userName: '@nguyen_van_an',
    comment:
      'Website tuyệt vời! Giúp tôi hiểu rõ hơn về lịch sử dân tộc. Giao diện đẹp và dễ sử dụng.',
  },
  {
    image: 'https://github.com/shadcn.png',
    name: 'Trần Thị Mai',
    userName: '@tran_thi_mai',
    comment:
      'Là giáo viên lịch sử, tôi thường xuyên sử dụng Niên Sử Việt để chuẩn bị bài giảng. Nội dung chính xác và phong phú, rất hữu ích cho công việc giảng dạy.',
  },

  {
    image: 'https://github.com/shadcn.png',
    name: 'Lê Hoàng Nam',
    userName: '@le_hoang_nam',
    comment:
      'Dòng thời gian tương tác rất ấn tượng! Tôi có thể dễ dàng tìm kiếm và theo dõi các sự kiện lịch sử từ thời Hùng Vương đến hiện đại. Đây là nguồn tài liệu tuyệt vời.',
  },
  {
    image: 'https://github.com/shadcn.png',
    name: 'Phạm Minh Châu',
    userName: '@pham_minh_chau',
    comment:
      'Con tôi học lớp 10 và rất thích lịch sử nhờ website này. Các hình ảnh và tài liệu giúp em học một cách sinh động hơn rất nhiều so với sách giáo khoa.',
  },
  {
    image: 'https://github.com/shadcn.png',
    name: 'Võ Thành Công',
    userName: '@vo_thanh_cong',
    comment:
      'Là người yêu thích lịch sử, tôi đã tìm thấy nhiều tài liệu quý giá trên Niên Sử Việt. Đặc biệt là phần bản đồ lịch sử rất trực quan và dễ hiểu.',
  },
  {
    image: 'https://github.com/shadcn.png',
    name: 'Hoàng Thị Lan',
    userName: '@hoang_thi_lan',
    comment:
      'Website được thiết kế rất chuyên nghiệp và khoa học. Tôi đã giới thiệu cho nhiều bạn bè cùng tìm hiểu về lịch sử Việt Nam qua trang web này.',
  },
];

export const Testimonials = () => {
  return (
    <section id="testimonials" className="container py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold">
        Khám Phá Lý Do
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {' '}
          Mọi Người Yêu Thích{' '}
        </span>
        Niên Sử Việt
      </h2>

      <p className="text-xl text-muted-foreground pt-4 pb-8">
        Hàng ngàn người dùng đã tin tưởng và sử dụng Niên Sử Việt để khám phá và
        học hỏi về lịch sử dân tộc
      </p>

      <div className="grid md:grid-cols-2 lg:grid-cols-4 sm:block columns-2 lg:columns-3 lg:gap-6 mx-auto space-y-4 lg:space-y-6">
        {testimonials.map(
          ({ image, name, userName, comment }: TestimonialProps) => (
            <Card
              key={userName}
              className="md:break-inside-avoid overflow-hidden gap-2"
            >
              <CardHeader className="flex flex-row items-center gap-4">
                <Avatar className="size-10">
                  <AvatarImage alt="" src={image} />
                  <AvatarFallback>OM</AvatarFallback>
                </Avatar>

                <div className="flex flex-col">
                  <CardTitle className="text-lg">{name}</CardTitle>
                  <CardDescription>{userName}</CardDescription>
                </div>
              </CardHeader>

              <CardContent>{comment}</CardContent>
            </Card>
          ),
        )}
      </div>
    </section>
  );
};
