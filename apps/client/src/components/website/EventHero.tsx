import { Sparkles, Clock, MapPin, ChevronDown } from 'lucide-react';

export const EventHero = () => {
  return (
    <section className="relative overflow-hidden py-12 md:py-16 pb-20">
      {/* Background gradient effect */}
      <div className="absolute inset-0 bg-gradient-to-br from-primary/5 via-secondary/5 to-background -z-10" />

      <div className="container mx-auto px-4">
        <div className="text-center space-y-6 max-w-4xl mx-auto">
          {/* Badge */}
          <div className="inline-flex items-center gap-2 px-4 py-2 bg-primary/10 rounded-full text-sm font-medium text-primary">
            <Sparkles className="h-4 w-4" />
            <span>Khám phá dòng chảy lịch sử</span>
          </div>

          {/* Main heading */}
          <h1 className="text-4xl md:text-6xl font-bold leading-tight">
            <span className="inline-block bg-gradient-to-r from-primary via-primary/80 to-secondary text-transparent bg-clip-text">
              Hành trình xuyên suốt
            </span>
            <br />
            <span className="text-foreground">4000 năm lịch sử Việt Nam</span>
          </h1>

          {/* Description */}
          <p className="text-lg md:text-xl text-muted-foreground max-w-2xl mx-auto leading-relaxed">
            Khám phá những sự kiện quan trọng đã định hình nên lịch sử dân tộc.
            Từ thời Hồng Bàng đến hiện đại, mỗi sự kiện là một mảnh ghép trong
            bức tranh hào hùng của Việt Nam.
          </p>

          {/* Feature highlights */}
          <div className="flex flex-wrap justify-center gap-6 md:gap-8 pt-4">
            <div className="flex items-center gap-2 text-sm md:text-base">
              <div className="p-2 rounded-lg bg-primary/10">
                <Clock className="h-5 w-5 text-primary" />
              </div>
              <div className="text-left">
                <div className="font-semibold">Dòng thời gian tương tác</div>
                <div className="text-xs text-muted-foreground">
                  Dễ dàng điều hướng
                </div>
              </div>
            </div>

            <div className="flex items-center gap-2 text-sm md:text-base">
              <div className="p-2 rounded-lg bg-secondary/10">
                <MapPin className="h-5 w-5 text-secondary" />
              </div>
              <div className="text-left">
                <div className="font-semibold">1000+ Sự kiện</div>
                <div className="text-xs text-muted-foreground">
                  Được kiểm chứng
                </div>
              </div>
            </div>

            <div className="flex items-center gap-2 text-sm md:text-base">
              <div className="p-2 rounded-lg bg-primary/10">
                <Sparkles className="h-5 w-5 text-primary" />
              </div>
              <div className="text-left">
                <div className="font-semibold">Nội dung chi tiết</div>
                <div className="text-xs text-muted-foreground">
                  Dễ hiểu, sinh động
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Decorative elements */}
      <div className="absolute top-1/4 left-10 w-72 h-72 bg-primary/5 rounded-full blur-3xl -z-10 animate-pulse" />
      <div
        className="absolute bottom-1/4 right-10 w-96 h-96 bg-secondary/5 rounded-full blur-3xl -z-10 animate-pulse"
        style={{ animationDelay: '1s' }}
      />

      {/* Scroll indicator */}
      <div className="absolute bottom-4 left-1/2 -translate-x-1/2 flex flex-col items-center gap-2 animate-bounce">
        <span className="text-xs text-muted-foreground font-medium">
          Cuộn xuống để khám phá
        </span>
        <ChevronDown className="h-5 w-5 text-primary" />
      </div>
    </section>
  );
};
