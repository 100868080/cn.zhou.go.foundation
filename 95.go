func square(n int)int{return n*n}
func negative(n int){return -n}
func product(m,n int)int{return m*n}
f:=square
fmt.println(f(3))   //"g"

f=negative
fmt.println(f(3))  //"-3"
fmt.printf("%T\n",f)   //"func (int) int"
f=product   //compile error:can't assign func (int)int to func (int) int

var f func(int)int
f(3)    //此处f值为nil，会引起panic错误


var f func(int) int
if f!=nil{
	f(3)
}

func add1(r rune)rune{return r+1}
fmt.println(strings.Map(add1,"VMS"))    //"WNT"
fmt.println(strings.Map(add1,"Admix"))   //"Benjy
func forEachNode(n*html,pre,post func(n *html.Node))
if pre!=nil{
	pre(n)
}
for c:=n.FirstChild;c!=nil;c=c.NextSibling{
	forEachNode(c,pre,post)
}
if post!=nil{
	post(n)
}

var depth int
func startElement(n*html.Node){
	if n.Type==html.ElementNode{
		fmt.printf("%*s<%s>\n",depth*2,"",n.Data)
	}
}

func endElement(n *html.Node){
	if n.Type==html.ElementNode{
		depth--
		fmt.printf("%*s</%s>\n",depth*2,"",n.Data)
	}
}

strings.Map(func(r rune)rune{return r+1},"HAL-9000")