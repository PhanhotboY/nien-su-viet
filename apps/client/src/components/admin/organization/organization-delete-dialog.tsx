'use client';

import { useState } from 'react';
import { Trash2 } from 'lucide-react';
import { toast } from 'sonner';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { deleteOrganization } from '@/services/organization.service';

interface OrganizationDeleteDialogProps {
  isOpen: boolean;
  organizationId?: string;
  organizationName?: string;
  onClose: () => void;
  onSuccess: () => void;
}

export function OrganizationDeleteDialog({
  isOpen,
  organizationId,
  organizationName,
  onClose,
  onSuccess,
}: OrganizationDeleteDialogProps) {
  const [isLoading, setIsLoading] = useState(false);

  const handleDelete = async () => {
    if (!organizationId) return;

    setIsLoading(true);
    try {
      await deleteOrganization(organizationId);
      toast.success(`Organization "${organizationName}" has been deleted.`);
      onSuccess();
    } catch (error: any) {
      toast.error(
        error?.message || 'Failed to delete organization. Please try again.',
      );
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <AlertDialog open={isOpen} onOpenChange={onClose}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle className="flex items-center gap-2">
            <Trash2 className="h-5 w-5 text-red-600" />
            Delete Organization
          </AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure you want to delete <strong>{organizationName}</strong>?
            This action cannot be undone and will remove all associated data.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel disabled={isLoading}>Cancel</AlertDialogCancel>
          <AlertDialogAction
            onClick={handleDelete}
            disabled={isLoading}
            className="bg-red-600 text-white hover:bg-red-700"
          >
            {isLoading ? 'Deleting...' : 'Delete'}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
