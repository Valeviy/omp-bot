package model

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ConvertPbTimeToNullableTime - convert timestamppb.Timestamp to sql.NullTime
func ConvertPbTimeToNullableTime(pbTime *timestamppb.Timestamp) sql.NullTime {
	var nullableTime sql.NullTime
	if pbTime != nil {
		nullableTime = sql.NullTime{Time: pbTime.AsTime(), Valid: true}
	}

	return nullableTime
}

// ConvertNullableTimeToPbTime - convert  sql.NullTime to timestamppb.Timestamp
func ConvertNullableTimeToPbTime(nullableTime sql.NullTime) *timestamppb.Timestamp {
	var pbTime *timestamppb.Timestamp
	if nullableTime.Valid {
		pbTime = timestamppb.New(nullableTime.Time)
	}

	return pbTime
}
