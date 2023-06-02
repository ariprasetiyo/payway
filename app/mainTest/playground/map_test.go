package playground

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {

	person := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(person)
	person["key3"] = "value3"
	fmt.Println(person)
	person["key3"] = "value33"
	fmt.Println(person)
	delete(person, "key3")
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["tittle"] = "belajar golang"
	book["auth"] = "ari"
	book["ups"] = "salah"
	fmt.Println(book)
}
