'use server';

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

export interface UploadedImage {
  img_url: string;
  public_id: string;
  width: number;
  height: number;
  format: string;
  bytes: number;
}

/**
 * Upload multiple images to Cloudinary
 * @param formData - FormData containing image files and folder name
 * @returns Array of uploaded image details
 */
export async function uploadImages(
  formData: FormData,
): Promise<UploadedImage[]> {
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
    return new Promise<UploadedImage>((resolve, reject) => {
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
            });
          }
        },
      );

      uploadStream.end(buffer);
    });
  });

  return Promise.all(uploadPromises);
}
