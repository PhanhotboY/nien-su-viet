'use client';

import { useState, useCallback } from 'react';
import { Input } from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { useImages } from '@/hooks/useImages';
import { ImageUploadZone } from '@/components/cmsdesk/images/ImageUploadZone';
import { ImageGalleryGrid } from '@/components/cmsdesk/images/ImageGalleryGrid';
import { ImageEditModal } from '@/components/cmsdesk/images/ImageEditModal';
import { IMAGE_CATEGORIES, CATEGORY_LABELS } from '@/constants/image.category';
import type { IImage, ImageMetadata } from '@/types/image';
import { toast } from 'sonner';

export default function ImageManagementPage() {
  const {
    images,
    isLoading,
    error,
    uploadImages,
    deleteImage,
    isUploading,
    updateImage,
  } = useImages();

  const [searchQuery, setSearchQuery] = useState('');
  const [categoryFilter, setCategoryFilter] = useState('');
  const [editingImage, setEditingImage] = useState<IImage | null>(null);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const handleFilesSelected = useCallback(
    async (files: File[]) => {
      try {
        await uploadImages(files);
        toast.success(`Successfully uploaded ${files.length} image(s)`);
      } catch (err) {
        console.error('Upload failed:', err);
        const message = err instanceof Error ? err.message : 'Upload failed';
        toast.error(message);
      }
    },
    [uploadImages],
  );

  const handleEditImage = useCallback((image: IImage) => {
    setEditingImage(image);
    setIsEditModalOpen(true);
  }, []);

  const handleDeleteImage = useCallback(
    async (publicId: string) => {
      try {
        await deleteImage(publicId);
        toast.success('Image deleted successfully');
      } catch (err) {
        console.error('Delete failed:', err);
        toast.error('Failed to delete image');
      }
    },
    [deleteImage],
  );

  const handleSaveMetadata = useCallback(
    async (publicId: string, metadata: Partial<ImageMetadata>) => {
      try {
        await updateImage(publicId, metadata);
        toast.success('Image metadata saved successfully');
      } catch (err) {
        console.error('Update failed:', err);
        toast.error('Failed to update image');
      }
    },
    [updateImage],
  );

  return (
    <div className="container space-y-6 py-6">
      {/* Header */}
      <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">
            Image Management
          </h1>
          <p className="text-muted-foreground">
            Manage and organize your Cloudinary images
          </p>
        </div>

        <ImageUploadZone
          onFilesSelected={handleFilesSelected}
          isUploading={isUploading}
          error={error}
        />
      </div>

      {/* Filters Section */}
      <div className="flex flex-col gap-4 sm:flex-row sm:items-end">
        <div className="flex-1">
          <label className="mb-2 block text-sm font-medium text-foreground">
            Search
          </label>
          <Input
            placeholder="Search by title, description, or filename..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
        </div>

        <div className="w-full sm:w-48">
          <label className="mb-2 block text-sm font-medium text-foreground">
            Category
          </label>
          <Select
            value={categoryFilter}
            onValueChange={(v) => setCategoryFilter(v.trim())}
          >
            <SelectTrigger>
              <SelectValue placeholder="All categories" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value=" ">All categories</SelectItem>
              {IMAGE_CATEGORIES.map((category) => (
                <SelectItem key={category} value={category}>
                  {CATEGORY_LABELS[category]}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>
      </div>

      {/* Image Count */}
      <div className="text-sm text-muted-foreground">
        {!isLoading && (
          <>
            {images.length === 0
              ? 'No images'
              : `${images.length} image${images.length !== 1 ? 's' : ''}`}
          </>
        )}
      </div>

      {/* Gallery Grid */}
      <ImageGalleryGrid
        images={images}
        isLoading={isLoading}
        onEdit={handleEditImage}
        onDelete={handleDeleteImage}
        searchQuery={searchQuery}
        categoryFilter={categoryFilter}
      />

      {/* Edit Modal */}
      {isEditModalOpen && (
        <ImageEditModal
          image={editingImage}
          isOpen={isEditModalOpen}
          onClose={() => {
            setIsEditModalOpen(false);
            setTimeout(() => setEditingImage(null), 300);
          }}
          onSave={handleSaveMetadata}
        />
      )}
    </div>
  );
}
