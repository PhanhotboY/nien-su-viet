import { Statistics } from './Statistics';
import Image from 'next/image';

export const About = () => {
  return (
    <section id="about" className="container py-24 sm:py-32">
      <div className="bg-muted/50 border rounded-lg py-12">
        <div className="px-6 flex flex-col-reverse md:flex-row gap-8 md:gap-12">
          {/*<Image
            src={pilot}
            alt=""
            className="w-[300px] object-contain rounded-lg"
          />*/}
          <div className="bg-green-0 flex flex-col justify-between">
            <div className="pb-6">
              <h2 className="text-3xl md:text-4xl font-bold">
                <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
                  Về{' '}
                </span>
                Niên Sử Việt
              </h2>
              <p className="text-xl text-muted-foreground mt-4">
                Niên Sử Việt là nền tảng số hóa lịch sử Việt Nam, mang đến cho
                bạn cái nhìn toàn diện và sinh động về hành trình phát triển của
                dân tộc. Chúng tôi tập hợp và sắp xếp các sự kiện lịch sử theo
                thời gian, giúp bạn dễ dàng theo dõi và hiểu rõ dòng chảy lịch
                sử từ thời Hùng Vương đến thời kỳ đổi mới. Với sứ mệnh bảo tồn
                và lan tỏa giá trị lịch sử, văn hóa Việt Nam đến thế hệ trẻ và
                cộng đồng quốc tế.
              </p>
            </div>

            <Statistics />
          </div>
        </div>
      </div>
    </section>
  );
};
