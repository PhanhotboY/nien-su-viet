import { createParamDecorator, ExecutionContext } from '@nestjs/common';
import { validateJson } from '../util';

/**
 * @description transform incoming payload into camelCase keys with validation
 */
export const ParsedMessage = (dto: any) =>
  createParamDecorator((_data: unknown, ctx: ExecutionContext) => {
    const rpcContext = ctx.switchToRpc();

    let payload = rpcContext.getContext().args[0]?.content?.toString();
    if (!payload) {
      payload = JSON.stringify(rpcContext.getData());
    }

    return validateJson(dto, payload);
  })();
