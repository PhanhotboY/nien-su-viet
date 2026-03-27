import {
  deleteImage,
  getImages,
  updateImageMetadata,
} from '@/services/image.service';
import { NextRequest, NextResponse } from 'next/server';

/**
 * GET /api/images - Fetch all images from a Cloudinary folder
 */
export async function GET(request: NextRequest) {
  try {
    const folder = request.nextUrl.searchParams.get('folder') || 'nien-su-viet';

    const images = await getImages(folder);

    return NextResponse.json({ images });
  } catch (error) {
    console.error('Failed to fetch images:', error);
    return NextResponse.json(
      { error: 'Failed to fetch images' },
      { status: 500 },
    );
  }
}

/**
 * DELETE /api/images - Delete an image from Cloudinary
 */
export async function DELETE(request: NextRequest) {
  try {
    const { public_id } = await request.json();

    const success = await deleteImage(public_id);
    if (!success) {
      return NextResponse.json(
        { error: 'public_id is required' },
        { status: 400 },
      );
    }
    return NextResponse.json({ success: true });
  } catch (error) {
    console.error('Failed to delete image:', error);
    return NextResponse.json(
      { error: 'Failed to delete image' },
      { status: 500 },
    );
  }
}

/**
 * PUT /api/images - Update an image's metadata in Cloudinary
 */
export async function PUT(request: NextRequest) {
  try {
    const { public_id, metadata } = await request.json();

    const success = await updateImageMetadata(public_id, metadata);
    if (!success) {
      return NextResponse.json(
        { error: 'public_id is required' },
        { status: 400 },
      );
    }
    return NextResponse.json({ success: true });
  } catch (error) {
    console.error('Failed to update image:', error);
    return NextResponse.json(
      { error: 'Failed to update image' },
      { status: 500 },
    );
  }
}
