import { Button } from '../ui/button';
import { buttonVariants } from '../ui/button';
import { HeroCards } from './HeroCards';
import { GitHubLogoIcon } from '@radix-ui/react-icons';
import Link from 'next/link';

export const Hero = () => {
  return (
    <section className="container grid lg:grid-cols-2 place-items-center py-20 md:py-32 gap-10">
      <div className="text-center lg:text-start space-y-6">
        <main className="text-5xl md:text-6xl font-bold">
          <h1 className="inline">
            <span className="inline bg-gradient-to-r from-secondary/50 to-secondary text-transparent bg-clip-text">
              Niên Sử Việt
            </span>{' '}
          </h1>{' '}
          <h2 className="inline">
            <span className="inline bg-gradient-to-r from-primary/80 via-primary to-primary text-transparent bg-clip-text">
              Khám phá lịch sử
            </span>{' '}
            dân tộc Việt Nam
          </h2>
        </main>

        <p className="text-xl text-muted-foreground md:w-10/12 mx-auto lg:mx-0">
          Hành trình xuyên suốt hàng ngàn năm lịch sử hào hùng của dân tộc Việt
          Nam. Từ thời kỳ Hồng Bàng đến hiện đại, khám phá những sự kiện, nhân
          vật và di sản văn hóa quý giá.
        </p>

        <div className="space-y-4 md:space-y-0 md:space-x-4">
          <Button className="w-full md:w-1/3">Khám phá ngay</Button>

          <Link
            rel="noreferrer noopener"
            href="#about"
            className={`w-full md:w-1/3 ${buttonVariants({
              variant: 'outline',
            })}`}
          >
            Tìm hiểu thêm
          </Link>
        </div>
      </div>

      {/* Hero cards sections */}
      <div className="z-10">
        <HeroCards />
      </div>

      {/* Shadow effect */}
      <div className="shadow"></div>
    </section>
  );
};
