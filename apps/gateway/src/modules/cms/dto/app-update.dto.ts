import { PartialType } from '@nestjs/swagger';
import { AppDto } from './app.dto';

export class AppUpdateDto extends PartialType(AppDto) {}
