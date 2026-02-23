import { Card, CardDescription, CardHeader, CardTitle } from '../ui/card';
import { MagnifierIcon, WalletIcon, ChartIcon } from '../Icons';
import { JSX } from 'react';
import Image from 'next/image';

interface ServiceProps {
  title: string;
  description: string;
  icon: JSX.Element;
}

const serviceList: ServiceProps[] = [
  {
    title: 'Tra cứu sự kiện',
    description:
      'Tìm kiếm và khám phá các sự kiện lịch sử quan trọng của Việt Nam. Xem chi tiết thời gian, địa điểm, nhân vật liên quan và ý nghĩa lịch sử.',
    icon: <ChartIcon />,
  },
  {
    title: 'Thư viện nhân vật',
    description:
      'Tìm hiểu về các anh hùng dân tộc, vua chúa, danh tướng, văn nhân và các nhân vật tiêu biểu trong lịch sử Việt Nam qua từng thời kỳ.',
    icon: <WalletIcon />,
  },
  {
    title: 'Tài liệu lịch sử',
    description:
      'Truy cập kho tài liệu số hóa với hình ảnh, văn bản cổ, bản đồ lịch sử và các nguồn tài liệu quý giá được sưu tầm và bảo quản.',
    icon: <MagnifierIcon />,
  },
];

export const Services = () => {
  return (
    <section className="container py-24 sm:py-32">
      <div className="grid lg:grid-cols-2 gap-8 place-items-center">
        <div>
          <h2 className="text-3xl md:text-4xl font-bold">
            <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
              Dịch vụ{' '}
            </span>
            Của Chúng Tôi
          </h2>

          <p className="text-muted-foreground text-xl mt-4 mb-8 ">
            Khám phá và tìm hiểu lịch sử Việt Nam một cách toàn diện với các
            dịch vụ chất lượng cao.
          </p>

          <div className="flex flex-col gap-8">
            {serviceList.map(({ icon, title, description }: ServiceProps) => (
              <Card key={title}>
                <CardHeader className="space-y-1 flex md:flex-row justify-start items-start gap-4">
                  <div className="mt-1 bg-primary/20 p-1 rounded-2xl">
                    {icon}
                  </div>
                  <div>
                    <CardTitle>{title}</CardTitle>
                    <CardDescription className="text-md mt-2">
                      {description}
                    </CardDescription>
                  </div>
                </CardHeader>
              </Card>
            ))}
          </div>
        </div>

        {/*<Image
          src={cubeLeg}
          className="w-[300px] md:w-[500px] lg:w-[600px] object-contain"
          alt="About services"
        />*/}
      </div>
    </section>
  );
};
