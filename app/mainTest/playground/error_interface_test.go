package playground

import (
	"errors"
	"testing"
)

func TestInterfaceError(t *testing.T) {
	result, error := pembagi(1)
	if error != nil {
		println(error.Error())
	} else {
		println(result)
	}

}

func pembagi(value int8) (int8, error) {
	if value == 0 {
		return 0, errors.New("pembagi 0 menyebabkan error ")
	} else {
		return 1, nil
	}
}
