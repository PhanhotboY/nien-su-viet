import { PartialType } from '@nestjs/swagger';
import { PostBaseCreateDto } from './post-base-create.dto';

// DTO for updating historical event
export class PostBaseUpdateDto extends PartialType(PostBaseCreateDto) {}
