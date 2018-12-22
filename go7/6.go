package mymath

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mult(a, b int) int {
	return a * b
}
func div(a, b int) float32 {
	if b != 0 {
		return float32(a) / float32(b)
	} else {
		return 0
	}

}

/*
func main() {
	Add(a,b int)int
	Sub(a,b int)int
	Mult(a,b int)int
	Div(a,b int)int
}*/

/*
GOPATH=$usr/local/goexport
PATH=$GOBIN:$GOROOGOROOT=$HOME/go
GOBIN=
GOARCH=amd64
GOOS=linux
CGO_ENABLED=1T/bin:$PATH*/
/*
GOROOT=$usr/local/go
GOBIN=
GOARCH=amd64
GOOS=linux
CGO_ENABLED=1
PATH=$GOBIN:$GOROOT/bin:$PATH*/
