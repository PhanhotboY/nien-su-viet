import {
  DynamicModule,
  Inject,
  Logger,
  Module,
  OnModuleInit,
} from '@nestjs/common';
import { ClientProxy, ClientsModule, Transport } from '@nestjs/microservices';
import { ConfigService } from '../providers';
import { RmqService } from './rmq.service';
import { RMQ } from '@phanhotboy/constants';

export interface RmqModuleOptions {
  name: string;
}

@Module({
  providers: [RmqService],
  exports: [RmqService],
})
export class RmqModule implements OnModuleInit {
  static register(): DynamicModule {
    return {
      module: RmqModule,
      imports: [
        ClientsModule.registerAsync([
          {
            name: RMQ.TOPIC_EVENTS_EXCHANGE,
            inject: [ConfigService],
            useFactory: (configService: ConfigService) => {
              const rmqUrl = configService.get('rabbitmq');
              if (!rmqUrl) {
                throw new Error('Missing RabbitMQ configuration in .env file.');
              }

              return {
                transport: Transport.RMQ,
                options: {
                  urls: [rmqUrl],
                  wildcards: true,
                  exchange: 'events',
                  exchangeType: 'topic',
                },
              };
            },
          },
        ]),
      ],
      providers: [RmqService],
      exports: [ClientsModule, RmqService],
    };
  }

  constructor(
    @Inject(RMQ.TOPIC_EVENTS_EXCHANGE) private readonly rmq: ClientProxy,
    private readonly logger: Logger,
  ) {}
  onModuleInit() {
    this.rmq.connect().then(() => {
      this.logger.log('Connected to RabbitMQ');
    });
  }
}
