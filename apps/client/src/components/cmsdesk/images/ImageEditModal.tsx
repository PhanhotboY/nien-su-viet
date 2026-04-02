'use client';

import { useState, useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from '@/components/ui/dialog';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Button } from '@/components/ui/button';
import { IMAGE_CATEGORIES, CATEGORY_LABELS } from '@/constants/image.category';
import type { IImage } from '@/types/image';
import type { ImageMetadata } from '@/types/image';

const imageMetadataSchema = z.object({
  category: z.enum(IMAGE_CATEGORIES),
  alt: z.string().max(200, 'Alt text must be less than 200 characters'),
  caption: z.string().max(500, 'Caption must be less than 500 characters'),
});

type ImageMetadataFormData = z.infer<typeof imageMetadataSchema>;

interface ImageEditModalProps {
  image: IImage | null;
  isOpen: boolean;
  onClose: () => void;
  onSave: (publicId: string, metadata: Partial<ImageMetadata>) => void;
}

export function ImageEditModal({
  image,
  isOpen,
  onClose,
  onSave,
}: ImageEditModalProps) {
  const [isSaving, setIsSaving] = useState(false);

  const form = useForm<ImageMetadataFormData>({
    resolver: zodResolver(imageMetadataSchema),
    defaultValues: {
      category: image?.metadata?.category || IMAGE_CATEGORIES[8],
      alt: image?.metadata?.alt || '',
      caption: image?.metadata?.caption || '',
    },
  });

  const onSubmit = async (data: ImageMetadataFormData) => {
    if (!image) return;

    setIsSaving(true);
    try {
      onSave(image.public_id, data);
      onClose();
    } catch (error) {
      console.error('Failed to save metadata:', error);
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>Edit Image Metadata</DialogTitle>
          <DialogDescription>
            Update the title, description, category, alt text, and caption for
            this image.
          </DialogDescription>
        </DialogHeader>

        {image && (
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
              {/* Caption */}
              <FormField
                control={form.control}
                name="caption"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Caption</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder="Caption to display with the image"
                        {...field}
                        rows={2}
                      />
                    </FormControl>
                    <FormDescription>
                      Displayed below the image in galleries
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              {/* Alt Text */}
              <FormField
                control={form.control}
                name="alt"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Alt Text</FormLabel>
                    <FormControl>
                      <Input
                        placeholder="Descriptive alt text for accessibility"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      Used for SEO and accessibility
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              {/* Category */}
              <FormField
                control={form.control}
                name="category"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Category</FormLabel>
                    <Select value={field.value} onValueChange={field.onChange}>
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {IMAGE_CATEGORIES.map((category) => (
                          <SelectItem key={category} value={category}>
                            {CATEGORY_LABELS[category]}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <DialogFooter>
                <Button
                  type="button"
                  variant="outline"
                  onClick={onClose}
                  disabled={isSaving}
                >
                  Cancel
                </Button>
                <Button type="submit" disabled={isSaving}>
                  {isSaving ? 'Saving...' : 'Save Changes'}
                </Button>
              </DialogFooter>
            </form>
          </Form>
        )}
      </DialogContent>
    </Dialog>
  );
}
