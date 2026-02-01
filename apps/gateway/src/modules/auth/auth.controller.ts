import { All, Controller, Get, Req, Res } from '@nestjs/common';
import { createProxyMiddleware } from 'http-proxy-middleware';

import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@gateway/config';

@Controller('auth')
export class AuthController {
  constructor(private readonly config: ConfigService<Config>) { }

  @All('*splat')
  async forwardAuthHandler(@Req() request, @Res() response) {
    const apiProxy = createProxyMiddleware({
      target: this.config.get('authServiceUrl'),
      changeOrigin: true, // for vhosted sites
    });
    apiProxy(request, response);
  }
}
