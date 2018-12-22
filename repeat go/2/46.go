func basename(s string) string {
	slash := string.LastIndex(s, "/") //-1 if "/" not found
	s := s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

var f float64=float64(3+0i)
f=float64(2)
f=float64(1e123)
f=foat64('a')

const(
  deadbeef=0xdeadbeef
  a=uint32(deadbeef)
  e=float32(1e309)
  f=uint(-1)

)