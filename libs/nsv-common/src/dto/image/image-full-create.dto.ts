import { OmitType } from '@nestjs/swagger';
import { ImageFullDto } from './image-full.dto';

// DTO for creating image
export class ImageFullCreateDto extends OmitType(ImageFullDto, ['id']) {}
