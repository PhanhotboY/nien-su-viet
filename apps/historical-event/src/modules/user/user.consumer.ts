import { Controller } from '@nestjs/common';
import { Ctx, EventPattern, Payload, RmqContext } from '@nestjs/microservices';
import {
  RmqService,
  USER_EVENT,
  UserBaseDto,
  UserDeleteDto,
} from '@phanhotboy/nsv-common';
import { UserService } from './user.service';

@Controller()
export class UserConsumer {
  constructor(
    private readonly userService: UserService,
    private readonly rmqService: RmqService,
  ) {}

  @EventPattern(USER_EVENT.REGISTERED)
  async handleUserRegisteredEvent(
    @Payload() data: UserBaseDto,
    @Ctx() context: RmqContext,
  ) {
    await this.userService
      .handleUserRegister(data)
      .catch((error) => {
        console.error('Error handling user registered event:', error);
      })
      .finally(() => {
        this.rmqService.ack(context);
      });
  }

  @EventPattern(USER_EVENT.DELETED)
  async handleUserDeletedEvent(
    @Payload() data: UserDeleteDto,
    @Ctx() context: RmqContext,
  ) {
    await this.userService
      .deleteUser(data.userId)
      .catch((error) => {
        console.error('Error handling user deleted event:', error);
      })
      .finally(() => {
        return this.rmqService.ack(context);
      });
  }

  @EventPattern(USER_EVENT.UPDATED)
  async handleUserUpdatedEvent(
    @Payload() data: UserBaseDto,
    @Ctx() context: RmqContext,
  ) {
    await this.userService
      .updateUser(data)
      .catch((error) => {
        console.error('Error handling user updated event:', error);
      })
      .finally(() => {
        return this.rmqService.ack(context);
      });
  }
}
