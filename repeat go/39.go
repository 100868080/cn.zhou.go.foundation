package main

func main() {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		//...ASCII letter or digit...
	}
	i := 0
	if b {
		i = 1
	}
}

//btoi returns 1 if b is true and 0 if false.
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob reports whether i is non-zero.
func itob(i int) bool {
	return i != 0
}
