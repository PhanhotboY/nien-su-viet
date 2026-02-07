import { UploadedImage } from './image.service';

export async function uploadImagesClient(formData: FormData) {
  const response = await fetch('/api/images/upload', {
    method: 'POST',
    body: formData,
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.toast?.message || 'Upload failed');
  }

  const data = await response.json();
  return data.images as UploadedImage[];
}
