'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { toast } from 'sonner';
import { Save, Loader2 } from 'lucide-react';

import { getAppInfo, updateApp } from '@/services/cms.service';
import { components } from '@nsv-interfaces/cms-service';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Label } from '@/components/ui/label';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';

type AppData = components['schemas']['AppData'];
type AppUpdateReq = components['schemas']['AppUpdateReq'];

export default function AppSettingsPage() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(true);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formData, setFormData] = useState<AppUpdateReq>({
    title: '',
    description: '',
    logo: '',
    social: {
      facebook: '',
      youtube: '',
      tiktok: '',
      zalo: '',
    },
    tax_code: '',
    address: {
      province: '',
      district: '',
      street: '',
    },
    msisdn: '',
    email: '',
    map: '',
    head_scripts: '',
    body_scripts: '',
  });

  useEffect(() => {
    const fetchAppInfo = async () => {
      try {
        setIsLoading(true);
        const response = await getAppInfo();
        if (response.data) {
          setFormData({
            title: response.data.title || '',
            description: response.data.description || '',
            logo: response.data.logo || '',
            social: {
              facebook: response.data.social?.facebook || '',
              youtube: response.data.social?.youtube || '',
              tiktok: response.data.social?.tiktok || '',
              zalo: response.data.social?.zalo || '',
            },
            tax_code: response.data.tax_code || '',
            address: {
              province: response.data.address?.province || '',
              district: response.data.address?.district || '',
              street: response.data.address?.street || '',
            },
            msisdn: response.data.msisdn || '',
            email: response.data.email || '',
            map: response.data.map || '',
            head_scripts: response.data.head_scripts || '',
            body_scripts: response.data.body_scripts || '',
          });
        }
      } catch (error: any) {
        toast.error('Failed to load app settings', {
          description: error.message || 'Please try again later.',
        });
      } finally {
        setIsLoading(false);
      }
    };

    fetchAppInfo();
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      setIsSubmitting(true);

      // Clean up empty strings to undefined for optional fields
      const payload: AppUpdateReq = {
        title: formData.title,
        description: formData.description || undefined,
        logo: formData.logo || undefined,
        social:
          formData.social?.facebook ||
          formData.social?.youtube ||
          formData.social?.tiktok ||
          formData.social?.zalo
            ? {
                facebook: formData.social.facebook || undefined,
                youtube: formData.social.youtube || undefined,
                tiktok: formData.social.tiktok || undefined,
                zalo: formData.social.zalo || undefined,
              }
            : undefined,
        tax_code: formData.tax_code || undefined,
        address:
          formData.address?.province ||
          formData.address?.district ||
          formData.address?.street
            ? {
                province: formData.address.province || undefined,
                district: formData.address.district || undefined,
                street: formData.address.street || undefined,
              }
            : undefined,
        msisdn: formData.msisdn || undefined,
        email: formData.email || undefined,
        map: formData.map || undefined,
        head_scripts: formData.head_scripts || undefined,
        body_scripts: formData.body_scripts || undefined,
      };

      await updateApp(payload);

      toast.success('App settings updated successfully!', {
        description: 'Your changes have been saved.',
      });

      router.refresh();
    } catch (error: any) {
      toast.error('Failed to update app settings', {
        description: error.message || 'Please try again later.',
      });
    } finally {
      setIsSubmitting(false);
    }
  };

  if (isLoading) {
    return (
      <div className="container mx-auto flex h-[calc(100vh-200px)] items-center justify-center">
        <div className="flex flex-col items-center gap-2">
          <Loader2 className="h-8 w-8 animate-spin text-primary" />
          <p className="text-sm text-muted-foreground">
            Loading app settings...
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto space-y-6 py-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">App Settings</h1>
        <p className="text-muted-foreground">
          Manage your website configuration and information
        </p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Basic Information */}
        <Card>
          <CardHeader>
            <CardTitle>Basic Information</CardTitle>
            <CardDescription>Main details about your website</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="title">
                Website Title <span className="text-destructive">*</span>
              </Label>
              <Input
                id="title"
                value={formData.title}
                onChange={(e) =>
                  setFormData({ ...formData, title: e.target.value })
                }
                placeholder="Nien Su Viet"
                required
                maxLength={255}
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                value={formData.description || ''}
                onChange={(e) =>
                  setFormData({ ...formData, description: e.target.value })
                }
                placeholder="Vietnam history timeline website"
                maxLength={1000}
                rows={3}
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="logo">Logo URL</Label>
              <Input
                id="logo"
                type="url"
                value={formData.logo || ''}
                onChange={(e) =>
                  setFormData({ ...formData, logo: e.target.value })
                }
                placeholder="https://example.com/logo.png"
              />
            </div>
          </CardContent>
        </Card>

        {/* Social Media Links */}
        <Card>
          <CardHeader>
            <CardTitle>Social Media</CardTitle>
            <CardDescription>
              Connect your social media profiles
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid gap-4 md:grid-cols-2">
              <div className="space-y-2">
                <Label htmlFor="facebook">Facebook</Label>
                <Input
                  id="facebook"
                  type="url"
                  value={formData.social?.facebook || ''}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      social: { ...formData.social, facebook: e.target.value },
                    })
                  }
                  placeholder="https://facebook.com/example"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="youtube">YouTube</Label>
                <Input
                  id="youtube"
                  type="url"
                  value={formData.social?.youtube || ''}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      social: { ...formData.social, youtube: e.target.value },
                    })
                  }
                  placeholder="https://youtube.com/@example"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="tiktok">TikTok</Label>
                <Input
                  id="tiktok"
                  type="url"
                  value={formData.social?.tiktok || ''}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      social: { ...formData.social, tiktok: e.target.value },
                    })
                  }
                  placeholder="https://tiktok.com/@example"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="zalo">Zalo</Label>
                <Input
                  id="zalo"
                  type="url"
                  value={formData.social?.zalo || ''}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      social: { ...formData.social, zalo: e.target.value },
                    })
                  }
                  placeholder="https://zalo.me/example"
                />
              </div>
            </div>
          </CardContent>
        </Card>

        {/* Contact Information */}
        <Card>
          <CardHeader>
            <CardTitle>Contact Information</CardTitle>
            <CardDescription>How people can reach you</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid gap-4 md:grid-cols-2">
              <div className="space-y-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  value={formData.email || ''}
                  onChange={(e) =>
                    setFormData({ ...formData, email: e.target.value })
                  }
                  placeholder="contact@example.com"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="msisdn">Phone Number</Label>
                <Input
                  id="msisdn"
                  type="tel"
                  value={formData.msisdn || ''}
                  onChange={(e) =>
                    setFormData({ ...formData, msisdn: e.target.value })
                  }
                  placeholder="+84987654321"
                />
              </div>
            </div>

            <div className="space-y-2">
              <Label htmlFor="tax_code">Tax Code</Label>
              <Input
                id="tax_code"
                value={formData.tax_code || ''}
                onChange={(e) =>
                  setFormData({ ...formData, tax_code: e.target.value })
                }
                placeholder="0123456789"
                maxLength={50}
              />
            </div>
          </CardContent>
        </Card>

        {/* Address */}
        <Card>
          <CardHeader>
            <CardTitle>Address</CardTitle>
            <CardDescription>Your physical location</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="street">Street Address</Label>
              <Input
                id="street"
                value={formData.address?.street || ''}
                onChange={(e) =>
                  setFormData({
                    ...formData,
                    address: { ...formData.address, street: e.target.value },
                  })
                }
                placeholder="123 Nguyen Hue Blvd"
                maxLength={255}
              />
            </div>

            <div className="grid gap-4 md:grid-cols-2">
              <div className="space-y-2">
                <Label htmlFor="district">District</Label>
                <Input
                  id="district"
                  value={formData.address?.district || ''}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      address: {
                        ...formData.address,
                        district: e.target.value,
                      },
                    })
                  }
                  placeholder="District 1"
                  maxLength={100}
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="province">Province</Label>
                <Input
                  id="province"
                  value={formData.address?.province || ''}
                  onChange={(e) =>
                    setFormData({
                      ...formData,
                      address: {
                        ...formData.address,
                        province: e.target.value,
                      },
                    })
                  }
                  placeholder="Ho Chi Minh City"
                  maxLength={100}
                />
              </div>
            </div>

            <div className="space-y-2">
              <Label htmlFor="map">Map URL</Label>
              <Input
                id="map"
                type="url"
                value={formData.map || ''}
                onChange={(e) =>
                  setFormData({ ...formData, map: e.target.value })
                }
                placeholder="https://maps.google.com/?q=location"
              />
            </div>
          </CardContent>
        </Card>

        {/* Scripts */}
        <Card>
          <CardHeader>
            <CardTitle>Custom Scripts</CardTitle>
            <CardDescription>
              Add custom scripts to your website (analytics, tracking, etc.)
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="head_scripts">Head Scripts</Label>
              <Textarea
                id="head_scripts"
                value={formData.head_scripts || ''}
                onChange={(e) =>
                  setFormData({ ...formData, head_scripts: e.target.value })
                }
                placeholder="<script>console.log('head');</script>"
                maxLength={5000}
                rows={5}
                className="font-mono text-xs"
              />
              <p className="text-xs text-muted-foreground">
                Scripts to be included in the {'<head>'} section
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="body_scripts">Body Scripts</Label>
              <Textarea
                id="body_scripts"
                value={formData.body_scripts || ''}
                onChange={(e) =>
                  setFormData({ ...formData, body_scripts: e.target.value })
                }
                placeholder="<script>console.log('body');</script>"
                maxLength={5000}
                rows={5}
                className="font-mono text-xs"
              />
              <p className="text-xs text-muted-foreground">
                Scripts to be included before the closing {'</body>'} tag
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Submit Button */}
        <div className="flex justify-end gap-4">
          <Button
            type="button"
            variant="outline"
            onClick={() => router.back()}
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
              <>
                <Save className="mr-2 h-4 w-4" />
                Save Changes
              </>
            )}
          </Button>
        </div>
      </form>
    </div>
  );
}
