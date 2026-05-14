'use client';

import { useState, useEffect } from 'react';
import { toast } from 'sonner';
import { Label } from '@/components/ui/label';
import { ConfirmationDialog } from '@/components/ui/confirmation-dialog';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { roles } from '@/lib/permissions';

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

interface UpdateMemberRoleDialogProps {
  isOpen: boolean;
  member?: Member;
  organizationId: string;
  onClose: () => void;
  onSuccess?: () => void;
}

export function UpdateMemberRoleDialog({
  isOpen,
  member,
  organizationId,
  onClose,
  onSuccess,
}: UpdateMemberRoleDialogProps) {
  const [isLoading, setIsLoading] = useState(false);
  const [selectedRole, setSelectedRole] = useState<string>(
    member?.role || 'member',
  );

  useEffect(() => {
    if (member) {
      setSelectedRole(member.role);
    }
  }, [member]);

  const handleUpdateRole = async () => {
    if (!member || selectedRole === member.role) {
      toast.error('Please select a different role');
      return;
    }

    try {
      setIsLoading(true);

      // Call update member role API
      // Note: This endpoint should be created in your backend at:
      // PUT /auth/admin/organizations/{organizationId}/members/{memberId}/role
      const response = await fetch(
        `/api/admin/organizations/${organizationId}/members/${member.id}/role`,
        {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ role: selectedRole }),
        },
      );

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || 'Failed to update member role');
      }

      toast.success('Member role updated successfully');
      onSuccess?.();
      onClose();
    } catch (error) {
      if (error instanceof Error) {
        toast.error(error.message);
      } else {
        toast.error('Failed to update member role');
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
      onConfirm={handleUpdateRole}
      title="Update Member Role"
      description={`Update the role for ${member.user?.name || 'this member'}`}
      confirmText={isLoading ? 'Updating...' : 'Update Role'}
    >
      <div className="grid gap-4 py-4">
        <div className="space-y-2">
          <p className="text-sm text-muted-foreground">
            <strong>Member:</strong> {member.user?.name || 'Unknown'}
          </p>
          <p className="text-sm text-muted-foreground">
            <strong>Email:</strong> {member.user?.email || 'N/A'}
          </p>
        </div>

        <div className="grid gap-2">
          <Label htmlFor="role">Select New Role</Label>
          <Select
            value={selectedRole}
            onValueChange={setSelectedRole}
            disabled={isLoading}
          >
            <SelectTrigger id="role" className="w-full">
              <SelectValue placeholder="Select role" />
            </SelectTrigger>
            <SelectContent>
              {Object.keys(roles).map((role) => (
                <SelectItem key={role} value={role}>
                  {role.charAt(0).toUpperCase() + role.slice(1)}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <p className="text-xs text-muted-foreground">
            Current role: <strong>{member.role}</strong>
          </p>
        </div>
      </div>
    </ConfirmationDialog>
  );
}
