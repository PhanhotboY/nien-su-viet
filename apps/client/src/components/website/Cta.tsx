import { Button } from '../ui/button';

export const Cta = () => {
  return (
    <section id="cta" className="bg-muted/50 py-16 my-24 sm:my-32">
      <div className="container lg:grid lg:grid-cols-2 place-items-center">
        <div className="lg:col-start-1">
          <h2 className="text-3xl md:text-4xl font-bold ">
            Khám Phá
            <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
              {' '}
              Lịch Sử Việt Nam{' '}
            </span>
            Ngay Hôm Nay
          </h2>
          <p className="text-muted-foreground text-xl mt-4 mb-8 lg:mb-0">
            Tham gia cùng hàng ngàn người dùng đang khám phá và học hỏi về lịch
            sử dân tộc. Trải nghiệm miễn phí ngay hôm nay!
          </p>
        </div>

        <div className="space-y-4 lg:col-start-2">
          <Button className="w-full md:mr-4 md:w-auto">Bắt đầu khám phá</Button>
          <Button variant="outline" className="w-full md:w-auto">
            Xem thêm tính năng
          </Button>
        </div>
      </div>
    </section>
  );
};
