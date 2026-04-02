import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { HttpModule } from '@nestjs/axios';
import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@gateway/config';
import { HttpProxyService } from '@gateway/common/services/http-proxy.service';

@Module({
  imports: [
    HttpModule.registerAsync({
      useFactory: (config: ConfigService<Config>) => ({
        baseURL: config.get('authServiceEndpoint') || 'http://localhost:4001',
        maxRedirects: 5,
        timeout: 5000,
      }),
      inject: [ConfigService],
    }),
  ],
  controllers: [AuthController],
  providers: [HttpProxyService],
})
export class AuthModule {}
