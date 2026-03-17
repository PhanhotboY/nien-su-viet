import { PartialType } from '@nestjs/swagger';
import {
  PostBaseCreateDto,
  PostBaseCreateGrpcDto,
} from './post-base-create.dto';

// DTO for updating historical event
export class PostBaseUpdateDto extends PartialType(PostBaseCreateDto) {}

export class PostBaseUpdateGrpcDto extends PartialType(PostBaseCreateGrpcDto) {}
