package playground

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	doProcess(false)
	doProcess(true)
}

func doProcess(isPanic bool) {
	defer endProcess()
	if isPanic {
		panic("error app panic")
	}
	fmt.Println("do process...")
}

func endProcess() {
	messagePanic := recover()
	if messagePanic != nil {
		fmt.Println("there are something problems", messagePanic)
	}
	fmt.Println("aplikasi selesai")
}
