import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar';
import { Badge } from '../ui/badge';
import { Button, buttonVariants } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
  CardFooter,
} from '@/components/ui/card';
import { Check, Linkedin } from 'lucide-react';
import { LightBulbIcon } from '../Icons';
import { GitHubLogoIcon } from '@radix-ui/react-icons';
import Link from 'next/link';

export const HeroCards = () => {
  return (
    <div className="hidden lg:flex flex-row gap-8 w-[700px] h-[500px]">
      <div className="flex flex-col items-end gap-8">
        {/* Testimonial */}
        <Card className="w-[340px] drop-shadow-xl shadow-black/10 dark:shadow-white/10 gap-2">
          <CardHeader className="flex flex-row items-center gap-4">
            <Avatar className="size-10">
              <AvatarImage alt="" src="https://github.com/shadcn.png" />
              <AvatarFallback>SH</AvatarFallback>
            </Avatar>

            <div className="flex flex-col">
              <CardTitle className="text-lg">Nguyễn Văn An</CardTitle>
              <CardDescription>@nguyen_van_an</CardDescription>
            </div>
          </CardHeader>

          <CardContent>
            Website tuyệt vời! Giúp tôi hiểu rõ hơn về lịch sử dân tộc.
          </CardContent>
        </Card>

        {/* Pricing */}
        <Card className="w-72 drop-shadow-xl shadow-black/10 dark:shadow-white/10 gap-4">
          <CardHeader>
            <CardTitle className="flex item-center justify-between">
              Miễn phí
              <Badge variant="secondary" className="text-sm">
                Phổ biến nhất
              </Badge>
            </CardTitle>
            <div>
              <span className="text-3xl font-bold">Miễn phí</span>
            </div>

            <CardDescription>
              Truy cập đầy đủ nội dung lịch sử Việt Nam hoàn toàn miễn phí
            </CardDescription>
          </CardHeader>

          <CardContent>
            <Button className="w-full">Bắt đầu khám phá</Button>
          </CardContent>

          <hr className="w-4/5 m-auto" />

          <CardFooter className="flex">
            <div className="space-y-4">
              {[
                'Xem dòng thời gian',
                'Tra cứu sự kiện',
                'Thư viện nhân vật',
              ].map((benefit: string) => (
                <span key={benefit} className="flex">
                  <Check className="text-green-500" />{' '}
                  <h3 className="ml-2">{benefit}</h3>
                </span>
              ))}
            </div>
          </CardFooter>
        </Card>
      </div>

      <div className="mt-8 flex flex-col items-start gap-8">
        {/* Team */}
        <Card className="w-80 flex flex-col justify-center items-center drop-shadow-xl shadow-black/10 dark:shadow-white/10 gap-2">
          <CardHeader className="w-full mt-8 flex justify-center flex-col items-center pb-2">
            <img
              src="https://i.pravatar.cc/150?img=58"
              alt="user avatar"
              className="absolute grayscale-[0%] -top-12 rounded-full w-24 h-24 aspect-square object-cover"
            />
            <CardTitle className="text-center">TS. Nguyễn Văn Minh</CardTitle>
            <CardDescription className="font-normal text-secondary">
              Chuyên gia Lịch sử Việt Nam
            </CardDescription>
          </CardHeader>

          <CardContent className="text-center pb-2">
            <p>
              Tôi tận tâm với sứ mệnh bảo tồn và lan tỏa kiến thức lịch sử Việt
              Nam đến thế hệ trẻ
            </p>
          </CardContent>
        </Card>

        {/* Service */}
        <Card className="w-[350px] drop-shadow-xl shadow-black/10 dark:shadow-white/10">
          <CardHeader className="space-y-1 flex md:flex-row justify-start items-start gap-4">
            <div className="mt-1 bg-primary/20 p-1 rounded-2xl">
              <LightBulbIcon />
            </div>
            <div>
              <CardTitle>Bản đồ lịch sử tương tác</CardTitle>
              <CardDescription className="text-md mt-2">
                Khám phá các sự kiện lịch sử trên bản đồ Việt Nam với giao diện
                trực quan và dễ sử dụng.
              </CardDescription>
            </div>
          </CardHeader>
        </Card>
      </div>
    </div>
  );
};
