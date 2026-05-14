import type { Metadata } from 'next';
import { Suspense } from 'react';
import { OrganizationsTable } from '@/components/admin/organization/organizations-table';
import { Skeleton } from '@/components/ui/skeleton';
import { getOrganizations } from '@/services/organization.service';

export const metadata: Metadata = {
  title: 'Admin Dashboard - Organizations',
  description: 'Manage organizations in the admin dashboard',
};

function OrganizationsTableSkeleton() {
  return (
    <div className="space-y-4">
      <div className="flex flex-wrap gap-2 items-end mb-2 w-full justify-between">
        <div className="flex gap-2 items-end">
          <Skeleton className="h-10 w-[200px]" />
          <Skeleton className="h-10 w-[140px]" />
        </div>
        <Skeleton className="h-10 w-[120px]" />
      </div>
      <div className="overflow-hidden rounded-lg border-muted border-2">
        <div className="space-y-2 p-4">
          {Array.from({ length: 5 }).map((_, i) => (
            <Skeleton key={i} className="h-16 w-full" />
          ))}
        </div>
      </div>
    </div>
  );
}

export default async function OrganizationsPage() {
  return (
    <div className="flex flex-col gap-4 p-4 md:p-6">
      <Suspense fallback={<OrganizationsTableSkeleton />}>
        <OrganizationsTable organizationsPromise={getOrganizations()} />
      </Suspense>
    </div>
  );
}
