import { HISTORICAL_EVENT } from '@phanhotboy/constants';
import { EventDateType } from '@phanhotboy/genproto/historical_event_service/historical_events';

function toGrpcEventDateType(
  dateType?: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>,
): EventDateType {
  switch (dateType) {
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT:
      return EventDateType.EXACT;
    case HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE:
    default:
      return EventDateType.APPROXIMATE;
  }
}

const eventDateTypeValues = Object.values(HISTORICAL_EVENT.EVENT_DATE_TYPE);
function toEventDateType(
  dateType?: Values<typeof EventDateType>,
): Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE> {
  if (eventDateTypeValues.includes(dateType as any)) {
    return dateType as any;
  }
  switch (dateType) {
    case EventDateType.EXACT:
      return HISTORICAL_EVENT.EVENT_DATE_TYPE.EXACT;
    case EventDateType.APPROXIMATE:
    default:
      return HISTORICAL_EVENT.EVENT_DATE_TYPE.APPROXIMATE;
  }
}

export { toGrpcEventDateType, toEventDateType };
