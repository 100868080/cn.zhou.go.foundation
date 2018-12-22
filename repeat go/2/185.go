package word

import (
	"unicode"
	"testing"
)

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t*testing.T){
	var tests=[]struct{
		input string
		want bool
		
	}{
		{"",true}
		{"a",true}
		{"aa",true}
		{"ab",false}
		{"kayak",true}
		{"detartrated",true}
		{"desserts",false}
	}
	for _,test :=range tests{
		if got:=IsPalindrome(test.input);got!=test.want{
			t.Errorf("IsPalindrome(%q)=%v",test.input,got)
		}
	}
}