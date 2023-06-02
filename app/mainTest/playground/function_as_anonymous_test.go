package playground

import (
	"fmt"
	"testing"
)

func TestAnonymousFunc(t *testing.T) {
	blacklist := func(name string) bool {
		return name == "black"
	}

	fmt.Println(blacklist("coba"))
}
