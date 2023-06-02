package playground

import (
	"fmt"
	"testing"
)

func TestLooping(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	slice := []string{
		"haha1", "haha2",
	}

	for _, value := range slice {
		fmt.Println(value)
	}
}
