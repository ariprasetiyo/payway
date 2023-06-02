package playground

import (
	"fmt"
	"testing"
)

func TestFuctionAsParameter(t *testing.T) {
	/**
	fucntion as paramter could be used for software design pattern strategy pattern
	**/
	functionAsParameter("coba", spamFilter)
	functionAsParameter("coba", spamFilter2)
}

func functionAsParameter(spamValue string, spamFilters func(string) string) {
	resultFilter := spamFilters(spamValue)
	fmt.Println(resultFilter)
}

func spamFilter(spamvalie string) string {
	switch spamvalie {
	case "coba":
		return "coba"
	default:
		return "tdk kena filter"
	}
}

func spamFilter2(spamvalie string) string {
	switch spamvalie {
	case "coba1":
		return "coba1"
	default:
		return "tdk kena filter"
	}
}
