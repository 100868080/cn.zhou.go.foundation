package main

import (
	"time"
	 "math/rand"
	"testing"
)

func randomPalindrome(rng*rand.Rand)string{
	n:=rng.Intn(25)
runes:=make([] rune,n)
for i:=0;i<(n+1)/2;i++{
	r:=rune(rng.Intn(0x1000))
	runes[i]=r
	runes[n-i-1]=r
}
return string(runes)

}

func TestRandomPalindromes(t*testing.T){
	seed:=time.Now().UTC().UnixNano()
	t.Logf("Random seed:%d",seed)
	rng:=rand.New(rand.NewSource(seed))
for i=:=0;i<1000;i++{
	p:= randomPalindrome(rng)
	if !IsPalindrome(p)
	t.Error("IsPalindrome(%q)=false",p)
}

}