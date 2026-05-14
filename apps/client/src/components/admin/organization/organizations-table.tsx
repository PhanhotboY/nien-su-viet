'use client';
import { Building, Trash2, MoreHorizontal, Plus } from 'lucide-react';
import { format } from 'date-fns';
import { useState, use, useMemo } from 'react';
import Link, { useRouter } from '@/i18n/navigation';

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { Organization } from '@/services/organization.service';
import { OrganizationDeleteDialog } from '@/components/admin/organization/organization-delete-dialog';
import { OrganizationAddDialog } from '@/components/admin/organization/organization-add-dialog';

export function OrganizationsTable({
  organizationsPromise,
}: {
  organizationsPromise: Promise<Organization[]>;
}) {
  const router = useRouter();
  const organizations = use(organizationsPromise);
  const [deleteDialogState, setDeleteDialogState] = useState<{
    isOpen: boolean;
    organizationId?: string;
    organizationName?: string;
  }>({
    isOpen: false,
  });
  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false);

  const handleDeleteClick = (org: Organization) => {
    setDeleteDialogState({
      isOpen: true,
      organizationId: org.id,
      organizationName: org.name,
    });
  };

  const handleDeleteSuccess = () => {
    setDeleteDialogState({ isOpen: false });
    router.refresh();
  };

  const addBtn = (
    <div className="flex flex-wrap gap-2 items-end mb-2 w-full justify-end">
      <Button onClick={() => setIsAddDialogOpen(true)}>
        <Plus className="h-4 w-4" />
        Add Organization
      </Button>
      <OrganizationAddDialog
        isOpen={isAddDialogOpen}
        onClose={() => setIsAddDialogOpen(false)}
        onSuccess={() => {
          setIsAddDialogOpen(false);
          // Refresh the organizations list
          router.refresh();
        }}
      />
    </div>
  );

  if (!organizations?.length)
    return (
      <div className="space-y-4 border-accent-foreground">
        {addBtn}
        <div className="overflow-hidden">
          <Table className="text-sm">
            <TableBody>
              <TableRow>
                <TableCell className="px-4 py-3">
                  <div className="flex items-center gap-4">No data</div>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>
      </div>
    );

  return (
    <div className="space-y-4">
      {addBtn}
      <div className="overflow-hidden rounded-lg border-muted border-2">
        <Table className="text-sm">
          <TableHeader className="bg-muted sticky top-0 z-10">
            <TableRow>
              {[
                { label: 'Name' },
                { label: 'Slug' },
                { label: 'Members' },
                { label: 'Created At' },
                { label: 'Actions', className: 'w-20' },
              ].map((col) => (
                <TableHead
                  key={col.label}
                  className={[
                    col.className,
                    'px-4 py-3 text-xs font-medium text-muted-foreground',
                  ]
                    .filter(Boolean)
                    .join(' ')}
                >
                  {col.label}
                </TableHead>
              ))}
            </TableRow>
          </TableHeader>
          <TableBody>
            {organizations.map((org) => (
              <TableRow key={org.id}>
                <TableCell className="px-4 py-3">
                  <Link
                    href={`organizations/${org.id}`}
                    className="hover:underline"
                  >
                    <div className="flex items-center gap-3">
                      <div className="flex items-center justify-center h-10 w-10 rounded-lg bg-blue-50 dark:bg-blue-900">
                        <Building className="h-5 w-5 text-blue-600 dark:text-blue-300" />
                      </div>
                      <div className="flex flex-col">
                        <span className="text-sm font-medium text-foreground">
                          {org.name}
                        </span>
                        <span className="text-xs text-muted-foreground">
                          ID: {org.id}
                        </span>
                      </div>
                    </div>
                  </Link>
                </TableCell>
                <TableCell className="px-4 py-3">
                  <Badge variant="outline" className="text-xs">
                    {org.slug}
                  </Badge>
                </TableCell>
                <TableCell className="px-4 py-3">
                  <Link href={`organizations/${org.id}`}>
                    <div className="flex items-center gap-2">
                      <Badge
                        variant="secondary"
                        className="text-xs font-semibold cursor-pointer hover:opacity-80"
                      >
                        {org.members?.length || 0} members
                      </Badge>
                    </div>
                  </Link>
                </TableCell>
                <TableCell className="px-4 py-3 text-xs text-muted-foreground">
                  {format(new Date(org.createdAt), "MMM d, yyyy 'at' h:mm a")}
                </TableCell>
                <TableCell className="px-4 py-3">
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="sm">
                        <MoreHorizontal className="h-4 w-4" />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuItem
                        onClick={() => handleDeleteClick(org)}
                        className="text-red-600 dark:text-red-400"
                      >
                        <Trash2 className="h-4 w-4 mr-2" />
                        Delete
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <OrganizationDeleteDialog
        isOpen={deleteDialogState.isOpen}
        organizationId={deleteDialogState.organizationId}
        organizationName={deleteDialogState.organizationName}
        onClose={() => setDeleteDialogState({ isOpen: false })}
        onSuccess={handleDeleteSuccess}
      />
    </div>
  );
}
