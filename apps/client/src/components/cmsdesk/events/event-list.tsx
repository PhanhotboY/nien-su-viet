'use client';

import { useState, useEffect, useMemo } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import useSWR from 'swr';

import { EventCard } from '@/components/cmsdesk/events/event-card';
import { EventListItem } from '@/components/cmsdesk/events/event-list-item';
import { DeleteEventDialog } from '@/components/cmsdesk/events/delete-event-dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Skeleton } from '@/components/ui/skeleton';
import {
  Plus,
  Grid3x3,
  List,
  Search,
  Calendar,
  SortAsc,
  SortDesc,
} from 'lucide-react';
import Link from 'next/link';
import { toast } from 'sonner';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { IPaginatedResponse } from '@/interfaces/response.interface';
import { swrFetcher } from '@/helper/swrFetcher';

type ViewMode = 'grid' | 'list';

export default function EventList({
  handleDelete,
}: {
  handleDelete: (eventId: string) => Promise<void>;
}) {
  const router = useRouter();
  const searchParams = useSearchParams();

  // State initialized from URL
  const [viewMode, setViewMode] = useState<ViewMode>(
    (searchParams.get('view') as ViewMode) || 'grid',
  );
  const [searchTerm, setSearchTerm] = useState(
    searchParams.get('search') || '',
  );
  const [debouncedSearch, setDebouncedSearch] = useState(searchTerm);
  const [selectedCategory, setSelectedCategory] = useState(
    searchParams.get('category') || 'all',
  );
  const [sortBy, setSortBy] = useState(
    searchParams.get('sortBy') || 'createdAt',
  );
  const [sortDirection, setSortDirection] = useState<'asc' | 'desc'>(
    (searchParams.get('sortDirection') as 'asc' | 'desc') || 'desc',
  );
  const [currentPage, setCurrentPage] = useState(
    Number(searchParams.get('page')) || 1,
  );
  const [deleteDialog, setDeleteDialog] = useState<{
    open: boolean;
    eventId: string;
    eventName: string;
  }>({
    open: false,
    eventId: '',
    eventName: '',
  });

  const itemsPerPage = viewMode === 'grid' ? 15 : 20;

  // Debounce search
  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedSearch(searchTerm);
    }, 300);

    return () => clearTimeout(timer);
  }, [searchTerm]);

  // Update URL when filters/sort/page change
  useEffect(() => {
    const params = new URLSearchParams();
    if (viewMode !== 'grid') params.set('view', viewMode);
    if (debouncedSearch) params.set('search', debouncedSearch);
    if (selectedCategory !== 'all') params.set('category', selectedCategory);
    if (sortBy !== 'createdAt') params.set('sortBy', sortBy);
    if (sortDirection !== 'desc') params.set('sortOrder', sortDirection);
    if (currentPage !== 1) params.set('page', String(currentPage));

    router.replace(`?${params.toString()}`);
  }, [
    viewMode,
    debouncedSearch,
    selectedCategory,
    sortBy,
    sortDirection,
    currentPage,
    router,
  ]);

  // Build SWR key with all params
  const swrKey = useMemo(() => {
    const params = new URLSearchParams();
    params.set('limit', String(itemsPerPage));
    params.set('page', String(currentPage));
    if (debouncedSearch) params.set('search', debouncedSearch);
    if (selectedCategory !== 'all') params.set('categoryId', selectedCategory);
    params.set('sortBy', sortBy);
    params.set('sortOrder', sortDirection);

    return `/api/historical-events?${params.toString()}`;
  }, [
    itemsPerPage,
    currentPage,
    debouncedSearch,
    selectedCategory,
    sortBy,
    sortDirection,
  ]);

  const { data, error, mutate, isLoading } = useSWR<
    IPaginatedResponse<components['schemas']['HistoricalEventBriefResponseDto']>
  >(swrKey, swrFetcher, {
    revalidateOnFocus: false,
    dedupingInterval: 2000,
  });

  const openDeleteDialog = (eventId: string, eventName: string) => {
    setDeleteDialog({ open: true, eventId, eventName });
  };

  const events = data?.data || [];
  const total = data?.pagination?.total || 0;
  const totalPages = Math.ceil(total / itemsPerPage);

  const toggleSortDirection = () => {
    setSortDirection((prev) => (prev === 'asc' ? 'desc' : 'asc'));
  };

  return (
    <div className="container mx-auto space-y-6">
      {/* Filters and Controls */}
      <div className="flex flex-col gap-4 rounded-lg border bg-card p-4">
        <div className="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
          {/* Search */}
          <div className="relative flex-1 lg:max-w-sm">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              placeholder="Search events..."
              value={searchTerm}
              onChange={(e) => {
                setSearchTerm(e.target.value);
                setCurrentPage(1);
              }}
              className="pl-9"
            />
          </div>

          {/* View Mode Toggle */}
          <div className="flex items-center gap-2">
            <Button
              variant={viewMode === 'grid' ? 'default' : 'outline'}
              size="sm"
              onClick={() => setViewMode('grid')}
              className="h-9 w-9 p-0"
            >
              <Grid3x3 className="h-4 w-4" />
            </Button>
            <Button
              variant={viewMode === 'list' ? 'default' : 'outline'}
              size="sm"
              onClick={() => setViewMode('list')}
              className="h-9 w-9 p-0"
            >
              <List className="h-4 w-4" />
            </Button>
          </div>
        </div>

        {/* Filters Row */}
        <div className="flex flex-col gap-4 sm:flex-row">
          <Select
            value={selectedCategory}
            onValueChange={(value) => {
              setSelectedCategory(value);
              setCurrentPage(1);
            }}
          >
            <SelectTrigger className="w-full sm:w-[200px]">
              <SelectValue placeholder="All Categories" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">All Categories</SelectItem>
              {/* TODO: Load categories from API */}
            </SelectContent>
          </Select>

          <Select value={sortBy} onValueChange={setSortBy}>
            <SelectTrigger className="w-full sm:w-[200px]">
              <SelectValue placeholder="Sort by" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="createdAt">Date Created</SelectItem>
              <SelectItem value="updatedAt">Last Updated</SelectItem>
              <SelectItem value="fromYear">Event Year</SelectItem>
              <SelectItem value="name">Name</SelectItem>
            </SelectContent>
          </Select>

          <Button
            variant="outline"
            size="sm"
            onClick={toggleSortDirection}
            className="w-full sm:w-auto"
          >
            {sortDirection === 'asc' ? (
              <>
                <SortAsc className="mr-2 h-4 w-4" />
                Ascending
              </>
            ) : (
              <>
                <SortDesc className="mr-2 h-4 w-4" />
                Descending
              </>
            )}
          </Button>
        </div>

        {/* Results Count */}
        <div className="text-sm text-muted-foreground">
          Showing {events.length} of {total} event{total !== 1 ? 's' : ''}
        </div>
      </div>

      {/* Events Grid/List */}
      {isLoading ? (
        <div
          className={
            viewMode === 'grid'
              ? 'grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5'
              : 'space-y-4'
          }
        >
          {Array.from({ length: itemsPerPage }).map((_, i) => (
            <Skeleton
              key={i}
              className={viewMode === 'grid' ? 'h-[400px]' : 'h-[140px]'}
            />
          ))}
        </div>
      ) : events.length === 0 ? (
        <div className="flex min-h-[400px] flex-col items-center justify-center rounded-lg border border-dashed bg-muted/20 p-8 text-center">
          <Calendar className="mb-4 h-16 w-16 text-muted-foreground/40" />
          <h3 className="mb-2 text-lg font-semibold">No events found</h3>
          <p className="mb-4 text-sm text-muted-foreground">
            {searchTerm || selectedCategory !== 'all'
              ? 'Try adjusting your filters'
              : 'Get started by creating your first historical event'}
          </p>
          <Button asChild>
            <Link href="/cmsdesk/historical-events/new">
              <Plus className="mr-2 h-4 w-4" />
              Create Event
            </Link>
          </Button>
        </div>
      ) : viewMode === 'grid' ? (
        <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5">
          {events.map((event) => (
            <EventCard
              key={event.id}
              event={event}
              onDelete={(id) => openDeleteDialog(id, event.name)}
            />
          ))}
        </div>
      ) : (
        <div className="space-y-4">
          {events.map((event) => (
            <EventListItem
              key={event.id}
              event={event}
              onDelete={(id) => openDeleteDialog(id, event.name)}
            />
          ))}
        </div>
      )}

      {/* Pagination */}
      {totalPages > 1 && (
        <div className="flex items-center justify-center gap-2">
          <Button
            variant="outline"
            onClick={() => setCurrentPage((p) => Math.max(1, p - 1))}
            disabled={currentPage === 1}
          >
            Previous
          </Button>
          <div className="flex items-center gap-1">
            {Array.from({ length: Math.min(5, totalPages) }, (_, i) => {
              let pageNum;
              if (totalPages <= 5) {
                pageNum = i + 1;
              } else if (currentPage <= 3) {
                pageNum = i + 1;
              } else if (currentPage >= totalPages - 2) {
                pageNum = totalPages - 4 + i;
              } else {
                pageNum = currentPage - 2 + i;
              }

              return (
                <Button
                  key={pageNum}
                  variant={currentPage === pageNum ? 'default' : 'outline'}
                  size="sm"
                  onClick={() => setCurrentPage(pageNum)}
                  className="h-9 w-9 p-0"
                >
                  {pageNum}
                </Button>
              );
            })}
          </div>
          <Button
            variant="outline"
            onClick={() => setCurrentPage((p) => Math.min(totalPages, p + 1))}
            disabled={currentPage === totalPages}
          >
            Next
          </Button>
        </div>
      )}

      {/* Delete Dialog */}
      <DeleteEventDialog
        open={deleteDialog.open}
        onOpenChange={(open) => setDeleteDialog((prev) => ({ ...prev, open }))}
        onConfirm={async () => {
          try {
            await handleDelete(deleteDialog.eventId);
            toast.success('Event deleted successfully');
            await mutate();
          } catch (error) {
            toast.error('Failed to delete event');
            console.error(error);
          }
        }}
        eventName={deleteDialog.eventName}
      />
    </div>
  );
}
