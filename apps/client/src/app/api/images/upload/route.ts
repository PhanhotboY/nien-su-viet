import { NextRequest, NextResponse } from 'next/server';
import { authClient } from '@/lib/auth-client'; // Adjust based on your auth setup
import { uploadImages } from '@/services/image.service'; // Adjust based on your service location

// Next.js App Router POST handler (replaces Remix action)
export async function POST(request: NextRequest) {
  // Get session using Better Auth SDK
  const session = await authClient.getSession();

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
    const formData = await request.formData();

    const images = await uploadImages(formData);

    return NextResponse.json(
      {
        images,
        success: 1,
        file: {
          url: images[0].img_url,
        },
        toast: { message: 'Upload ảnh thành công!', type: 'success' },
      },
      { status: 200 },
    );
  } catch (error: any) {
    console.error(error);
    return NextResponse.json(
      {
        success: 0,
        toast: { message: error.message, type: 'error' },
      },
      { status: 500 },
    );
  }
}
