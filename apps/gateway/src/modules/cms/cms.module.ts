import { Module } from '@nestjs/common';
import { HttpModule } from '@nestjs/axios';

import { PingController } from './ping.controller';
import { AppController } from './app.controller';
import { HeaderNavItemController } from './header-nav-item.controller';
import { FooterNavItemController } from './footer-nav-item.controller';
import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { MediaController } from './media.controller';
import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@gateway/config';

@Module({
  imports: [
    HttpModule.registerAsync({
      useFactory: (config: ConfigService<Config>) => ({
        baseURL: config.get('cmsServiceEndpoint') || 'http://localhost:4001',
        maxRedirects: 5,
        timeout: 5000,
      }),
      inject: [ConfigService],
    }),
  ],
  controllers: [
    PingController,
    AppController,
    MediaController,
    HeaderNavItemController,
    FooterNavItemController,
  ],
  providers: [HttpProxyService],
})
export class CmsModule {}
