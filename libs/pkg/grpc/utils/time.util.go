package grpcUtils

import (
	"time"

	"google.golang.org/protobuf/types/known/structpb"
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

func JsonToStruct(jsonStr string) *structpb.Struct {
	if jsonStr == "" {
		return &structpb.Struct{}
	}
	s, err := structpb.NewStruct(map[string]any{})
	if err != nil {
		return &structpb.Struct{}
	}
	err = s.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		return &structpb.Struct{}
	}
	return s
}

func StructToJson(s *structpb.Struct) string {
	if s == nil {
		return "{}"
	}
	jsonBytes, err := s.MarshalJSON()
	if err != nil {
		return "{}"
	}
	return string(jsonBytes)
}
