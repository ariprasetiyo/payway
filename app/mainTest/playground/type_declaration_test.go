package playground

import (
	"fmt"
	"testing"
)

func TestTypeDeclaration(t *testing.T) {
	type NoKTP string
	type Married bool

	var noKTPAri NoKTP = "1234"
	fmt.Println(noKTPAri)

	fmt.Println("is blacklist admin", isBlacklist("admin", blackListAdmin))
	fmt.Println("is blacklist user", isBlacklist("user", blackListUser))
}

/*
*
contoh lain type declaration
*/
type Blacklist func(string) bool

func isBlacklist(name string, blacklist Blacklist) bool {
	if blacklist(name) {
		return true
	}
	return false
}

func blackListAdmin(name string) bool {
	return name == "admin"
}

func blackListUser(name string) bool {
	return name == "user"
}
