'use server';

import { IImage, ImageMetadata } from '@/types/image';
import { v2 as cloudinary } from 'cloudinary';

// Configure Cloudinary
// Parse CLOUDINARY_URL format: cloudinary://api_key:api_secret@cloud_name
const cloudinaryUrl = process.env.CLOUDINARY_URL;
if (!cloudinaryUrl) {
  throw new Error('CLOUDINARY_URL environment variable is not set');
}

const urlMatch = cloudinaryUrl.match(/cloudinary:\/\/(\d+):([^@]+)@(.+)/);
if (!urlMatch) {
  throw new Error('Invalid CLOUDINARY_URL format');
}

const [, apiKey, apiSecret, cloudName] = urlMatch;

cloudinary.config({
  cloud_name: cloudName,
  api_key: apiKey,
  api_secret: apiSecret,
  secure: true,
});

/**
 * Upload multiple images to Cloudinary
 * @param formData - FormData containing image files and folder name
 * @returns Array of uploaded image details
 */
export async function uploadImages(formData: FormData): Promise<IImage[]> {
  const folder = (formData.get('folder') as string) || 'nien-su-viet';
  const files = formData.getAll('image') as File[];

  if (!files || files.length === 0) {
    throw new Error('No images provided');
  }

  const uploadPromises = files.map(async (file) => {
    // Convert File to Buffer
    const arrayBuffer = await file.arrayBuffer();
    const buffer = Buffer.from(arrayBuffer);

    // Upload to Cloudinary
    return new Promise<IImage>((resolve, reject) => {
      const uploadStream = cloudinary.uploader.upload_stream(
        {
          folder: folder,
          resource_type: 'image',
          transformation: [{ quality: 'auto:good' }, { fetch_format: 'auto' }],
        },
        (error, result) => {
          if (error) {
            reject(error);
          } else if (result) {
            resolve({
              img_url: result.secure_url,
              public_id: result.public_id,
              width: result.width,
              height: result.height,
              format: result.format,
              bytes: result.bytes,
              created_at: result.created_at,
              updated_at: result.updated_at,
              metadata: result.metadata || {},
            });
          }
        },
      );

      uploadStream.end(buffer);
    });
  });

  return Promise.all(uploadPromises);
}

export async function getImages(
  folder: string = 'nien-su-viet',
): Promise<IImage[]> {
  const result = await cloudinary.search
    .expression(`folder="${folder}"`)
    .with_field('context')
    .max_results(100)
    .execute();

  const images = result.resources.map((resource: any) => ({
    public_id: resource.public_id,
    img_url: resource.secure_url,
    width: resource.width,
    height: resource.height,
    format: resource.format,
    bytes: resource.bytes,
    created_at: resource.created_at,
    updated_at: resource.updated_at,
    metadata: resource.context || {},
  }));
  return images;
}

export async function deleteImage(publicId: string): Promise<boolean> {
  if (!publicId) {
    return false;
  }
  await cloudinary.uploader.destroy(publicId);
  return true;
}

export async function updateImageMetadata(
  publicId: string,
  metadata: ImageMetadata,
): Promise<boolean> {
  if (!publicId) {
    return false;
  }

  const context = {
    ...(metadata.alt && { alt: metadata.alt }),
    ...(metadata.caption && { caption: metadata.caption }),
    ...(metadata.category && { category: metadata.category }),
  };
  await cloudinary.uploader.explicit(publicId, {
    type: 'upload',
    context,
    ...(metadata.transformations && {
      transformation: metadata.transformations,
    }),
  });
  return true;
}
