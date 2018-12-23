package main

import (
	"fmt"
	"testing"
	"unicode"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man,a plan,a cannal:Panama")
	}
}
n:=len( letters)/2
for i:=0;i<n;i++{
	if letters[i]!=letters[len(letters)-1-i]{
		return false
	}
}
return true


letters:=make([]rune,0,len(s))
for _,r:=range s{
	if unicode.IsLetter(r){
		letters=append(letters,unicode.ToLower(r))
	}
}

func ExampleIsPalindrome(){
	fmt.Println(IsPalindrome("A man,a plan,a canal:Panama"))
	fmt.Println(IsPalindrome("palindrome"))
}