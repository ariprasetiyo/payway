package playground

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {

	// mengambil char index ke 0
	fmt.Println("Ini test string"[0])
	// menghitung len = 3
	fmt.Println(len("Ini"))
	// implementasi char get index
	thisString := "Ari Prasetiyo"
	byteThisValue := thisString[0]
	println(string(byteThisValue))
}
