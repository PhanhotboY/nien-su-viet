import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { MailModule, MailService } from '@auth/mail';
import { RmqModule, RMQ } from '@phanhotboy/nsv-common';
import { AuthController } from './auth.controller';

@Module({
  imports: [
    // MailModule,
    // RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE }),
  ],
  controllers: [AuthController],
  providers: [AuthService],
  exports: [AuthService],
})
export class AuthModule {}
