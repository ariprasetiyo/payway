package playground

import (
	"fmt"
	"testing"
)

func TestFunctionAsNameRetutn(t *testing.T) {
	fmt.Println(functionNamedReturn("coba"))
}

func functionNamedReturn(name string) (name1 string, address string) {
	address = "hahaha"
	name1 = name
	return name1, address
}
