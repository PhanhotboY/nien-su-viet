import { firstValueFrom } from 'rxjs';
import { Request } from 'express';
import { HttpService } from '@nestjs/axios';
import { HttpException, Injectable } from '@nestjs/common';

@Injectable()
export class HttpProxyService {
  constructor(private readonly http: HttpService) {}

  async proxyRequest<T extends Record<string, any> = any>(
    req: Request,
  ): Promise<T> {
    const upstream = await firstValueFrom(
      this.http.request<T>({
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
    return upstream.data?.data || upstream.data;
  }
}
