package playground

import "testing"

func randomData() interface{} {
	return "OK"
}
func TestTypeAssertions(t *testing.T) {
	var dataRandom any = randomData()
	switch dataRandom.(type) {
	case string:
		println("this string")
		break
	default:
		println("unknown type")
	}

	convertToString := dataRandom.(string)
	println(convertToString)

}
