'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Loader2, Calendar } from 'lucide-react';
import { components } from '@nsv-interfaces/historical-event';
import TextEditor from '@/components/TextEditor';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';

interface EventFormProps {
  initialData?: components['schemas']['HistoricalEventBaseCreateDto'];
  categories: components['schemas']['EventCategoryBaseDto'][];
  onSubmit: (
    data: components['schemas']['HistoricalEventBaseCreateDto'],
  ) => Promise<any>;
  submitLabel?: string;
}

export function EventForm({
  initialData,
  categories,
  onSubmit,
  submitLabel = 'Create Event',
}: EventFormProps) {
  const [formData, setFormData] = useState<
    components['schemas']['HistoricalEventBaseCreateDto']
  >({
    name: initialData?.name || '',
    content: initialData?.content || '',
    fromYear: initialData?.fromYear || new Date().getFullYear(),
    fromMonth: initialData?.fromMonth || null,
    fromDay: initialData?.fromDay || null,
    toYear: initialData?.toYear || null,
    toMonth: initialData?.toMonth || null,
    toDay: initialData?.toDay || null,
    thumbnailId: initialData?.thumbnailId,
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const router = useRouter();

  // const [selectedCategories, setSelectedCategories] = useState<Set<string>>(
  //   new Set(initialData?.categoryIds || []),
  // );

  const handleSubmit = async (e: React.FormEvent) => {
    setIsSubmitting(true);
    e.preventDefault();

    try {
      await onSubmit(formData);
      toast.success('Event updated successfully!');
      router.push('/cmsdesk/historical-events');
    } catch (error) {
      toast.error('Failed to update event');
      console.error(error);
    }

    setIsSubmitting(false);
  };

  // const handleCategoryToggle = (categoryId: string) => {
  //   const newSelected = new Set(selectedCategories);
  //   if (newSelected.has(categoryId)) {
  //     newSelected.delete(categoryId);
  //   } else {
  //     newSelected.add(categoryId);
  //   }
  //   setSelectedCategories(newSelected);
  // };

  const updateField = <
    K extends keyof components['schemas']['HistoricalEventBaseCreateDto'],
  >(
    field: K,
    value: components['schemas']['HistoricalEventBaseCreateDto'][K],
  ) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  const months = [
    'January',
    'February',
    'March',
    'April',
    'May',
    'June',
    'July',
    'August',
    'September',
    'October',
    'November',
    'December',
  ];

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      {/* Basic Information */}
      <Card>
        <CardHeader>
          <CardTitle>Basic Information</CardTitle>
          <CardDescription>
            Enter the main details of the historical event
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name">Event Name *</Label>
            <Input
              id="name"
              placeholder="e.g., Battle of Bach Dang River"
              value={formData.name}
              onChange={(e) => updateField('name', e.target.value)}
              required
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="content">Content *</Label>

            <TextEditor
              value={formData.content}
              onChange={(value) => updateField('content', value)}
            />
          </div>
        </CardContent>
      </Card>

      {/* Date Information */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Calendar className="h-5 w-5" />
            Date Information
          </CardTitle>
          <CardDescription>Specify when the event occurred</CardDescription>
        </CardHeader>
        <CardContent className="space-y-6">
          {/* Start Date */}
          <div className="space-y-3">
            <Label className="text-base font-semibold">Start Date *</Label>
            <div className="grid gap-4 sm:grid-cols-3">
              <div className="space-y-2">
                <Label htmlFor="fromYear" className="text-sm font-normal">
                  Year *
                </Label>
                <Input
                  id="fromYear"
                  type="number"
                  placeholder="YYYY"
                  value={formData.fromYear}
                  onChange={(e) => {
                    const value = e.target.value;
                    updateField('fromYear', value as any);
                  }}
                  required
                />
              </div>
              {!!formData.fromYear && (
                <div className="space-y-2">
                  <Label htmlFor="fromMonth" className="text-sm font-normal">
                    Month
                  </Label>
                  <Select
                    value={formData.fromMonth?.toString() || ''}
                    onValueChange={(value) =>
                      updateField('fromMonth', value ? parseInt(value) : null)
                    }
                  >
                    <SelectTrigger className="w-full" id="fromMonth">
                      <SelectValue placeholder="Select month" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="0">Không chọn</SelectItem>
                      {months.map((month, index) => (
                        <SelectItem key={index} value={(index + 1).toString()}>
                          {month}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
              )}
              {!!(formData.fromYear && formData.fromMonth) && (
                <div className="space-y-2">
                  <Label htmlFor="fromDay" className="text-sm font-normal">
                    Day
                  </Label>
                  <Input
                    id="fromDay"
                    type="number"
                    min="1"
                    max="31"
                    placeholder="DD"
                    value={formData.fromDay || ''}
                    onChange={(e) =>
                      updateField(
                        'fromDay',
                        e.target.value ? parseInt(e.target.value) : null,
                      )
                    }
                  />
                </div>
              )}
            </div>
          </div>

          {/* End Date */}
          <div className="space-y-3">
            <Label className="text-base font-semibold">
              End Date (Optional)
            </Label>
            <p className="text-sm text-muted-foreground">
              Leave empty for single-day events
            </p>
            <div className="grid gap-4 sm:grid-cols-3">
              {!!formData.fromYear && (
                <div className="space-y-2">
                  <Label htmlFor="toYear" className="text-sm font-normal">
                    Year
                  </Label>
                  <Input
                    id="toYear"
                    type="number"
                    placeholder="YYYY"
                    value={formData.toYear || ''}
                    onChange={(e) => {
                      updateField('toYear', e.target.value as any);
                    }}
                  />
                </div>
              )}
              {!!formData.toYear && (
                <div className="space-y-2">
                  <Label htmlFor="toMonth" className="text-sm font-normal">
                    Month
                  </Label>
                  <Select
                    value={formData.toMonth?.toString() || ''}
                    onValueChange={(value) =>
                      updateField('toMonth', value ? parseInt(value) : null)
                    }
                  >
                    <SelectTrigger id="toMonth" className="w-full">
                      <SelectValue placeholder="Select month" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="0">Không chọn</SelectItem>
                      {months.map((month, index) => (
                        <SelectItem key={index} value={(index + 1).toString()}>
                          {month}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
              )}
              {!!(formData.toYear && formData.toMonth) && (
                <div className="space-y-2">
                  <Label htmlFor="toDay" className="text-sm font-normal">
                    Day
                  </Label>
                  <Input
                    id="toDay"
                    type="number"
                    min="1"
                    max="31"
                    placeholder="DD"
                    value={formData.toDay || ''}
                    onChange={(e) =>
                      updateField(
                        'toDay',
                        e.target.value ? parseInt(e.target.value) : null,
                      )
                    }
                  />
                </div>
              )}
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Categories */}
      <Card>
        <CardHeader>
          <CardTitle>Categories</CardTitle>
          <CardDescription>
            Select categories that apply to this event
          </CardDescription>
        </CardHeader>
        <CardContent>
          {categories.length > 0 ? (
            <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
              {/* {categories.map((category) => (
                <div
                  key={category.id}
                  className="flex items-start space-x-3 space-y-0"
                >
                  <Checkbox
                    id={`category-${category.id}`}
                    checked={selectedCategories.has(category.id)}
                    onCheckedChange={() => handleCategoryToggle(category.id)}
                  />
                  <div className="space-y-1 leading-none">
                    <Label
                      htmlFor={`category-${category.id}`}
                      className="cursor-pointer font-medium"
                    >
                      {category.name}
                    </Label>
                    {category.description && (
                      <p className="text-xs text-muted-foreground">
                        {category.description}
                      </p>
                    )}
                  </div>
                </div>
              ))} */}
            </div>
          ) : (
            <p className="text-sm text-muted-foreground">
              No categories available. Create categories first.
            </p>
          )}
        </CardContent>
      </Card>

      {/* Submit Button */}
      <div className="flex justify-end gap-4">
        <Button
          type="button"
          variant="outline"
          onClick={() => window.history.back()}
          disabled={isSubmitting}
        >
          Cancel
        </Button>
        <Button type="submit" disabled={isSubmitting}>
          {isSubmitting ? (
            <>
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              Saving...
            </>
          ) : (
            submitLabel
          )}
        </Button>
      </div>
    </form>
  );
}
