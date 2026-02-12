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
import { Separator } from '@/components/ui/separator';
import { Alert, AlertDescription } from '@/components/ui/alert';
import {
  Mail,
  MapPin,
  Phone,
  Clock,
  Send,
  CheckCircle2,
  Facebook,
  Twitter,
  Instagram,
  Linkedin,
} from 'lucide-react';
import Link from 'next/link';

export default function ContactPage() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    phone: '',
    subject: '',
    message: '',
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

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    // Simulate API call
    await new Promise((resolve) => setTimeout(resolve, 2000));

    setIsSubmitting(false);
    setIsSubmitted(true);

    // Reset form
    setFormData({
      name: '',
      email: '',
      phone: '',
      subject: '',
      message: '',
    });

    // Reset success message after 5 seconds
    setTimeout(() => {
      setIsSubmitted(false);
    }, 5000);
  };

  return (
    <div className="container mx-auto py-12">
      <div className="space-y-8">
        {/* Header */}
        <div className="text-center space-y-4">
          <h1 className="text-4xl font-bold tracking-tight">Liên Hệ</h1>
          <p className="text-muted-foreground text-lg mx-auto">
            Chúng tôi luôn sẵn sàng lắng nghe và hỗ trợ bạn. Hãy liên hệ với
            chúng tôi qua biểu mẫu bên dưới hoặc thông qua các kênh liên lạc
            khác.
          </p>
        </div>

        <Separator />

        {/* Success Alert */}
        {isSubmitted && (
          <Alert className="bg-green-50 border-green-200">
            <CheckCircle2 className="h-4 w-4 text-green-600" />
            <AlertDescription className="text-green-800">
              Cảm ơn bạn đã liên hệ! Chúng tôi sẽ phản hồi trong thời gian sớm
              nhất.
            </AlertDescription>
          </Alert>
        )}

        <div className="grid md:grid-cols-3 gap-8">
          {/* Contact Information */}
          <div className="md:col-span-1 space-y-6">
            {/* Contact Cards */}
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">Thông Tin Liên Hệ</CardTitle>
              </CardHeader>
              <CardContent className="space-y-6">
                {/* Address */}
                <div className="flex gap-4">
                  <div className="flex-shrink-0">
                    <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                      <MapPin className="h-5 w-5 text-primary" />
                    </div>
                  </div>
                  <div className="space-y-1">
                    <h3 className="font-semibold">Địa Chỉ</h3>
                    <p className="text-sm text-muted-foreground">
                      123 Đường Lê Lợi
                      <br />
                      Quận 1, TP. Hồ Chí Minh
                      <br />
                      Việt Nam
                    </p>
                  </div>
                </div>

                {/* Phone */}
                <div className="flex gap-4">
                  <div className="flex-shrink-0">
                    <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                      <Phone className="h-5 w-5 text-primary" />
                    </div>
                  </div>
                  <div className="space-y-1">
                    <h3 className="font-semibold">Điện Thoại</h3>
                    <p className="text-sm text-muted-foreground">
                      +84 (28) 1234 5678
                      <br />
                      +84 912 345 678
                    </p>
                  </div>
                </div>

                {/* Email */}
                <div className="flex gap-4">
                  <div className="flex-shrink-0">
                    <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                      <Mail className="h-5 w-5 text-primary" />
                    </div>
                  </div>
                  <div className="space-y-1">
                    <h3 className="font-semibold">Email</h3>
                    <p className="text-sm text-muted-foreground">
                      info@niensu.viet
                      <br />
                      support@niensu.viet
                    </p>
                  </div>
                </div>

                {/* Working Hours */}
                <div className="flex gap-4">
                  <div className="flex-shrink-0">
                    <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center">
                      <Clock className="h-5 w-5 text-primary" />
                    </div>
                  </div>
                  <div className="space-y-1">
                    <h3 className="font-semibold">Giờ Làm Việc</h3>
                    <p className="text-sm text-muted-foreground">
                      Thứ 2 - Thứ 6: 8:00 - 17:00
                      <br />
                      Thứ 7: 8:00 - 12:00
                      <br />
                      Chủ Nhật: Nghỉ
                    </p>
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* Social Media */}
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">Mạng Xã Hội</CardTitle>
                <CardDescription>
                  Kết nối với chúng tôi trên các nền tảng
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="flex gap-3">
                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href="https://facebook.com"
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label="Facebook"
                    >
                      <Facebook className="h-5 w-5" />
                    </a>
                  </Button>
                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href="https://twitter.com"
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label="Twitter"
                    >
                      <Twitter className="h-5 w-5" />
                    </a>
                  </Button>
                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href="https://instagram.com"
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label="Instagram"
                    >
                      <Instagram className="h-5 w-5" />
                    </a>
                  </Button>
                  <Button
                    variant="outline"
                    size="icon"
                    className="rounded-full"
                    asChild
                  >
                    <a
                      href="https://linkedin.com"
                      target="_blank"
                      rel="noopener noreferrer"
                      aria-label="LinkedIn"
                    >
                      <Linkedin className="h-5 w-5" />
                    </a>
                  </Button>
                </div>
              </CardContent>
            </Card>

            {/* FAQ Link */}
            <Card className="bg-primary/5 border-primary/20">
              <CardContent className="pt-6">
                <h3 className="font-semibold mb-2">Câu Hỏi Thường Gặp</h3>
                <p className="text-sm text-muted-foreground mb-4">
                  Bạn có thể tìm thấy câu trả lời cho các câu hỏi phổ biến trong
                  trang FAQ của chúng tôi.
                </p>
                <Button variant="outline" className="w-full" asChild>
                  <Link href="/gioi-thieu#faq">Xem FAQ</Link>
                </Button>
              </CardContent>
            </Card>
          </div>

          {/* Contact Form */}
          <div className="md:col-span-2">
            <Card>
              <CardHeader>
                <CardTitle>Gửi Tin Nhắn</CardTitle>
                <CardDescription>
                  Điền thông tin vào biểu mẫu bên dưới và chúng tôi sẽ liên hệ
                  lại với bạn sớm nhất có thể.
                </CardDescription>
              </CardHeader>
              <CardContent>
                <form onSubmit={handleSubmit} className="space-y-6">
                  {/* Name and Email Row */}
                  <div className="grid md:grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="name">
                        Họ và Tên <span className="text-red-500">*</span>
                      </Label>
                      <Input
                        id="name"
                        name="name"
                        placeholder="Nguyễn Văn A"
                        value={formData.name}
                        onChange={handleChange}
                        required
                      />
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="email">
                        Email <span className="text-red-500">*</span>
                      </Label>
                      <Input
                        id="email"
                        name="email"
                        type="email"
                        placeholder="email@example.com"
                        value={formData.email}
                        onChange={handleChange}
                        required
                      />
                    </div>
                  </div>

                  {/* Phone and Subject Row */}
                  <div className="grid md:grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="phone">Số Điện Thoại</Label>
                      <Input
                        id="phone"
                        name="phone"
                        type="tel"
                        placeholder="0912 345 678"
                        value={formData.phone}
                        onChange={handleChange}
                      />
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="subject">
                        Tiêu Đề <span className="text-red-500">*</span>
                      </Label>
                      <Input
                        id="subject"
                        name="subject"
                        placeholder="Chủ đề tin nhắn"
                        value={formData.subject}
                        onChange={handleChange}
                        required
                      />
                    </div>
                  </div>

                  {/* Message */}
                  <div className="space-y-2">
                    <Label htmlFor="message">
                      Nội Dung <span className="text-red-500">*</span>
                    </Label>
                    <Textarea
                      id="message"
                      name="message"
                      placeholder="Nhập nội dung tin nhắn của bạn..."
                      rows={6}
                      value={formData.message}
                      onChange={handleChange}
                      required
                    />
                  </div>

                  {/* Privacy Notice */}
                  <div className="bg-muted p-4 rounded-lg">
                    <p className="text-sm text-muted-foreground">
                      Bằng việc gửi biểu mẫu này, bạn đồng ý với{' '}
                      <Link
                        href="/privacy"
                        className="text-primary hover:underline"
                      >
                        Chính Sách Bảo Mật
                      </Link>{' '}
                      và{' '}
                      <Link
                        href="/terms"
                        className="text-primary hover:underline"
                      >
                        Điều Khoản Sử Dụng
                      </Link>{' '}
                      của chúng tôi.
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
                        Đang Gửi...
                      </>
                    ) : (
                      <>
                        <Send className="mr-2 h-4 w-4" />
                        Gửi Tin Nhắn
                      </>
                    )}
                  </Button>
                </form>
              </CardContent>
            </Card>

            {/* Map (Optional) */}
            <Card className="mt-6">
              <CardHeader>
                <CardTitle>Bản Đồ</CardTitle>
                <CardDescription>
                  Tìm đường đến văn phòng của chúng tôi
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="aspect-video bg-muted rounded-lg flex items-center justify-center">
                  <p className="text-muted-foreground">
                    [Tích hợp Google Maps hoặc bản đồ khác tại đây]
                  </p>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>

        {/* Additional Help Section */}
        <Card className="bg-gradient-to-r from-primary/10 to-primary/5 border-primary/20">
          <CardContent className="">
            <div className="text-center space-y-4">
              <h2 className="text-2xl font-bold">Cần Hỗ Trợ Ngay?</h2>
              <p className="text-muted-foreground mx-auto">
                Nếu bạn cần hỗ trợ khẩn cấp hoặc có câu hỏi cần giải đáp ngay
                lập tức, đừng ngần ngại liên hệ với chúng tôi qua hotline hoặc
                chat trực tuyến.
              </p>
              <div className="flex gap-4 justify-center flex-wrap">
                <Button size="lg" asChild>
                  <a href="tel:+842812345678">
                    <Phone className="mr-2 h-4 w-4" />
                    Gọi Hotline
                  </a>
                </Button>
                <Button size="lg" variant="outline" asChild>
                  <a href="mailto:support@niensu.viet">
                    <Mail className="mr-2 h-4 w-4" />
                    Gửi Email
                  </a>
                </Button>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
