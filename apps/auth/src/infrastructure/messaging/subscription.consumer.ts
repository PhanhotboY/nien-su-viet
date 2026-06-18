import { Controller } from '@nestjs/common';
import { Ctx, EventPattern, Payload, RmqContext } from '@nestjs/microservices';
import { RmqService } from '@phanhotboy/nsv-common';
import { UserService } from '@auth/application/services/user.service';
import { SUBSCRIPTION_EVENT } from '@phanhotboy/constants';
// import { UserBaseDto, UserDeleteDto } from './dto';

@Controller()
export class SubscriptionConsumer {
  constructor(
    private readonly userService: UserService,
    private readonly rmqService: RmqService,
  ) {}

  @EventPattern(SUBSCRIPTION_EVENT.CREATED)
  async handleSubscriptionCreatedEvent(
    @Payload() data: any,
    @Ctx() context: RmqContext,
  ) {
    // await this.userService
    //   .handleUserRegister(data)
    //   .catch((error) => {
    //     console.error('Error handling user registered event:', error);
    //   })
    //   .finally(() => {
    //     this.rmqService.ack(context);
    //   });
  }
}
