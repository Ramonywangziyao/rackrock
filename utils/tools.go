package utils

import (
	"errors"
	"rackrock/model"
	"strconv"
	"time"
)

func ConvertStringToInt64(val string) (int64, error) {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New(model.DataTypeConversionError)
	}

	return i, nil
}

func ConvertStringToTime(val string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", val)
	if err != nil {
		return time.Time{}, errors.New(model.DataTypeConversionError)
	}

	return date, nil
}