func basename(s string) string {
	slash := string.LastIndex(s, "/") //-1 if "/" not found
	s := s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}