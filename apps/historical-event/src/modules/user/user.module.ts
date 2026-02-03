import { Module } from '@nestjs/common';
import { UserService } from './user.service';
import { UserConsumer } from './user.consumer';
import { RmqModule } from '@phanhotboy/nsv-common';
import { RMQ } from '@phanhotboy/constants';

@Module({
  imports: [RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE })],
  controllers: [UserConsumer],
  providers: [UserService],
  exports: [UserService],
})
export class UserModule {}
