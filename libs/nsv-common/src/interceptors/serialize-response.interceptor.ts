import {
  Injectable,
  NestInterceptor,
  ExecutionContext,
  CallHandler,
} from '@nestjs/common';
import { Observable, of, pipe } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

@Injectable()
export class SerializeResponseInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    console.log('Before...');
    const res = context.switchToHttp().getResponse();

    const now = Date.now();
    return next
      .handle()
      .pipe(tap(() => console.log(`After... ${Date.now() - now}ms`)))
      .pipe(
        map((data) => {
          const commonizedResponse = (data: Record<string, any>) => {
            return {
              ...data,
              statusCode: res.statusCode,
              timestamp: new Date().toISOString(),
            };
          };

          if (!data) return commonizedResponse({ data: null });

          // if data has data and pagination properties, it's a paginated response
          if (data && data.data && data.pagination) {
            return commonizedResponse(data);
          }

          return commonizedResponse({ data });
        }),
      )
      .pipe(
        catchError((err) => {
          console.log(`After... ${Date.now() - now}ms`);
          console.log('Error occurred: ', err);
          let errMessage = '';

          if (Array.isArray(err?.response?.message)) {
            // class validator error
            errMessage = err?.response?.message?.[0];
          } else if (typeof err?.response?.message === 'string') {
            errMessage = err?.response?.message;
          } else {
            errMessage = err?.message || 'Internal Server Error';
          }

          const errStatus = err?.status || 500;

          return of({
            data: null,
            statusCode: errStatus,
            message: errMessage,
            timestamp: new Date().toISOString(),
          });
        }),
      );
  }
}
