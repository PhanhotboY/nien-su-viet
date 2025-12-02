import { DynamicModule, Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { ConfigService } from '../providers';
import { RmqService } from './rmq.service';

export interface RmqModuleOptions {
  name: string;
}

@Module({
  providers: [RmqService],
  exports: [RmqService],
})
export class RmqModule {
  static register({ name }: RmqModuleOptions): DynamicModule {
    return {
      module: RmqModule,
      imports: [
        ClientsModule.registerAsync([
          {
            name,
            inject: [ConfigService],
            useFactory: (configService: ConfigService) => {
              const rmqUrl = configService.get('RABBITMQ_URL');
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
      exports: [ClientsModule],
    };
  }
}
