package playground

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var thisArray = [...]string{"coba"}
	thisArray[0] = "data"
	fmt.Println(thisArray)

	var thisArray2 [3]string
	thisArray2[0] = "data1"
	thisArray2[1] = "data2"
	thisArray2[2] = "data3"
	fmt.Println(thisArray2)
}
