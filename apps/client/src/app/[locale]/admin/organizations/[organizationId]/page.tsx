import type { Metadata } from 'next';
import { Suspense } from 'react';
import { ArrowLeft } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Skeleton } from '@/components/ui/skeleton';
import { OrganizationMembersTable } from '@/components/admin/organization/organization-members-table';
import { getOrganizationById } from '@/services/organization.service';
import { getOrganizationMembers } from '@/services/member.service';
import Link from '@/i18n/navigation';

export const metadata: Metadata = {
  title: 'Admin Dashboard - Organization Details',
  description: 'Manage organization members',
};

function MembersTableSkeleton() {
  return (
    <div className="space-y-4">
      <div className="flex gap-2 items-center mb-4">
        <Skeleton className="h-10 w-10" />
        <Skeleton className="h-8 w-[300px]" />
      </div>
      <div className="flex justify-end mb-4">
        <Skeleton className="h-10 w-[140px]" />
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

export default async function OrganizationDetailPage({
  params,
}: {
  params: Promise<{ organizationId: string; locale: string }>;
}) {
  const resolvedParams = await params;

  // Fetch organization details and members
  const organization = await getOrganizationById(resolvedParams.organizationId);
  const baseLink = `/admin/organizations`;

  if (!organization) {
    return (
      <div className="flex flex-col gap-4 p-4 md:p-6">
        <div className="flex items-center gap-2 mb-4">
          <Link href={baseLink}>
            <Button variant="outline" size="sm">
              <ArrowLeft className="h-4 w-4" />
            </Button>
          </Link>
          <h1 className="text-3xl font-bold">Organization Not Found</h1>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-4 p-4 md:p-6">
      <div className="flex items-center gap-2 mb-4">
        <Link href={baseLink}>
          <Button variant="outline" size="sm">
            <ArrowLeft className="h-4 w-4" />
          </Button>
        </Link>
        <div>
          <h1 className="text-3xl font-bold">{organization.name}</h1>
          <p className="text-sm text-muted-foreground">
            Manage organization members and roles
          </p>
        </div>
      </div>

      <Suspense fallback={<MembersTableSkeleton />}>
        <OrganizationMembersTable
          organizationId={resolvedParams.organizationId}
          membersPromise={getOrganizationMembers(resolvedParams.organizationId)}
        />
      </Suspense>
    </div>
  );
}
