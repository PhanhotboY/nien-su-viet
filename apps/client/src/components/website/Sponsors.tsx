import { Radar } from 'lucide-react';
import { JSX } from 'react';

interface SponsorProps {
  icon: JSX.Element;
  name: string;
}

const sponsors: SponsorProps[] = [
  {
    icon: <Radar size={34} />,
    name: 'Viện Sử học Việt Nam',
  },
  {
    icon: <Radar size={34} />,
    name: 'Bảo tàng Lịch sử Quốc gia',
  },
  {
    icon: <Radar size={34} />,
    name: 'Thư viện Quốc gia',
  },
  {
    icon: <Radar size={34} />,
    name: 'Viện Hán Nôm',
  },
  {
    icon: <Radar size={34} />,
    name: 'Trung tâm Lưu trữ Quốc gia',
  },
  {
    icon: <Radar size={34} />,
    name: 'Bảo tàng Dân tộc học',
  },
];

export const Sponsors = () => {
  return (
    <section id="sponsors" className="container pt-24 sm:py-32">
      <h2 className="text-center text-md lg:text-2xl font-bold mb-8 text-secondary">
        Đối tác và nguồn tham khảo
      </h2>

      <div className="flex flex-wrap justify-center items-center gap-4 md:gap-8">
        {sponsors.map(({ icon, name }: SponsorProps) => (
          <div
            key={name}
            className="flex items-center gap-1 text-muted-foreground/60"
          >
            <span>{icon}</span>
            <h3 className="text-xl  font-bold">{name}</h3>
          </div>
        ))}
      </div>
    </section>
  );
};
