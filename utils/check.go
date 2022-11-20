package utils

import "fmt"

func IsTrue(condition bool, msg string, params ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(msg, params...))
	}
}

func IsEmpty() {

}
