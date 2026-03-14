'use client';

import { useState } from 'react';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Button } from '@/components/ui/button';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { CheckCircle2, Mail, Phone, Send } from 'lucide-react';
import Link from '@/i18n/navigation';
import { useTranslations } from 'next-intl';

export type ContactFormProps = {
  /**
   * Optional prefilled values (useful if you want to inject from URL params/user profile).
   */
  initialValues?: Partial<{
    name: string;
    email: string;
    phone: string;
    subject: string;
    message: string;
  }>;

  /**
   * Called when the user submits the form.
   * If omitted, the form will simulate a request (same behavior as the previous inline implementation).
   */
  onSubmit?: (data: {
    name: string;
    email: string;
    phone: string;
    subject: string;
    message: string;
  }) => Promise<void>;

  /**
   * Optional contact shortcuts rendered under the form (e.g. from metadata).
   */
  contactShortcuts?: {
    msisdn?: string;
    email?: string;
  };
};

export default function ContactForm(props: ContactFormProps) {
  const t = useTranslations('ContactForm');

  const [formData, setFormData] = useState({
    name: props.initialValues?.name ?? '',
    email: props.initialValues?.email ?? '',
    phone: props.initialValues?.phone ?? '',
    subject: props.initialValues?.subject ?? '',
    message: props.initialValues?.message ?? '',
  });

  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isSubmitted, setIsSubmitted] = useState(false);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const resetForm = () => {
    setFormData({
      name: '',
      email: '',
      phone: '',
      subject: '',
      message: '',
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    setIsSubmitting(true);
    try {
      if (props.onSubmit) {
        await props.onSubmit(formData);
      } else {
        // Default behavior: simulate API call (keeps parity with old page implementation)
        await new Promise((resolve) => setTimeout(resolve, 2000));
      }

      setIsSubmitted(true);
      resetForm();

      // Reset success message after 5 seconds
      setTimeout(() => {
        setIsSubmitted(false);
      }, 5000);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="space-y-6">
      {/* Success Alert */}
      {isSubmitted && (
        <Alert className="bg-green-50 border-green-200">
          <CheckCircle2 className="h-4 w-4 text-green-600" />
          <AlertDescription className="text-green-800">
            {t('alerts.success')}
          </AlertDescription>
        </Alert>
      )}

      <Card>
        <CardHeader>
          <CardTitle>{t('header.title')}</CardTitle>
          <CardDescription>{t('header.description')}</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
            {/* Name and Email Row */}
            <div className="grid md:grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="name">
                  {t('fields.name.label')}{' '}
                  <span className="text-red-500">*</span>
                </Label>
                <Input
                  id="name"
                  name="name"
                  placeholder={t('fields.name.placeholder')}
                  value={formData.name}
                  onChange={handleChange}
                  required
                  autoComplete="name"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="email">
                  {t('fields.email.label')}{' '}
                  <span className="text-red-500">*</span>
                </Label>
                <Input
                  id="email"
                  name="email"
                  type="email"
                  placeholder={t('fields.email.placeholder')}
                  value={formData.email}
                  onChange={handleChange}
                  required
                  autoComplete="email"
                />
              </div>
            </div>

            {/* Phone and Subject Row */}
            <div className="grid md:grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="phone">{t('fields.phone.label')}</Label>
                <Input
                  id="phone"
                  name="phone"
                  type="tel"
                  placeholder={t('fields.phone.placeholder')}
                  value={formData.phone}
                  onChange={handleChange}
                  autoComplete="tel"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="subject">
                  {t('fields.subject.label')}{' '}
                  <span className="text-red-500">*</span>
                </Label>
                <Input
                  id="subject"
                  name="subject"
                  placeholder={t('fields.subject.placeholder')}
                  value={formData.subject}
                  onChange={handleChange}
                  required
                />
              </div>
            </div>

            {/* Message */}
            <div className="space-y-2">
              <Label htmlFor="message">
                {t('fields.message.label')}{' '}
                <span className="text-red-500">*</span>
              </Label>
              <Textarea
                id="message"
                name="message"
                placeholder={t('fields.message.placeholder')}
                rows={6}
                value={formData.message}
                onChange={handleChange}
                required
              />
            </div>

            {/* Privacy Notice */}
            <div className="bg-muted p-4 rounded-lg">
              <p className="text-sm text-muted-foreground">
                {t('privacy.prefix')}{' '}
                <Link
                  href={t('privacy.privacyHref')}
                  className="text-primary hover:underline"
                >
                  {t('privacy.privacyLinkText')}
                </Link>{' '}
                {t('privacy.and')}{' '}
                <Link
                  href={t('privacy.termsHref')}
                  className="text-primary hover:underline"
                >
                  {t('privacy.termsLinkText')}
                </Link>{' '}
                {t('privacy.suffix')}
              </p>
            </div>

            {/* Submit Button */}
            <Button
              type="submit"
              className="w-full"
              size="lg"
              disabled={isSubmitting}
            >
              {isSubmitting ? (
                <>
                  <span className="animate-spin mr-2">⏳</span>
                  {t('actions.submitting')}
                </>
              ) : (
                <>
                  <Send className="mr-2 h-4 w-4" />
                  {t('actions.submit')}
                </>
              )}
            </Button>
          </form>
        </CardContent>
      </Card>

      {/* Optional shortcuts */}
      {(props.contactShortcuts?.msisdn || props.contactShortcuts?.email) && (
        <Card className="bg-gradient-to-r from-primary/10 to-primary/5 border-primary/20">
          <CardContent>
            <div className="text-center space-y-4">
              <h2 className="text-2xl font-bold">{t('shortcuts.title')}</h2>
              <p className="text-muted-foreground mx-auto">
                {t('shortcuts.description')}
              </p>

              <div className="flex gap-4 justify-center flex-wrap">
                {props.contactShortcuts?.msisdn && (
                  <Button size="lg" asChild>
                    <a href={`tel:${props.contactShortcuts.msisdn}`}>
                      <Phone className="mr-2 h-4 w-4" />
                      {t('shortcuts.callHotline')}
                    </a>
                  </Button>
                )}

                {props.contactShortcuts?.email && (
                  <Button size="lg" variant="outline" asChild>
                    <a href={`mailto:${props.contactShortcuts.email}`}>
                      <Mail className="mr-2 h-4 w-4" />
                      {t('shortcuts.sendEmail')}
                    </a>
                  </Button>
                )}
              </div>
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
