'use client';

import { useState, useEffect } from 'react';
import { toast } from 'sonner';
import { Loader2, Check } from 'lucide-react';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { ConfirmationDialog } from '@/components/ui/confirmation-dialog';
import { Checkbox } from '@/components/ui/checkbox';
import { ScrollArea } from '@/components/ui/scroll-area';
import { getUsers } from '@/services/user.service';
import { getOrganizationMembers } from '@/services/member.service';
import type { components } from '@nsv-interfaces/auth-service';

interface AddMembersDialogProps {
  isOpen: boolean;
  onClose: () => void;
  onSuccess?: () => void;
  organizationId: string;
}

export function AddMembersDialog({
  isOpen,
  onClose,
  onSuccess,
  organizationId,
}: AddMembersDialogProps) {
  const [isLoading, setIsLoading] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');
  const [users, setUsers] = useState<components['schemas']['User'][]>([]);
  const [selectedUserIds, setSelectedUserIds] = useState<Set<string>>(
    new Set(),
  );
  const [isFetching, setIsFetching] = useState(false);
  const [existingMemberIds, setExistingMemberIds] = useState<Set<string>>(
    new Set(),
  );

  // Fetch existing members to prevent duplicates
  useEffect(() => {
    if (isOpen) {
      const fetchExistingMembers = async () => {
        try {
          const membersData = await getOrganizationMembers(organizationId);
          const memberIds = new Set<string>(
            (membersData?.data || membersData || []).map((m: any) => m.userId),
          );
          setExistingMemberIds(memberIds);
        } catch (error) {
          console.error('Failed to fetch existing members:', error);
        }
      };

      fetchExistingMembers();
    }
  }, [isOpen, organizationId]);

  // Fetch users based on search term
  useEffect(() => {
    const fetchUsers = async () => {
      if (!searchTerm.trim()) {
        setUsers([]);
        return;
      }

      setIsFetching(true);
      try {
        const response = await getUsers({
          email: searchTerm,
          limit: '20',
          offset: '0',
        });
        setUsers(response.data || []);
      } catch (error) {
        console.error('Failed to fetch users:', error);
        toast.error('Failed to fetch users');
      } finally {
        setIsFetching(false);
      }
    };

    const timer = setTimeout(() => {
      fetchUsers();
    }, 300);

    return () => clearTimeout(timer);
  }, [searchTerm]);

  const handleToggleUser = (userId: string) => {
    const newSelected = new Set(selectedUserIds);
    if (newSelected.has(userId)) {
      newSelected.delete(userId);
    } else {
      newSelected.add(userId);
    }
    setSelectedUserIds(newSelected);
  };

  const handleAddMembers = async () => {
    if (selectedUserIds.size === 0) {
      toast.error('Please select at least one user');
      return;
    }

    try {
      setIsLoading(true);
      const selectedUsers = Array.from(selectedUserIds);
      let successCount = 0;
      let errorCount = 0;

      for (const userId of selectedUsers) {
        try {
          const response = await fetch(
            `/api/admin/organizations/${organizationId}/members/${userId}`,
            {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
            },
          );

          if (!response.ok) {
            const errorData = await response.json().catch(() => ({}));
            throw new Error(errorData.message || 'Failed to add user');
          }
          successCount++;
        } catch (error) {
          console.error(`Failed to add user ${userId}:`, error);
          errorCount++;
        }
      }

      if (errorCount === 0) {
        toast.success(
          `Successfully added ${successCount} member${successCount > 1 ? 's' : ''}`,
        );
      } else {
        toast.warning(
          `Added ${successCount} member${successCount > 1 ? 's' : ''}, but ${errorCount} failed`,
        );
      }

      onSuccess?.();
      handleClose();
    } catch (error) {
      if (error instanceof Error) {
        toast.error(error.message);
      } else {
        toast.error('Failed to add members');
      }
    } finally {
      setIsLoading(false);
    }
  };

  const handleClose = () => {
    setSearchTerm('');
    setUsers([]);
    setSelectedUserIds(new Set());
    onClose();
  };

  const availableUsers = users.filter(
    (user) => !existingMemberIds.has(user.id!),
  );

  return (
    <ConfirmationDialog
      isOpen={isOpen}
      onClose={handleClose}
      onConfirm={handleAddMembers}
      title="Add Members to Organization"
      description="Search for users and select those you want to add to this organization."
      confirmText={
        isLoading
          ? 'Adding...'
          : `Add ${selectedUserIds.size} Member${selectedUserIds.size !== 1 ? 's' : ''}`
      }
    >
      <div className="grid gap-4 py-4">
        <div className="grid gap-2">
          <Label htmlFor="search">Search Users by Email</Label>
          <Input
            id="search"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            placeholder="Type email address..."
            disabled={isLoading}
          />
        </div>

        <div className="grid gap-2">
          <Label>
            Available Users{' '}
            {availableUsers.length > 0 && `(${availableUsers.length})`}
          </Label>
          {isFetching ? (
            <div className="flex items-center justify-center py-4">
              <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
            </div>
          ) : searchTerm.trim() === '' ? (
            <div className="text-sm text-muted-foreground py-4 text-center">
              Start typing to search for users...
            </div>
          ) : availableUsers.length === 0 ? (
            <div className="text-sm text-muted-foreground py-4 text-center">
              {users.length === 0
                ? 'No users found'
                : 'All found users are already members'}
            </div>
          ) : (
            <ScrollArea className="h-[300px] rounded-md border p-2">
              <div className="space-y-2">
                {availableUsers.map((user) => (
                  <div
                    key={user.id}
                    className="flex items-center space-x-2 p-2 rounded hover:bg-accent cursor-pointer"
                    onClick={() => handleToggleUser(user.id!)}
                  >
                    <Checkbox
                      id={user.id}
                      checked={selectedUserIds.has(user.id!)}
                      onCheckedChange={() => handleToggleUser(user.id!)}
                    />
                    <div className="flex-1 min-w-0">
                      <Label
                        htmlFor={user.id}
                        className="text-sm cursor-pointer font-medium truncate"
                      >
                        {user.name || 'Unknown'}
                      </Label>
                      <p className="text-xs text-muted-foreground truncate">
                        {user.email}
                      </p>
                    </div>
                    {selectedUserIds.has(user.id!) && (
                      <Check className="h-4 w-4 text-green-600" />
                    )}
                  </div>
                ))}
              </div>
            </ScrollArea>
          )}
        </div>
      </div>
    </ConfirmationDialog>
  );
}
