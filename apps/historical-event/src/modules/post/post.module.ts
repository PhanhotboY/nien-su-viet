import { Module } from '@nestjs/common';
import { RmqModule } from '@phanhotboy/nsv-common';
import { RMQ } from '@phanhotboy/constants';
import { PostController } from './post.controller';
import { PostService } from './post.service';

@Module({
  imports: [RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE })],
  controllers: [PostController],
  providers: [PostService],
})
export class PostModule {}
