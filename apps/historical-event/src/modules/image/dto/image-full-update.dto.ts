import { PartialType } from '@nestjs/swagger';
import { ImageFullCreateDto } from './image-full-create.dto';

export class ImageFullUpdateDto extends PartialType(ImageFullCreateDto) {}
