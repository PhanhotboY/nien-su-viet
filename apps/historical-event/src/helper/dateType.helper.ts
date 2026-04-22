import { HISTORICAL_EVENT } from '@phanhotboy/constants';
import { EventDateType } from '@phanhotboy/genproto/historical_event_service/historical_events';

function toGrpcEventDateType(
  dateType?: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>,
): EventDateType {
  switch (dateType) {
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT:
      return EventDateType.EXACT;
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE:
      return EventDateType.APPROXIMATE;
    default:
      return EventDateType.EVENT_DATE_TYPE_UNSPECIFIED;
  }
}

function toEventDateType(
  dateType?: Values<typeof EventDateType>,
): Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE> {
  switch (dateType) {
    case EventDateType.EXACT:
      return HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT;
    case EventDateType.APPROXIMATE:
    case EventDateType.EVENT_DATE_TYPE_UNSPECIFIED:
    default:
      return HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE;
  }
}

export { toGrpcEventDateType, toEventDateType };
