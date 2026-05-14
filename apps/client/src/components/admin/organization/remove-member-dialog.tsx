'use client';

import { useState } from 'react';
import { toast } from 'sonner';
import { ConfirmationDialog } from '@/components/ui/confirmation-dialog';

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

interface RemoveMemberDialogProps {
  isOpen: boolean;
  member?: Member;
  organizationId: string;
  onClose: () => void;
  onSuccess?: () => void;
}

export function RemoveMemberDialog({
  isOpen,
  member,
  organizationId,
  onClose,
  onSuccess,
}: RemoveMemberDialogProps) {
  const [isLoading, setIsLoading] = useState(false);

  const handleRemoveMember = async () => {
    if (!member) return;

    try {
      setIsLoading(true);

      // Call server action via API route
      const response = await fetch(
        `/api/admin/organizations/${organizationId}/members/${member.id}`,
        {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json',
          },
        },
      );

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || 'Failed to remove member');
      }

      toast.success(
        `${member.user?.name || 'Member'} removed from organization`,
      );
      onSuccess?.();
      onClose();
    } catch (error) {
      if (error instanceof Error) {
        toast.error(error.message);
      } else {
        toast.error('Failed to remove member');
      }
    } finally {
      setIsLoading(false);
    }
  };

  if (!member) return null;

  return (
    <ConfirmationDialog
      isOpen={isOpen}
      onClose={onClose}
      onConfirm={handleRemoveMember}
      title="Remove Member"
      description={`Are you sure you want to remove ${member.user?.name || 'this member'} from the organization? This action cannot be undone.`}
      confirmText={isLoading ? 'Removing...' : 'Remove Member'}
    >
      <div className="grid gap-4 py-4">
        <div className="space-y-2">
          <p className="text-sm text-muted-foreground">
            <strong>Member:</strong> {member.user?.name || 'Unknown'}
          </p>
          <p className="text-sm text-muted-foreground">
            <strong>Email:</strong> {member.user?.email || 'N/A'}
          </p>
          <p className="text-sm text-muted-foreground">
            <strong>Role:</strong> {member.role}
          </p>
        </div>
        <div className="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-900 rounded p-3">
          <p className="text-sm text-red-800 dark:text-red-200">
            This member will lose access to this organization and all its
            resources.
          </p>
        </div>
      </div>
    </ConfirmationDialog>
  );
}
