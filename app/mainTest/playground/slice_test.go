package playground

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// jika array pake[...]
	var sliceData [2]string
	sliceData[0] = "hal1"
	sliceData[1] = "hal2"
	fmt.Println(sliceData)

	//array
	var arraydata [2]string
	arraydata[0] = "hal1"
	arraydata[1] = "hal2"
	fmt.Println(sliceData)

	sliceD := append(arraydata[:], "nambah data")
	fmt.Println(sliceD)

	//membuat slice lebih aman tanpa mendefinisikan array dlu
	var slice2 = make([]string, 2, 5)
	slice2[0] = "1"
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))

	copySlice := make([]string, len(slice2), cap(slice2))
	copy(slice2, copySlice)
	fmt.Println(copySlice)

}
