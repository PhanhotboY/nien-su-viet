import { PartialType } from '@nestjs/swagger';
import { PostBaseCreateDto } from './post-base-create.dto';

// DTO for updating post
export class PostBaseUpdateDto extends PartialType(PostBaseCreateDto) {}
