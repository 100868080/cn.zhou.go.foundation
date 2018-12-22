package word

import "testing"

func Testalindrome(t *testing.T) {
	if !IsPalindrome("detartrtated") {
		t.Error(`IsPalindrome("detartarted")=false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak")=false`)
	}
}

func TestNonPalindrome(r *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome")=true`)
	}
}
