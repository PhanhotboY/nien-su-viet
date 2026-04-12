import { status } from '@grpc/grpc-js';
import { HttpStatus } from '@nestjs/common';

function mapGrpcCodeToHttpStatusCode(code: number): number {
  switch (code) {
    case status.OK: // OK
      return HttpStatus.OK; // 200
    case status.CANCELLED: // CANCELLED
      return HttpStatus.REQUEST_TIMEOUT; // 408
    case status.UNKNOWN: // UNKNOWN
      return HttpStatus.INTERNAL_SERVER_ERROR; // 500
    case status.INVALID_ARGUMENT: // INVALID_ARGUMENT
      return HttpStatus.BAD_REQUEST; // 400
    case status.DEADLINE_EXCEEDED: // DEADLINE_EXCEEDED
      return HttpStatus.REQUEST_TIMEOUT; // 408
    case status.NOT_FOUND: // NOT_FOUND
      return HttpStatus.NOT_FOUND; // 404
    case status.ALREADY_EXISTS: // ALREADY_EXISTS
      return HttpStatus.CONFLICT; // 409
    case status.PERMISSION_DENIED: // PERMISSION_DENIED
      return HttpStatus.FORBIDDEN; // 403
    case status.RESOURCE_EXHAUSTED: // RESOURCE_EXHAUSTED
      return HttpStatus.TOO_MANY_REQUESTS; // 429
    case status.FAILED_PRECONDITION: // FAILED_PRECONDITION
      return HttpStatus.BAD_REQUEST; // 400
    case status.ABORTED: // ABORTED
      return HttpStatus.CONFLICT; // 409
    case status.OUT_OF_RANGE: // OUT_OF_RANGE
      return HttpStatus.BAD_REQUEST; // 400
    case status.UNIMPLEMENTED: // UNIMPLEMENTED
      return HttpStatus.NOT_IMPLEMENTED; // 501
    case status.INTERNAL: // INTERNAL
      return HttpStatus.INTERNAL_SERVER_ERROR; // 500
    case status.UNAVAILABLE: // UNAVAILABLE
      return HttpStatus.SERVICE_UNAVAILABLE; // 503
    case status.DATA_LOSS: // DATA_LOSS
      return HttpStatus.INTERNAL_SERVER_ERROR; // 500
    case status.UNAUTHENTICATED: // UNAUTHENTICATED
      return HttpStatus.UNAUTHORIZED; // 401
    default:
      return HttpStatus.INTERNAL_SERVER_ERROR; // Default to 500 for unknown codes
  }
}

export { mapGrpcCodeToHttpStatusCode };
