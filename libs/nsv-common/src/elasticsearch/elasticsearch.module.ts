import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { ElasticsearchModule } from '@nestjs/elasticsearch';

@Module({
  imports: [
    ConfigModule,
    ElasticsearchModule.registerAsync({
      useFactory: (config: ConfigService) => {
        const node = config.get<string>('ELASTICSEARCH_NODE');
        if (!node) {
          throw new Error('ELASTICSEARCH_NODE is not defined');
        }

        return {
          node,
        };
      },
      inject: [ConfigService],
    }),
  ],
  exports: [ElasticsearchModule],
})
export class SearchModule {}
