import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion';
import Link from 'next/link';

interface FAQProps {
  question: string;
  answer: string;
  value: string;
}

const FAQList: FAQProps[] = [
  {
    question: 'Niên Sử Việt có miễn phí không?',
    answer:
      'Có. Tất cả nội dung trên Niên Sử Việt hoàn toàn miễn phí. Chúng tôi cam kết mang kiến thức lịch sử đến với mọi người mà không thu bất kỳ chi phí nào.',
    value: 'item-1',
  },
  {
    question: 'Làm thế nào để tìm kiếm sự kiện lịch sử trên website?',
    answer:
      'Bạn có thể sử dụng thanh tìm kiếm ở đầu trang để tra cứu theo tên sự kiện, nhân vật, hoặc thời gian. Ngoài ra, bạn có thể duyệt qua dòng thời gian tương tác hoặc lọc theo triều đại, chủ đề lịch sử.',
    value: 'item-2',
  },
  {
    question: 'Nguồn tài liệu trên Niên Sử Việt có đáng tin cậy không?',
    answer:
      'Tất cả thông tin trên Niên Sử Việt được sưu tầm, biên soạn và kiểm chứng kỹ lưỡng từ các nguồn uy tín như Viện Sử học Việt Nam, Bảo tàng Lịch sử Quốc gia, và các tài liệu học thuật đã được công nhận.',
    value: 'item-3',
  },
  {
    question: 'Tôi có thể đóng góp nội dung cho Niên Sử Việt không?',
    answer:
      'Có. Chúng tôi luôn hoan nghênh sự đóng góp từ cộng đồng. Bạn có thể gửi tài liệu, hình ảnh, hoặc thông tin bổ sung qua mục "Liên hệ". Mọi đóng góp sẽ được kiểm duyệt trước khi đăng tải.',
    value: 'item-4',
  },
  {
    question: 'Website có ứng dụng di động không?',
    answer:
      'Hiện tại Niên Sử Việt được tối ưu hóa cho mọi thiết bị (responsive design), bạn có thể truy cập dễ dàng trên điện thoại, máy tính bảng và máy tính. Chúng tôi đang phát triển ứng dụng mobile để mang đến trải nghiệm tốt hơn.',
    value: 'item-5',
  },
];

export const FAQ = () => {
  return (
    <section id="faq" className="container py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold mb-4">
        Câu Hỏi{' '}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          Thường Gặp
        </span>
      </h2>

      <Accordion type="single" collapsible className="w-full AccordionRoot">
        {FAQList.map(({ question, answer, value }: FAQProps) => (
          <AccordionItem key={value} value={value} className="!border-b-1">
            <AccordionTrigger className="text-left text-base">
              {question}
            </AccordionTrigger>

            <AccordionContent className="text-base">{answer}</AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>

      <h3 className="font-medium mt-4">
        Vẫn còn thắc mắc?{' '}
        <Link
          rel="noreferrer noopener"
          href="#"
          className="text-secondary transition-all border-primary hover:border-b-2"
        >
          Liên hệ với chúng tôi
        </Link>
      </h3>
    </section>
  );
};
