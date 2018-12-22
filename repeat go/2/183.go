package word

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != S[len(s)-i] {
			return false
		}
	}
	return true
}
