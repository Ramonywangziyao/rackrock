package utils

import (
	"encoding/json"
	"fmt"
)

func InterfaceContains(src interface{}, arr interface{}) bool {
	return true
}

//  check fn  last err
func panicFunc(fn interface{}) {

}

func MustMarshal(data interface{}) []byte {
	var res, err = json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("marshal data: %+v err: %s", data, err))
	}

	return res
}

func MustUnmarshal(bytes []byte, data interface{}) {
	if err := json.Unmarshal(bytes, data); err != nil {
		panic(fmt.Sprintf("unmarshal data err: %s", err.Error()))
	}
}
