'use client';

import { use, useState } from 'react';
import { MoreHorizontal, Plus, User } from 'lucide-react';
import { format } from 'date-fns';
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
import { AddMembersDialog } from './add-members-dialog';
import { UpdateMemberRoleDialog } from './update-member-role-dialog';
import { RemoveMemberDialog } from './remove-member-dialog';
import { IPaginatedResponse } from '@/interfaces/response.interface';
import { OrganizationMember } from '@/services/member.service';

interface Member {
  id: string;
  userId: string;
  organizationId: string;
  role: string;
  user?: {
    id: string;
    name: string;
    email: string;
  };
  createdAt: string;
}

export function OrganizationMembersTable({
  organizationId,
  membersPromise,
}: {
  organizationId: string;
  membersPromise: Promise<IPaginatedResponse<OrganizationMember>>;
}) {
  const data = use(membersPromise);
  const members: Member[] = data?.data || data || [];

  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false);
  const [updateRoleDialogState, setUpdateRoleDialogState] = useState<{
    isOpen: boolean;
    member?: Member;
  }>({ isOpen: false });
  const [removeMemberDialogState, setRemoveMemberDialogState] = useState<{
    isOpen: boolean;
    member?: Member;
  }>({ isOpen: false });

  const handleUpdateRoleClick = (member: Member) => {
    setUpdateRoleDialogState({
      isOpen: true,
      member,
    });
  };

  const handleRemoveClick = (member: Member) => {
    setRemoveMemberDialogState({
      isOpen: true,
      member,
    });
  };

  const handleRefresh = () => {
    // Trigger a page reload or use SWR if available
    window.location.reload();
  };

  const addBtn = (
    <div className="flex flex-wrap gap-2 items-end mb-2 w-full justify-end">
      <Button onClick={() => setIsAddDialogOpen(true)}>
        <Plus className="h-4 w-4" />
        Add Member
      </Button>
    </div>
  );

  if (!members?.length) {
    return (
      <div className="space-y-4 border-accent-foreground">
        {addBtn}
        <AddMembersDialog
          isOpen={isAddDialogOpen}
          onClose={() => setIsAddDialogOpen(false)}
          organizationId={organizationId}
          onSuccess={handleRefresh}
        />
        <div className="overflow-hidden">
          <Table className="text-sm">
            <TableBody>
              <TableRow>
                <TableCell className="px-4 py-3">
                  <div className="flex items-center gap-4">No members yet</div>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {addBtn}
      <div className="overflow-hidden rounded-lg border-muted border-2">
        <Table className="text-sm">
          <TableHeader className="bg-muted sticky top-0 z-10">
            <TableRow>
              {[
                { label: 'Name' },
                { label: 'Email' },
                { label: 'Role' },
                { label: 'Joined At' },
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
            {members.map((member) => (
              <TableRow key={member.id}>
                <TableCell className="px-4 py-3">
                  <div className="flex items-center gap-3">
                    <div className="flex items-center justify-center h-10 w-10 rounded-lg bg-purple-50 dark:bg-purple-900">
                      <User className="h-5 w-5 text-purple-600 dark:text-purple-300" />
                    </div>
                    <div className="flex flex-col">
                      <span className="text-sm font-medium text-foreground">
                        {member.user?.name || 'Unknown'}
                      </span>
                      <span className="text-xs text-muted-foreground">
                        ID: {member.userId}
                      </span>
                    </div>
                  </div>
                </TableCell>
                <TableCell className="px-4 py-3">
                  <span className="text-sm">{member.user?.email || '-'}</span>
                </TableCell>
                <TableCell className="px-4 py-3">
                  <Badge variant="outline" className="text-xs capitalize">
                    {member.role}
                  </Badge>
                </TableCell>
                <TableCell className="px-4 py-3 text-xs text-muted-foreground">
                  {format(
                    new Date(member.createdAt),
                    "MMM d, yyyy 'at' h:mm a",
                  )}
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
                        onClick={() => handleUpdateRoleClick(member)}
                      >
                        Update Role
                      </DropdownMenuItem>
                      <DropdownMenuItem
                        onClick={() => handleRemoveClick(member)}
                        className="text-red-600 dark:text-red-400"
                      >
                        Remove Member
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <AddMembersDialog
        isOpen={isAddDialogOpen}
        onClose={() => setIsAddDialogOpen(false)}
        organizationId={organizationId}
        onSuccess={handleRefresh}
      />

      <UpdateMemberRoleDialog
        isOpen={updateRoleDialogState.isOpen}
        member={updateRoleDialogState.member}
        organizationId={organizationId}
        onClose={() => setUpdateRoleDialogState({ isOpen: false })}
        onSuccess={handleRefresh}
      />

      <RemoveMemberDialog
        isOpen={removeMemberDialogState.isOpen}
        member={removeMemberDialogState.member}
        organizationId={organizationId}
        onClose={() => setRemoveMemberDialogState({ isOpen: false })}
        onSuccess={handleRefresh}
      />
    </div>
  );
}
