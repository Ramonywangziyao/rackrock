package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"rackrock/model"
	"strconv"
	"time"
)

func ConvertStringToUint64(val string) (uint64, error) {
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, errors.New(model.DataTypeConversionError)
	}

	return i, nil
}

func ConvertStringToTime(val string) (time.Time, error) {
	fmt.Println(val)
	date, err := time.Parse("2006-01-02", val)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, errors.New(model.DataTypeConversionError)
	}

	return date, nil
}

func GenerateRandomId() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}
