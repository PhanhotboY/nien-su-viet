import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Check } from 'lucide-react';

enum PopularPlanType {
  NO = 0,
  YES = 1,
}

interface PricingProps {
  title: string;
  popular: PopularPlanType;
  price: number;
  description: string;
  buttonText: string;
  benefitList: string[];
}

const pricingList: PricingProps[] = [
  {
    title: 'Miễn phí',
    popular: 0,
    price: 0,
    description: 'Dành cho người mới bắt đầu tìm hiểu lịch sử Việt Nam',
    buttonText: 'Bắt đầu ngay',
    benefitList: [
      'Xem dòng thời gian cơ bản',
      'Tra cứu 100+ sự kiện',
      'Đọc tiểu sử nhân vật',
      'Hỗ trợ cộng đồng',
      'Truy cập trên mọi thiết bị',
    ],
  },
  {
    title: 'Cao cấp',
    popular: 1,
    price: 0,
    description: 'Trải nghiệm đầy đủ với tất cả tính năng cao cấp',
    buttonText: 'Dùng thử miễn phí',
    benefitList: [
      'Truy cập toàn bộ nội dung',
      'Tải tài liệu và hình ảnh',
      'Bản đồ lịch sử tương tác',
      'Không quảng cáo',
      'Hỗ trợ ưu tiên 24/7',
    ],
  },
  {
    title: 'Tổ chức',
    popular: 0,
    price: 0,
    description: 'Giải pháp cho trường học, thư viện và tổ chức',
    buttonText: 'Liên hệ',
    benefitList: [
      'Tài khoản không giới hạn',
      'Quản lý tập trung',
      'Tùy chỉnh nội dung',
      'API tích hợp',
      'Đào tạo và hỗ trợ chuyên biệt',
    ],
  },
];

export const Pricing = () => {
  return (
    <section id="pricing" className="container py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold text-center">
        Chọn Gói
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          {' '}
          Phù Hợp{' '}
        </span>
        Với Bạn
      </h2>
      <h3 className="text-xl text-center text-muted-foreground pt-4 pb-8">
        Tất cả các gói đều miễn phí. Chúng tôi cam kết mang kiến thức lịch sử
        đến với mọi người.
      </h3>
      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {pricingList.map((pricing: PricingProps) => (
          <Card
            key={pricing.title}
            className={
              pricing.popular === PopularPlanType.YES
                ? 'drop-shadow-primary/25 drop-shadow-xl shadow-primary/10 dark:shadow-white/10'
                : ''
            }
          >
            <CardHeader>
              <CardTitle className="flex item-center justify-between">
                {pricing.title}
                {pricing.popular === PopularPlanType.YES ? (
                  <Badge variant="secondary">Phổ biến nhất</Badge>
                ) : null}
              </CardTitle>
              <div>
                <span className="text-3xl font-bold">
                  {pricing.price === 0 ? 'Miễn phí' : `${pricing.price}đ`}
                </span>
                {pricing.price > 0 && (
                  <span className="text-muted-foreground"> /tháng</span>
                )}
              </div>

              <CardDescription>{pricing.description}</CardDescription>
            </CardHeader>

            <CardContent>
              <Button className="w-full">{pricing.buttonText}</Button>
            </CardContent>

            <hr className="w-4/5 m-auto mb-4" />

            <CardFooter className="flex">
              <div className="space-y-4">
                {pricing.benefitList.map((benefit: string) => (
                  <span key={benefit} className="flex">
                    <Check className="text-green-500" />{' '}
                    <h3 className="ml-2">{benefit}</h3>
                  </span>
                ))}
              </div>
            </CardFooter>
          </Card>
        ))}
      </div>
    </section>
  );
};
