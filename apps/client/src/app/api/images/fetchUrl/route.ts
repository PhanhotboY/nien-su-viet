import { authClient } from '@/lib/auth-client';
import { uploadImages } from '@/services/image.service';
import { NextRequest, NextResponse } from 'next/server';

export async function POST(request: NextRequest) {
  const session = await authClient.getSession();
  const body = await request.json();

  if (!session) {
    return NextResponse.json(
      {
        success: false,
        toast: { type: 'error', message: 'Vui lòng đăng nhập!' },
      },
      { status: 401 },
    );
  }

  try {
    const res = await fetch(body.url);

    const blob = await res.blob();
    const file = new File([blob], body.url.split('/').at(-1), {
      type: blob.type,
    });
    const formData = new FormData();
    formData.append('image', file);
    const image = await uploadImages(formData);

    return NextResponse.json({
      image,
      success: 1,
      file: {
        url: image[0].img_url,
      },
      toast: { message: 'Upload ảnh thành công!', type: 'success' },
    });
  } catch (error: any) {
    console.error(error);
    return NextResponse.json({
      success: 0,
      toast: { message: error.message, type: 'error' },
    });
  }
}
