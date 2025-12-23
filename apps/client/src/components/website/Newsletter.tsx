'use client';

import { Button } from '../ui/button';
import { Input } from '../ui/input';

export const Newsletter = () => {
  const handleSubmit = (e: any) => {
    e.preventDefault();
    console.log('Subscribed!');
  };

  return (
    <section id="newsletter">
      <hr className="w-11/12 mx-auto" />

      <div className="container py-24 sm:py-32">
        <h3 className="text-center text-4xl md:text-5xl font-bold">
          Đăng Ký Nhận{' '}
          <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
            Bản Tin
          </span>
        </h3>
        <p className="text-xl text-muted-foreground text-center mt-4 mb-8">
          Nhận thông tin về sự kiện lịch sử mỗi ngày và cập nhật nội dung mới
          nhất
        </p>

        <form
          className="flex flex-col w-full md:flex-row md:w-6/12 lg:w-4/12 mx-auto gap-4 md:gap-2"
          onSubmit={handleSubmit}
        >
          <Input
            placeholder="email@example.com"
            className="bg-accent/50 dark:bg-muted/80 "
            aria-label="email"
          />
          <Button>Đăng ký</Button>
        </form>
      </div>

      <hr className="w-11/12 mx-auto" />
    </section>
  );
};
