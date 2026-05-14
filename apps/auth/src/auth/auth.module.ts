import { Module } from '@nestjs/common';
import { AuthService } from './auth.service';
import { RmqModule } from '@phanhotboy/nsv-common';
import { AuthController } from './auth.controller';
import { RMQ } from '@phanhotboy/constants';

@Module({
  imports: [
    // MailModule,
    RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE }),
  ],
  controllers: [AuthController],
  providers: [AuthService],
  exports: [AuthService],
})
export class AuthModule {}
