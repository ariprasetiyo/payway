package playground

import (
	"fmt"
	"testing"
)

func TestInterfaceKosong(T *testing.T) {
	fmt.Println(coba())

}

/*
interface kosonbg bisa pake any
*/
func coba() interface{} {

	switch "" {
	case "1":
		return 1
	default:
		return ""
	}
}

func coba1() any {

	switch "" {
	case "1":
		return 1
	default:
		return ""
	}
}
