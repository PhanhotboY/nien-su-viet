'use client';

import { useEffect, useState } from 'react';
import { components } from '@nsv-interfaces/cms-service';
import {
  getHeaderNavItems,
  getFooterNavItems,
  updateHeaderNavItem,
  updateFooterNavItem,
  createHeaderNavItem,
  createFooterNavItem,
  deleteHeaderNavItem,
  deleteFooterNavItem,
} from '@/services/cms.service';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Switch } from '@/components/ui/switch';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import { toast } from 'sonner';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
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
import { Skeleton } from '@/components/ui/skeleton';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { Plus, MoreVertical, Pencil, Trash2, ExternalLink } from 'lucide-react';

type HeaderNavItem = components['schemas']['HeaderNavItemData'];
type FooterNavItem = components['schemas']['FooterNavItemData'];
type LinkType = 'internal' | 'external';

interface EditingItem {
  id?: string;
  order: number;
  link_label: string;
  link_url: string;
  link_type: LinkType;
  link_new_tab: boolean;
}

export default function NavMenusPage() {
  const [headerItems, setHeaderItems] = useState<HeaderNavItem[]>([]);
  const [footerItems, setFooterItems] = useState<FooterNavItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [editDialogOpen, setEditDialogOpen] = useState(false);
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [editingItem, setEditingItem] = useState<EditingItem | null>(null);
  const [deletingItem, setDeletingItem] = useState<{
    id: string;
    label: string;
    type: 'header' | 'footer';
  } | null>(null);
  const [editingType, setEditingType] = useState<'header' | 'footer'>('header');
  const [creatingType, setCreatingType] = useState<'header' | 'footer'>(
    'header',
  );
  const [saving, setSaving] = useState(false);
  const [deleting, setDeleting] = useState(false);

  useEffect(() => {
    loadNavItems();
  }, []);

  const loadNavItems = async () => {
    try {
      setLoading(true);
      const [headerResponse, footerResponse] = await Promise.all([
        getHeaderNavItems(),
        getFooterNavItems(),
      ]);

      if (
        !Array.isArray(headerResponse.data) ||
        !Array.isArray(footerResponse.data)
      ) {
        throw new Error('Failed to fetch navigation items');
      }

      setHeaderItems(headerResponse.data);
      setFooterItems(footerResponse.data);
    } catch (error) {
      toast.error('Failed to load navigation items');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const handleEdit = (
    item: HeaderNavItem | FooterNavItem,
    type: 'header' | 'footer',
  ) => {
    setEditingItem({
      id: item.id,
      order: item.order,
      link_label: item.link_label,
      link_url: item.link_url,
      link_type: item.link_type as LinkType,
      link_new_tab: item.link_new_tab ?? false,
    });
    setEditingType(type);
    setEditDialogOpen(true);
  };

  const handleCreateNew = (type: 'header' | 'footer') => {
    const maxOrder =
      type === 'header'
        ? Math.max(0, ...headerItems.map((i) => i.order))
        : Math.max(0, ...footerItems.map((i) => i.order));

    setEditingItem({
      order: maxOrder + 1,
      link_label: '',
      link_url: '',
      link_type: 'internal',
      link_new_tab: false,
    });
    setCreatingType(type);
    setCreateDialogOpen(true);
  };

  const handleSave = async () => {
    if (!editingItem) return;

    try {
      setSaving(true);
      const updateData = {
        order: editingItem.order,
        link_label: editingItem.link_label,
        link_url: editingItem.link_url,
        link_type: editingItem.link_type,
        link_new_tab: editingItem.link_new_tab,
      };

      if (editingType === 'header') {
        await updateHeaderNavItem(editingItem.id!, updateData);
      } else {
        await updateFooterNavItem(editingItem.id!, updateData);
      }

      toast.success('Navigation item updated successfully');
      setEditDialogOpen(false);
      setEditingItem(null);
      await loadNavItems();
    } catch (error) {
      toast.error('Failed to update navigation item');
      console.error(error);
    } finally {
      setSaving(false);
    }
  };

  const handleCreate = async () => {
    if (!editingItem) return;

    if (!editingItem.link_label || !editingItem.link_url) {
      toast.error('Please fill in all required fields');
      return;
    }

    try {
      setSaving(true);
      const createData = {
        order: editingItem.order,
        link_label: editingItem.link_label,
        link_url: editingItem.link_url,
        link_type: editingItem.link_type,
        link_new_tab: editingItem.link_new_tab,
      };

      if (creatingType === 'header') {
        await createHeaderNavItem(createData);
      } else {
        await createFooterNavItem(createData);
      }

      toast.success('Navigation item created successfully');
      setCreateDialogOpen(false);
      setEditingItem(null);
      await loadNavItems();
    } catch (error) {
      toast.error('Failed to create navigation item');
      console.error(error);
    } finally {
      setSaving(false);
    }
  };

  const handleDeleteClick = (
    item: HeaderNavItem | FooterNavItem,
    type: 'header' | 'footer',
  ) => {
    setDeletingItem({
      id: item.id,
      label: item.link_label,
      type,
    });
    setDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!deletingItem) return;

    try {
      setDeleting(true);
      if (deletingItem.type === 'header') {
        await deleteHeaderNavItem(deletingItem.id);
      } else {
        await deleteFooterNavItem(deletingItem.id);
      }

      toast.success('Navigation item deleted successfully');
      setDeleteDialogOpen(false);
      setDeletingItem(null);
      await loadNavItems();
    } catch (error) {
      toast.error('Failed to delete navigation item');
      console.error(error);
    } finally {
      setDeleting(false);
    }
  };

  const renderNavItemCard = (
    item: HeaderNavItem | FooterNavItem,
    type: 'header' | 'footer',
  ) => (
    <Card key={item.id} className="hover:shadow-md transition-shadow">
      <CardContent className="pt-4 pb-4">
        <div className="flex items-center justify-between gap-3">
          <div className="flex items-center gap-2 flex-1 min-w-0">
            <Badge variant="outline" className="font-mono shrink-0 text-xs">
              #{item.order}
            </Badge>
            <div className="flex-1 min-w-0">
              <div className="font-semibold truncate text-sm">
                {item.link_label}
              </div>
              <div className="flex items-center gap-1.5 text-xs text-muted-foreground mt-0.5">
                <span className="truncate">{item.link_url}</span>
                {item.link_type === 'external' && (
                  <ExternalLink className="h-3 w-3 shrink-0" />
                )}
              </div>
            </div>
          </div>
          <div className="flex items-center gap-2 shrink-0">
            {item.link_new_tab && (
              <Badge variant="outline" className="text-xs">
                New tab
              </Badge>
            )}
            <Badge
              variant={item.link_type === 'internal' ? 'default' : 'secondary'}
              className="text-xs"
            >
              {item.link_type}
            </Badge>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" size="sm" className="h-7 w-7 p-0">
                  <MoreVertical className="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent align="end">
                <DropdownMenuItem onClick={() => handleEdit(item, type)}>
                  <Pencil className="h-4 w-4 mr-2" />
                  Edit
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem
                  onClick={() => handleDeleteClick(item, type)}
                  className="text-destructive focus:text-destructive"
                >
                  <Trash2 className="h-4 w-4 mr-2" />
                  Delete
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </div>
      </CardContent>
    </Card>
  );

  const renderEditDialog = () => (
    <Dialog open={editDialogOpen} onOpenChange={setEditDialogOpen}>
      <DialogContent className="sm:max-w-[500px]">
        <DialogHeader>
          <DialogTitle>Edit Navigation Item</DialogTitle>
          <DialogDescription>
            Update the {editingType} navigation menu item details.
          </DialogDescription>
        </DialogHeader>
        {editingItem && (
          <div className="grid gap-4 py-4">
            <div className="grid gap-2">
              <Label htmlFor="order">Display Order</Label>
              <Input
                id="order"
                type="number"
                value={editingItem.order}
                onChange={(e) =>
                  setEditingItem({
                    ...editingItem,
                    order: parseInt(e.target.value) || 0,
                  })
                }
              />
              <p className="text-xs text-muted-foreground">
                Lower numbers appear first
              </p>
            </div>

            <div className="grid gap-2">
              <Label htmlFor="label">
                Label <span className="text-destructive">*</span>
              </Label>
              <Input
                id="label"
                value={editingItem.link_label}
                onChange={(e) =>
                  setEditingItem({
                    ...editingItem,
                    link_label: e.target.value,
                  })
                }
                placeholder="Enter menu label"
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="url">
                URL <span className="text-destructive">*</span>
              </Label>
              <Input
                id="url"
                value={editingItem.link_url}
                onChange={(e) =>
                  setEditingItem({
                    ...editingItem,
                    link_url: e.target.value,
                  })
                }
                placeholder="/about or https://example.com"
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="type">Link Type</Label>
              <Select
                value={editingItem.link_type}
                onValueChange={(value: LinkType) =>
                  setEditingItem({
                    ...editingItem,
                    link_type: value,
                  })
                }
              >
                <SelectTrigger id="type">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="internal">Internal</SelectItem>
                  <SelectItem value="external">External</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="flex items-center justify-between rounded-lg border p-3">
              <div className="space-y-0.5">
                <Label htmlFor="new-tab">Open in new tab</Label>
                <div className="text-sm text-muted-foreground">
                  Open link in a new browser tab
                </div>
              </div>
              <Switch
                id="new-tab"
                checked={editingItem.link_new_tab}
                onCheckedChange={(checked) =>
                  setEditingItem({
                    ...editingItem,
                    link_new_tab: checked,
                  })
                }
              />
            </div>
          </div>
        )}
        <DialogFooter>
          <Button
            variant="outline"
            onClick={() => setEditDialogOpen(false)}
            disabled={saving}
          >
            Cancel
          </Button>
          <Button onClick={handleSave} disabled={saving}>
            {saving ? 'Saving...' : 'Save Changes'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );

  const renderCreateDialog = () => (
    <Dialog open={createDialogOpen} onOpenChange={setCreateDialogOpen}>
      <DialogContent className="sm:max-w-[500px]">
        <DialogHeader>
          <DialogTitle>Create Navigation Item</DialogTitle>
          <DialogDescription>
            Add a new {creatingType} navigation menu item.
          </DialogDescription>
        </DialogHeader>
        {editingItem && (
          <div className="grid gap-4 py-4">
            <div className="grid gap-2">
              <Label htmlFor="create-order">Display Order</Label>
              <Input
                id="create-order"
                type="number"
                value={editingItem.order}
                onChange={(e) =>
                  setEditingItem({
                    ...editingItem,
                    order: parseInt(e.target.value) || 0,
                  })
                }
              />
              <p className="text-xs text-muted-foreground">
                Lower numbers appear first
              </p>
            </div>

            <div className="grid gap-2">
              <Label htmlFor="create-label">
                Label <span className="text-destructive">*</span>
              </Label>
              <Input
                id="create-label"
                value={editingItem.link_label}
                onChange={(e) =>
                  setEditingItem({
                    ...editingItem,
                    link_label: e.target.value,
                  })
                }
                placeholder="Enter menu label"
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="create-url">
                URL <span className="text-destructive">*</span>
              </Label>
              <Input
                id="create-url"
                value={editingItem.link_url}
                onChange={(e) =>
                  setEditingItem({
                    ...editingItem,
                    link_url: e.target.value,
                  })
                }
                placeholder="/about or https://example.com"
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="create-type">Link Type</Label>
              <Select
                value={editingItem.link_type}
                onValueChange={(value: LinkType) =>
                  setEditingItem({
                    ...editingItem,
                    link_type: value,
                  })
                }
              >
                <SelectTrigger id="create-type">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="internal">Internal</SelectItem>
                  <SelectItem value="external">External</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="flex items-center justify-between rounded-lg border p-3">
              <div className="space-y-0.5">
                <Label htmlFor="create-new-tab">Open in new tab</Label>
                <div className="text-sm text-muted-foreground">
                  Open link in a new browser tab
                </div>
              </div>
              <Switch
                id="create-new-tab"
                checked={editingItem.link_new_tab}
                onCheckedChange={(checked) =>
                  setEditingItem({
                    ...editingItem,
                    link_new_tab: checked,
                  })
                }
              />
            </div>
          </div>
        )}
        <DialogFooter>
          <Button
            variant="outline"
            onClick={() => setCreateDialogOpen(false)}
            disabled={saving}
          >
            Cancel
          </Button>
          <Button onClick={handleCreate} disabled={saving}>
            {saving ? 'Creating...' : 'Create Item'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );

  const renderDeleteDialog = () => (
    <AlertDialog open={deleteDialogOpen} onOpenChange={setDeleteDialogOpen}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This will permanently delete the navigation item &quot;
            {deletingItem?.label}&quot;. This action cannot be undone.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel disabled={deleting}>Cancel</AlertDialogCancel>
          <AlertDialogAction
            onClick={handleDelete}
            disabled={deleting}
            className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
          >
            {deleting ? 'Deleting...' : 'Delete'}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );

  const renderLoadingSkeleton = () => (
    <div className="grid gap-3">
      {[1, 2, 3].map((i) => (
        <Card key={i}>
          <CardContent className="pt-4 pb-4">
            <div className="flex items-center justify-between gap-3">
              <div className="flex items-center gap-2 flex-1">
                <Skeleton className="h-5 w-10" />
                <div className="flex-1">
                  <Skeleton className="h-4 w-32 mb-1.5" />
                  <Skeleton className="h-3 w-48" />
                </div>
              </div>
              <div className="flex items-center gap-2">
                <Skeleton className="h-5 w-16" />
                <Skeleton className="h-5 w-16" />
                <Skeleton className="h-7 w-7" />
              </div>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );

  return (
    <div className="container mx-auto py-8 px-4">
      <div className="space-y-6">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">
              Navigation Menus
            </h1>
            <p className="text-muted-foreground mt-2">
              Manage your website&apos;s header and footer navigation items
            </p>
          </div>
        </div>

        <Separator />

        {/* Tabs for Header and Footer */}
        <Tabs defaultValue="header" className="space-y-6 max-w-5xl mx-auto">
          <TabsList className="grid w-full grid-cols-2 max-w-md py-0">
            <TabsTrigger value="header">
              Header Navigation
              {!loading && (
                <Badge variant="secondary" className="ml-2">
                  {headerItems.length}
                </Badge>
              )}
            </TabsTrigger>
            <TabsTrigger value="footer">
              Footer Navigation
              {!loading && (
                <Badge variant="secondary" className="ml-2">
                  {footerItems.length}
                </Badge>
              )}
            </TabsTrigger>
          </TabsList>

          {/* Header Items */}
          <TabsContent value="header" className="space-y-4">
            <Card>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <div>
                    <CardTitle>Header Menu Items</CardTitle>
                    <CardDescription>
                      Navigation items displayed in the website header
                    </CardDescription>
                  </div>
                  <Button
                    onClick={() => handleCreateNew('header')}
                    size="sm"
                    className="shrink-0"
                  >
                    <Plus className="h-4 w-4 mr-2" />
                    Add Item
                  </Button>
                </div>
              </CardHeader>
              <CardContent>
                {loading ? (
                  renderLoadingSkeleton()
                ) : headerItems.length === 0 ? (
                  <div className="text-center py-12">
                    <p className="text-muted-foreground mb-4">
                      No header navigation items found
                    </p>
                    <Button
                      onClick={() => handleCreateNew('header')}
                      variant="outline"
                    >
                      <Plus className="h-4 w-4 mr-2" />
                      Create your first item
                    </Button>
                  </div>
                ) : (
                  <div className="grid gap-3">
                    {headerItems
                      .sort((a, b) => a.order - b.order)
                      .map((item) => renderNavItemCard(item, 'header'))}
                  </div>
                )}
              </CardContent>
            </Card>
          </TabsContent>

          {/* Footer Items */}
          <TabsContent value="footer" className="space-y-4">
            <Card>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <div>
                    <CardTitle>Footer Menu Items</CardTitle>
                    <CardDescription>
                      Navigation items displayed in the website footer
                    </CardDescription>
                  </div>
                  <Button
                    onClick={() => handleCreateNew('footer')}
                    size="sm"
                    className="shrink-0"
                  >
                    <Plus className="h-4 w-4 mr-2" />
                    Add Item
                  </Button>
                </div>
              </CardHeader>
              <CardContent>
                {loading ? (
                  renderLoadingSkeleton()
                ) : footerItems.length === 0 ? (
                  <div className="text-center py-12">
                    <p className="text-muted-foreground mb-4">
                      No footer navigation items found
                    </p>
                    <Button
                      onClick={() => handleCreateNew('footer')}
                      variant="outline"
                    >
                      <Plus className="h-4 w-4 mr-2" />
                      Create your first item
                    </Button>
                  </div>
                ) : (
                  <div className="grid gap-3">
                    {footerItems
                      .sort((a, b) => a.order - b.order)
                      .map((item) => renderNavItemCard(item, 'footer'))}
                  </div>
                )}
              </CardContent>
            </Card>
          </TabsContent>
        </Tabs>

        {/* Dialogs */}
        {renderEditDialog()}
        {renderCreateDialog()}
        {renderDeleteDialog()}
      </div>
    </div>
  );
}
