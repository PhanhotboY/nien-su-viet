import { Module } from '@nestjs/common';

import { AuthService } from './auth.service';
import { RmqModule } from '@phanhotboy/nsv-common';
import { AuthController } from './auth.controller';

@Module({
  imports: [
    // MailModule,
    RmqModule.register(),
  ],
  controllers: [AuthController],
  providers: [AuthService],
  exports: [AuthService],
})
export class AuthModule {}
