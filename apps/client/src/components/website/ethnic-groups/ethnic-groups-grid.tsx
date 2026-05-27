'use client';

import { useState, useMemo } from 'react';
import { Search } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { EthnicGroupCard } from './ethnic-group-card';

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
  };
}

interface EthnicGroupsGridProps {
  groups: EthnicGroup[];
  searchPlaceholder?: string;
  foundResults?: string;
  results?: string;
  noResultsFound?: string;
  tryAnotherKeyword?: string;
}

export function EthnicGroupsGrid({
  groups,
  searchPlaceholder = 'Search ethnic groups, other names, history...',
  foundResults = 'Found',
  results = 'results',
  noResultsFound = 'No results found',
  tryAnotherKeyword = 'Try searching with different keywords or view all ethnic groups',
}: EthnicGroupsGridProps) {
  const [searchTerm, setSearchTerm] = useState('');

  const filteredGroups = useMemo(() => {
    if (!searchTerm.trim()) {
      return groups;
    }

    const term = searchTerm.toLowerCase();
    return groups.filter((group) => {
      const nameMatch = group.name.toLowerCase().includes(term);
      const slugMatch = group.slug.toLowerCase().includes(term);
      const codeMatch = group.code.includes(term);
      const otherNamesMatch = group.detail?.other_names
        ?.toLowerCase()
        .includes(term);
      const historyMatch = group.detail?.history?.toLowerCase().includes(term);
      const altNamesMatch = group.alternative_names
        ?.toLowerCase()
        .includes(term);

      return (
        nameMatch ||
        slugMatch ||
        codeMatch ||
        otherNamesMatch ||
        historyMatch ||
        altNamesMatch
      );
    });
  }, [searchTerm, groups]);

  return (
    <>
      {/* Search Section */}
      <div className="border-b border-border/40 bg-card/30 backdrop-blur-sm">
        <div className="container mx-auto px-4 py-8 md:py-12">
          <div className="max-w-2xl">
            <div className="relative">
              <Search className="absolute left-3 top-3 h-5 w-5 text-muted-foreground pointer-events-none" />
              <Input
                type="search"
                placeholder={searchPlaceholder}
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="pl-10 py-6 text-base border-border/50 bg-background hover:border-border transition-colors"
              />
            </div>
            {searchTerm && (
              <p className="mt-3 text-sm text-muted-foreground">
                {foundResults}{' '}
                <span className="font-semibold">{filteredGroups.length}</span>{' '}
                {results}
              </p>
            )}
          </div>
        </div>
      </div>

      {/* Main Grid */}
      <div className="container mx-auto px-4 py-16 md:py-24">
        {filteredGroups.length > 0 ? (
          <div className="grid gap-4 sm:gap-5 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
            {filteredGroups.map((group) => (
              <EthnicGroupCard key={group.code} group={group} />
            ))}
          </div>
        ) : (
          <div className="flex flex-col items-center justify-center py-12">
            <div className="rounded-lg bg-muted/50 p-8 text-center max-w-md">
              <Search className="h-12 w-12 text-muted-foreground/50 mx-auto mb-4" />
              <h3 className="text-lg font-semibold text-foreground mb-2">
                {noResultsFound}
              </h3>
              <p className="text-sm text-muted-foreground">
                {tryAnotherKeyword}
              </p>
            </div>
          </div>
        )}
      </div>
    </>
  );
}
