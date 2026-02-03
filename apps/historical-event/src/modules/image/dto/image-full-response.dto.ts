import { PickType } from '@nestjs/swagger';
import { ImageFullDto } from './image-full.dto';

export class ImageBriefResponseDto extends PickType(ImageFullDto, [
  'id',
  'caption',
  'publicUrl',
  'redirectLink',
  'type',
]) {}
