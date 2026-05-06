import { HISTORICAL_EVENT } from '@phanhotboy/constants';
import { EventDateType } from '@phanhotboy/genproto/historical_event_service/historical_events';

function toEventDateType(dateType?: EventDateType) {
  switch (dateType) {
    case EventDateType.EXACT:
      return HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT;
    case EventDateType.APPROXIMATE:
    default:
      return HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE;
  }
}

function toGrpcEventDateType(
  dateType?: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>,
) {
  switch (dateType) {
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT:
      return EventDateType.EXACT;
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE:
    default:
      return EventDateType.APPROXIMATE;
  }
}

export { toEventDateType, toGrpcEventDateType };
