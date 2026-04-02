package grpcUtils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

func TimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

func TimestampToString(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return TimeToString(TimestampToTime(ts))
}

func TimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
