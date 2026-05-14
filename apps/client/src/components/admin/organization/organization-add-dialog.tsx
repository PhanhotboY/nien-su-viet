'use client';

import { useState } from 'react';
import { toast } from 'sonner';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { ConfirmationDialog } from '@/components/ui/confirmation-dialog';
import { createOrganization } from '@/services/organization.service';

interface OrganizationAddDialogProps {
  isOpen: boolean;
  onClose: () => void;
  onSuccess?: () => void;
}

export function OrganizationAddDialog({
  isOpen,
  onClose,
  onSuccess,
}: OrganizationAddDialogProps) {
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    slug: '',
  });

  const handleCreateOrganization = async () => {
    // Validate inputs
    if (!formData.name.trim()) {
      toast.error('Organization name is required');
      return;
    }
    if (!formData.slug.trim()) {
      toast.error('Organization slug is required');
      return;
    }

    try {
      setIsLoading(true);
      await createOrganization({
        name: formData.name,
        slug: formData.slug,
      });
      toast.success(`Organization "${formData.name}" created successfully`);
      onSuccess?.();
      onClose();
      // Reset form
      setFormData({
        name: '',
        slug: '',
      });
    } catch (error) {
      if (error instanceof Error) {
        toast.error(error.message);
      } else {
        toast.error('Failed to create organization. Please try again.');
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <ConfirmationDialog
      isOpen={isOpen}
      onClose={onClose}
      onConfirm={handleCreateOrganization}
      title="Add New Organization"
      description="Create a new organization with the following details."
      confirmText={isLoading ? 'Creating...' : 'Create Organization'}
    >
      <div className="grid gap-4 py-4">
        <div className="grid gap-2">
          <Label htmlFor="org-name">Organization Name</Label>
          <Input
            id="org-name"
            value={formData.name}
            onChange={(e) =>
              setFormData((prev) => ({ ...prev, name: e.target.value }))
            }
            placeholder="Enter organization name"
            required
          />
        </div>
        <div className="grid gap-2">
          <Label htmlFor="org-slug">Organization Slug</Label>
          <Input
            id="org-slug"
            value={formData.slug}
            onChange={(e) =>
              setFormData((prev) => ({
                ...prev,
                slug: e.target.value.toLowerCase().replace(/\s+/g, '-'),
              }))
            }
            placeholder="Enter organization slug"
            required
          />
        </div>
      </div>
    </ConfirmationDialog>
  );
}
