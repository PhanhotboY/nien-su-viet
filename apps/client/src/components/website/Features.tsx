import { Badge } from '../ui/badge';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import image from '@public/growth.png';
import image3 from '@public/reflecting.png';
import image4 from '@public/looking-ahead.png';
import Image, { StaticImageData } from 'next/image';

interface FeatureProps {
  title: string;
  description: string;
  image: StaticImageData;
}

const features: FeatureProps[] = [
  {
    title: 'Dòng thời gian tương tác',
    description:
      'Khám phá lịch sử Việt Nam qua dòng thời gian trực quan và tương tác. Dễ dàng di chuyển qua các thời kỳ, xem chi tiết sự kiện và nhân vật lịch sử.',
    image: image4,
  },
  {
    title: 'Nội dung phong phú',
    description:
      'Kho tàng kiến thức với hàng ngàn sự kiện, nhân vật, và tài liệu lịch sử được biên soạn và kiểm chứng kỹ lưỡng từ nhiều nguồn uy tín.',
    image: image3,
  },
  {
    title: 'Giao diện thân thiện',
    description:
      'Thiết kế hiện đại, dễ sử dụng trên mọi thiết bị. Hỗ trợ tìm kiếm nhanh, lọc theo thời kỳ, triều đại và chủ đề lịch sử.',
    image: image,
  },
];

const featureList: string[] = [
  'Dòng thời gian',
  'Bản đồ lịch sử',
  'Tra cứu sự kiện',
  'Nhân vật lịch sử',
  'Triều đại Việt Nam',
  'Tài liệu số hóa',
  'Hình ảnh lịch sử',
  'Trắc nghiệm kiến thức',
  'Giao diện đa thiết bị',
];

export const Features = () => {
  return (
    <section id="features" className="container py-24 sm:py-32 space-y-8">
      <h2 className="text-3xl lg:text-4xl font-bold md:text-center">
        Nhiều{' '}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          Tính Năng Nổi Bật
        </span>
      </h2>

      <div className="flex flex-wrap md:justify-center gap-4">
        {featureList.map((feature: string) => (
          <div key={feature}>
            <Badge variant="secondary" className="text-sm">
              {feature}
            </Badge>
          </div>
        ))}
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {features.map(({ title, description, image }: FeatureProps) => (
          <Card key={title}>
            <CardHeader>
              <CardTitle>{title}</CardTitle>
            </CardHeader>

            <CardContent>{description}</CardContent>

            <CardFooter>
              <Image
                src={image}
                alt="About feature"
                className="w-[200px] lg:w-[300px] mx-auto"
              />
            </CardFooter>
          </Card>
        ))}
      </div>
    </section>
  );
};
