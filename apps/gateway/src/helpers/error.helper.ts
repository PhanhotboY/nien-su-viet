import { HttpStatus } from '@nestjs/common';

function mapErrorNameToStatusCode(errorName: string): number {
  switch (errorName) {
    case 'ConflictException':
    case 'RpcException.ConflictException':
      return HttpStatus.CONFLICT; // 409
    case 'NotFoundException':
    case 'RpcException.NotFoundException':
      return HttpStatus.NOT_FOUND; // 404
    case 'BadRequestException':
    case 'RpcException.BadRequestException':
      return HttpStatus.BAD_REQUEST; // 400
    case 'UnauthorizedException':
    case 'RpcException.UnauthorizedException':
      return HttpStatus.UNAUTHORIZED; // 401
    case 'ForbiddenException':
    case 'RpcException.ForbiddenException':
      return HttpStatus.FORBIDDEN; // 403
    case 'UnprocessableEntityException':
    case 'RpcException.UnprocessableEntityException':
      return HttpStatus.UNPROCESSABLE_ENTITY; // 422
    default:
      return HttpStatus.INTERNAL_SERVER_ERROR; // Default to 500 for unknown error names
  }
}

function mapErrorMessageToStatusCode(errorMessage: string): number {
  if (
    errorMessage.includes('already exists') ||
    errorMessage.includes('duplicate') ||
    errorMessage.includes('unique constraint') ||
    errorMessage.includes('UQ_')
  ) {
    return HttpStatus.CONFLICT; // 409
  }

  if (
    errorMessage.includes('not found') ||
    errorMessage.includes('does not exist')
  ) {
    return HttpStatus.NOT_FOUND; // 404
  }

  if (
    errorMessage.includes('validation') ||
    errorMessage.includes('invalid') ||
    errorMessage.includes('required') ||
    errorMessage.includes('must be') ||
    errorMessage.includes('should not be empty')
  ) {
    return HttpStatus.BAD_REQUEST; // 400
  }

  if (
    errorMessage.includes('timeout') ||
    errorMessage.includes('connection') ||
    errorMessage.includes('ETIMEDOUT')
  ) {
    return HttpStatus.REQUEST_TIMEOUT; // 408
  }

  if (
    errorMessage.includes('unauthorized') ||
    errorMessage.includes('access denied')
  ) {
    return HttpStatus.UNAUTHORIZED; // 401
  }

  if (errorMessage.includes('forbidden')) {
    return HttpStatus.FORBIDDEN; // 403
  }

  if (
    errorMessage.includes('insufficient') ||
    errorMessage.includes('not enough')
  ) {
    return HttpStatus.UNPROCESSABLE_ENTITY; // 422
  }

  // Default to BAD_GATEWAY for microservice communication issues
  return HttpStatus.BAD_GATEWAY; // 502
}

function cleanupErrorMessage(message: string): string {
  // Handle TypeORM constraint errors
  if (message.includes('unique constraint')) {
    if (
      message.includes('sku') ||
      message.includes('UQ_5ec10f972b1fa4f1e60d66d28bc')
    ) {
      return 'Record with this SKU already exists';
    }
    if (message.includes('email')) {
      return 'User with this email already exists';
    }
    if (message.includes('username')) {
      return 'User with this username already exists';
    }
    if (message.includes('name')) {
      return 'Record with this name already exists';
    }
    if (message.includes('phone')) {
      return 'User with this phone number already exists';
    }
    if (message.includes('order_number')) {
      return 'Order with this number already exists';
    }
    if (message.includes('transaction_id')) {
      return 'Transaction with this ID already exists';
    }
    return 'Duplicate entry found';
  }

  // Handle common database errors
  if (message.includes('QueryFailedError')) {
    return 'Database operation failed';
  }

  if (message.includes('duplicate key value violates')) {
    return 'Duplicate entry found';
  }

  // Handle validation errors
  if (message.includes('should not be empty')) {
    return message.replace(/\b\w+\b should not be empty/g, (match) => {
      const field = match.split(' ')[0];
      return `${field} is required`;
    });
  }

  // Handle insufficient stock/points errors
  if (message.includes('insufficient stock')) {
    return 'Insufficient stock available';
  }

  if (message.includes('insufficient points')) {
    return 'Insufficient points available';
  }

  if (message.includes('not enough stock')) {
    return 'Not enough stock available';
  }

  if (message.includes('not enough points')) {
    return 'Not enough points available';
  }

  // Remove technical prefixes
  message = message.replace(
    /^(Error: |QueryFailedError: |ValidationError: )/i,
    '',
  );

  // Remove ip addresses
  message = message.replaceAll(
    /^((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5]):([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$/gm,
    '',
  );

  // Remove url
  message = message.replaceAll(
    /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&\/=]*)/gm,
    '',
  );

  return message;
}

export {
  mapErrorNameToStatusCode,
  mapErrorMessageToStatusCode,
  cleanupErrorMessage,
};
