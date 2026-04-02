import { catchError, firstValueFrom, throwError } from 'rxjs';
import { Request } from 'express';
import { HttpService } from '@nestjs/axios';
import { HttpException, Injectable } from '@nestjs/common';
import { createProxyMiddleware } from 'http-proxy-middleware';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class HttpProxyService {
  private readonly proxyMiddleware: ReturnType<
    typeof createProxyMiddleware<Request, Response>
  >;

  constructor(private readonly httpService: HttpService) {
    this.proxyMiddleware = createProxyMiddleware({
      target: this.httpService.axiosRef.getUri(),
      changeOrigin: true, // for vhosted sites
      timeout: 5000,
    });
  }

  async makeRequest<T extends Record<string, any> = any>(
    req: Request,
  ): Promise<T> {
    const upstream = await firstValueFrom(
      this.httpService.request<T>({
        method: req.method,
        url: req.originalUrl,
        headers: { ...req.headers, host: undefined },
        data: req.body,
        validateStatus: () => true, // propagate non-2xx as well
      }),
    );

    if (upstream.status >= 400) {
      throw new HttpException(upstream.data, upstream.status);
    }

    // Let Nest handle the response so CacheInterceptor can cache it
    return upstream.data;
  }

  async proxyRequest(req: Request, res: any): Promise<void> {
    this.proxyMiddleware(req, res, (err) => {
      if (err) {
        console.error('Proxy error: ', err);
        throw new HttpException('Proxy error', 500);
      }
    });
  }
}
