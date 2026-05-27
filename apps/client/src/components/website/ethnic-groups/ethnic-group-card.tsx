'use client';

import Image from 'next/image';
import Link from '@/i18n/navigation';
import { ArrowRight } from 'lucide-react';

interface EthnicGroup {
  code: string;
  name: string;
  slug: string;
  alternative_names: string;
  detail: {
    title?: string;
    other_names?: string;
    population?: string;
    history?: string;
    images?: Array<{
      url: string;
      alt?: string;
    }>;
  };
}

export function EthnicGroupCard({ group }: { group: EthnicGroup }) {
  const firstImage = group.detail?.images?.[0];
  const imageUrl = firstImage?.url || '/assets/image/not-found.webp';
  const imageAlt = firstImage?.alt || group.detail.title || group.name;

  return (
    <Link href={`/dan-toc/${group.slug}`} className="group">
      <div className="relative h-full rounded-xl overflow-hidden bg-card/30 backdrop-blur-sm border border-border/50 hover:border-primary/50 transition-all duration-300 hover:bg-card/50 hover:shadow-lg flex flex-col">
        {/* Image Section */}
        <div className="relative w-full h-48 overflow-hidden bg-muted/50">
          <Image
            src={imageUrl}
            alt={imageAlt}
            fill
            className="object-contain group-hover:scale-105 transition-transform duration-300"
            sizes="(max-width: 640px) 100vw, (max-width: 1024px) 50vw, 33vw"
          />
          <div className="absolute inset-0 bg-linear-to-t from-black/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
        </div>

        {/* Content Section */}
        <div className="relative z-10 flex flex-col flex-1 p-5">
          {/* Gradient accent */}
          <div className="absolute inset-0 bg-linear-to-br from-primary/5 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300 pointer-events-none" />

          {/* Header: Name */}
          <div className="relative z-10 mb-3">
            <h2 className="text-2xl font-bold text-foreground group-hover:text-primary transition-colors duration-300">
              {group.detail.title || group.name}
            </h2>
          </div>

          {/* Other Names */}
          {group.detail?.other_names && (
            <div className="mb-3 pb-3 border-b border-border/30">
              <p className="text-xs font-semibold text-muted-foreground mb-1 uppercase tracking-wide">
                Tên khác
              </p>
              <p className="text-xs text-muted-foreground line-clamp-2 leading-relaxed">
                {group.detail.other_names}
              </p>
            </div>
          )}

          {/* Population */}
          {group.detail?.population && (
            <div className="mb-3 pb-3 border-b border-border/30">
              <p className="text-xs font-semibold text-muted-foreground mb-1 uppercase tracking-wide">
                Dân số
              </p>
              <p className="text-xs text-muted-foreground line-clamp-2 leading-relaxed">
                {group.detail.population}
              </p>
            </div>
          )}

          {/* History */}
          {group.detail?.history && (
            <div className="mb-4">
              <p className="text-xs font-semibold text-muted-foreground mb-1 uppercase tracking-wide">
                Lịch sử
              </p>
              <p className="text-xs text-muted-foreground line-clamp-2 leading-relaxed">
                {group.detail.history}
              </p>
            </div>
          )}

          {/* Alternative Names (fallback) */}
          {!group.detail?.other_names && group.alternative_names && (
            <div className="mb-4">
              <p className="text-xs font-semibold text-muted-foreground mb-1.5 uppercase tracking-wide">
                Tên khác
              </p>
              <p className="text-sm text-muted-foreground line-clamp-2 leading-relaxed">
                {group.alternative_names}
              </p>
            </div>
          )}

          {/* Footer with arrow */}
          <div className="relative z-10 flex items-center justify-between pt-3 mt-auto border-t border-border/30 group-hover:border-primary/30 transition-colors duration-300">
            <span className="text-xs font-semibold text-primary">
              Xem chi tiết
            </span>
            <ArrowRight className="h-4 w-4 text-primary transition-transform duration-300 group-hover:translate-x-1" />
          </div>
        </div>
      </div>
    </Link>
  );
}
